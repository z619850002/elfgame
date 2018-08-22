#!/usr/bin/env python3
# -*- coding: utf8 -*-

from concurrent import futures
from datetime import datetime
from gevent.queue import Queue

from rlpytorch import Evaluator, load_env
from console_lib import GoConsoleGTP, move2xy, xy2move

from google.rpc import code_pb2 as google_dot_rpc_dot_code__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2

import gevent
import grpc

import os
import signal
import sys
import traceback

import game_pb2
import game_pb2_grpc
import threading
from time import sleep

# Load env
additional_to_load = {
    'evaluator': (
        Evaluator.get_option_spec(),
        lambda object_map: Evaluator(object_map, stats=None)),
}

# Set os environment
os.environ.update({
    'game': 'elfgames.go.game',
    'model': 'df_pred',
    'model_file': 'elfgames.go.df_model3',
})

overrides = {
    'num_games': 1,
    'greedy': True,
    'T': 1,
    'model': 'online',
    'additional_labels': ['aug_code', 'move_idx'],
    'check_loaded_options': False,
    'keys_in_reply':  ['V', 'rv'],
    'use_mcts': True,
    'mcts_verbose_time': True,
    'mcts_use_prior': True,
    'mcts_persistent_tree': True,
    'load': '/home/wuhan/pytorch/pretrained-go-19x19-v1.bin',
    'server_addr': 'localhost',
    'port': 1234,
    'replace_prefix': ['resnet.module,resnet'],
    'no_parameter_print': True,
    'verbose': True,
    'gpu': 0,
    'num_block': 20,
    'dim': 224,
    'mcts_puct': 1.50,
    'batchsize': 16,
    'mcts_rollout_per_batch': 16,
    'mcts_threads': 2,
    'mcts_rollout_per_thread': 64,  # bigger value will spend more time to genmove
    'resign_thres': 0.05,
    'mcts_virtual_loss': 1,
}

# Set game to online model.
env = load_env(
    os.environ,
    overrides=overrides,
    additional_to_load=additional_to_load)

GC_PARAMS = {
    'ACTION_CLEAR': -97,
    'ACTION_PASS': -99,
    'ACTION_RESIGN': -98,
    'ACTION_SKIP': -100,
    'board_size': 19,
    'num_action': 362,
    'num_future_actions': 1,
    'num_planes': 18,
    'opponent_stone_plane': 1,
    'our_stone_plane': 0,
    'num_group': 2,
    'T': 1
}

model_loader = env["model_loaders"][0]
model = model_loader.load_model(GC_PARAMS)

mi = env['mi']
mi.add_model("model", model)
mi.add_model("actor", model)
mi["model"].eval()
mi["actor"].eval()
####


def debug_print(message):
    print("\033[92m{message}\033[0m".format(message=message))

#one game context is just one chess game
#the game context will deal with the command in its channel continuely
class GameContext:

    def __init__(self):
        self.action = None
        self.last_cmd = ""
        #at the same time the ai can only make response to the same operation
        self.reqChan = Queue(maxsize=1)
        self.respChan = Queue(maxsize=1)
        self._load_env()

    def _load_env(self):
        #the console of one game context is just the object of the GOConsoleGTP
        self.console = GoConsoleGTP(env["game"].initialize(), env["evaluator"])

    #the return value of this function is a dictionary, it`s sent to the console, while the reply to the game object is by the response channel
    def human_actor(self, batch):
        self.batch = batch

        if self.last_cmd == "play" or self.last_cmd == "genmove":
            #resign means give up
            if self.console.get_last_move(batch).lower() == "resign":
                #the last game is end， then we need to calculate for the final score
                self.respChan.put_nowait(game_pb2.Reply(final_score=self.console.get_final_score(batch),
                                                 status=google_dot_rpc_dot_status__pb2.Status(code=google_dot_rpc_dot_code__pb2.OK),
                                                 resigned=True))
            else:
                #last game is not end, some new step is needed
                x, y = move2xy(self.console.get_last_move(batch))
                self.respChan.put_nowait(game_pb2.Reply(coordinate=game_pb2.Coordinate(x=x, y=y),
                                                 next_player=self.console.get_next_player(batch),
                                                 final_score=self.console.get_final_score(batch),
                                                 resigned=False,
                                                 status=google_dot_rpc_dot_status__pb2.Status(code=google_dot_rpc_dot_code__pb2.OK),
                                                 last_move=self.console.get_last_move(batch),
                                                 board=batch.GC.getGame(0).showBoard()))
        elif self.last_cmd == "clear_board":
            self.respChan.put_nowait(google_dot_rpc_dot_status__pb2.Status(code=google_dot_rpc_dot_code__pb2.OK))

        self.last_cmd = ""
        while self.reqChan.empty():
            gevent.sleep(0)

        cmd = self.reqChan.get()
        #update the cmd to this time
        self.last_cmd = cmd
        if cmd == "play":
            return dict(pi=None, a=self.action, V=0)
        elif cmd == "genmove":
            return dict(pi=None, a=self.console.actions["skip"], V=0)
        elif cmd == "pass":
            return dict(pi=None, a=self.console.actions["pass"], V=0)
        elif cmd == "clear_board":
            return dict(pi=None, a=self.console.actions["clear"], V=0)
        elif cmd == "resign":
            return dict(pi=None, a=self.console.actions["resign"], V=0)
        elif cmd == "exit":
            self.console.exit = True
            return dict(pi=None, a=None, V=0)


#this is a game
#it should be regist to the server
#the command input by user will be dealt by this class
#at the same time, it will send message to the game context
#the game context has a channel which can help it deal with the command orderly
class GoGame(game_pb2_grpc.GameServicer):

    def finishNewThread(self):
        self.finish = 0

    def __init__(self, maxsize=10):
        self.finish = 1
        self.connectionTime = 0
        self._players = {}
        self.gc_pools = Queue(maxsize=maxsize)
        self.greenlets = []
        self.start_console()
        self.lock = False

        def heartBeat():
            while(self.finish==1):
                #lock this block
                #connection is broken
                if (self.connectionTime > 5):
                    self.lock = True
                    #clear the board now
                    for gcItem in self._players.values():
                        gcItem.reqChan.putnowait("clear_board")
                        while gcItem.respChan.empty():
                            gevent.sleep(0)
                        gcItem.respChan.get()
                    #free all game contexts
                    keys = []
                    for playerId in self._players.keys():
                        keys.append(playerId)
                    for playerId in keys:
                        self.gc_pools.put_nowait(self._players.pop(playerId))
                    print("connection out of time, free all gc")
                    self.connectionTime = -1
                    self.lock = False
                #calculate for the connection time
                self.connectionTime = self.connectionTime + 1
                #release the lock
                self.lock.release()
                sleep(1)
        try:
            self.newThread = threading.Thread(target=heartBeat)
            self.newThread.start()
        except:
            print("Error: unable to start thread")

    #start the console module
    def _start_console(self):
        #create a game context first
        #the gc is just the game context below
        gc = GameContext()


        #add the game context to the pool( 10 gc is allowed to add to the pool in default)
        self.gc_pools.put_nowait(gc)

        #define 2 functions here
        #the console of gc is the GoConsoleGTP
        def actor(batch):
            return gc.console.actor(batch)

        #prompt a command of train
        #now I consider the batch as a kind of operaion
        def train(batch):
            gc.console.prompt("DF Train> ", batch)


        #the first gc is just the game context below while the second one
        #is the game context in source code
        gc.console.evaluator.setup(sampler=env["sampler"], mi=mi)

        gc.console.GC.reg_callback_if_exists("actor_black", actor)
        gc.console.GC.reg_callback_if_exists("human_actor", gc.human_actor)
        gc.console.GC.reg_callback_if_exists("train", train)

        gc.console.GC.start()
        gc.console.GC.GC.getClient().setRequest(
            mi["actor"].step, -1, env['game'].options.resign_thres, -1)

        gc.console.evaluator.episode_start(0)

        while True:
            gc.console.GC.run()
            if gc.console.exit:
                break
        gc.console.GC.stop()

    def start_console(self):
        while not self.gc_pools.full():
            g = gevent.spawn(self._start_console)
            g.start()
            gevent.sleep(0)
            self.greenlets.append(g)

    def stop_console(self):
        gevent.killall(self.greenlets, timeout=10)

    #one game context is related to one game
    #the function here will return the protobuf
    #the function like clear board, play etc., will block here until the response channel is not empty
    def NewGC(self, request, context):
        while (self.lock):
            pass
        #offered by the client
        playerId = request.id
        if not playerId in self._players:
            if self.gc_pools.empty():
                return google_dot_rpc_dot_status__pb2.Status(code=google_dot_rpc_dot_code__pb2.RESOURCE_EXHAUSTED,
                                                             message="no more gc is available")
            self._players[playerId] = self.gc_pools.get()
        return google_dot_rpc_dot_status__pb2.Status(code=google_dot_rpc_dot_code__pb2.OK)

    def FreeGC(self, request, context):
        while (self.lock):
            pass
        playerId = request.id
        if playerId not in self._players:
            return google_dot_rpc_dot_status__pb2.Status(code=google_dot_rpc_dot_code__pb2.NOT_FOUND)
        self.gc_pools.put_nowait(self._players.pop(playerId))
        return google_dot_rpc_dot_status__pb2.Status(code=google_dot_rpc_dot_code__pb2.OK)

    def ClearBoard(self, request, context):
        while (self.lock):
            pass
        gc = self._players[request.id]
        #now the game will send message to the game context`s channel,
        #  then those contexts will deal with them orderly
        gc.reqChan.put_nowait("clear_board")
        while gc.respChan.empty():
            gevent.sleep(0)
        debug_print("Clear board for player {id}".format(id=request.id))
        return gc.respChan.get()

    def Play(self, request, context):
        while (self.lock):
            pass
        gc = self._players[request.player.id]
        try:
            gc.action = gc.console.move2action(request.move)
        except ValueError as e:
            return game_pb2.Reply(status=google_dot_rpc_dot_status__pb2.Status(code=google_dot_rpc_dot_code__pb2.INVALID_ARGUMENT,
                                                                               message=str(e)))
        gc.reqChan.put_nowait("play")
        while gc.respChan.empty():
            gevent.sleep(0)
        return gc.respChan.get()

    def GenMove(self, request, context):
        while (self.lock):
            pass
        gc = self._players[request.id]
        gc.reqChan.put_nowait("genmove")
        #until the gc make response this object will be blocked
        while gc.respChan.empty():
            gevent.sleep(0)
        return gc.respChan.get()

    def Pass(self, request, context):
        while (self.lock):
            pass
        gc = self._players[request.id]
        gc.reqChan.put_nowait("pass")
        while gc.respChan.empty():
            gevent.sleep(0)
        return gc.respChan.get()

    def Resign(self, request, context):
        while (self.lock):
            pass
        gc = self._players[request.id]
        gc.reqChan.put_nowait("resign")
        while gc.respChan.empty():
            gevent.sleep(0)
        return gc.respChan.get()

    def HeartBeat(self, request , context):
        while (self.lock):
            pass
        self.connectionTime = 0
        return game_pb2.BeatReply(beatReply = 1)

class GameServer:

    def __init__(self, listen, workers=1):
        #initialize the exit channel
        self.exitChan = Queue(maxsize=1)
        #server is the grpc server
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=workers))

        #initialize the game
        self.game = GoGame(maxsize=10)

        #the server here is the grpc server
        self.server.add_insecure_port(listen)

        #regist the game to the server
        #the game here is just the gogame created above
        game_pb2_grpc.add_GameServicer_to_server(self.game, self.server)

    def start(self):
        #start the server
        self.server.start()
        while self.exitChan.empty():
            gevent.sleep(0)
        #stop the console of the game
        self.game.stop_console()
    
    def stop(self):
        #stop the server
        self.server.stop(0)
        #put a element into the channel , so the channel is not empty and the loop will end
        self.exitChan.put_nowait(0)


def main():
    #listen to the port
    listen = os.getenv("GAME_LISTEN", "[::]:10000")
    gameserver = GameServer(listen)
    gevent.signal(signal.SIGTERM, gameserver.stop)
    gevent.signal(signal.SIGINT, gameserver.stop)
    debug_print("Starting game server on {listen}".format(listen=listen))
    gevent.spawn(gameserver.start).join()
    gameserver.game.finishNewThread()
    debug_print("Gameserver is stopped")


if __name__ == "__main__":
    main()
