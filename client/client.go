package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/rpc/code"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	game "hub000.xindong.com/pytorch/elfgame/client/proto"
)

func readCoordinate(name string) int32 {
	fmt.Print(color.BlueString(fmt.Sprintf("enter coordinate %s: ", name)))
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("read from console error: %v", err)
	}
	coordinate, err := strconv.ParseInt(strings.TrimSpace(input), 10, 32)
	if err != nil {
		log.Fatalf("Invalid input: %v", err)
	}
	return int32(coordinate)
}

func readMove() string {
	fmt.Print(color.BlueString(fmt.Sprintf("Enter move: ")))
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("read from console error: %v", err)
	}
	move := strings.TrimSpace(input)
	if len(move) == 0 {
		log.Fatalf("Invalid input: %v", err)
	}
	return move
}

func main() {
	var (
		playerID string
		address  string
	)

	flag.StringVar(&playerID, "player", uuid.New().String(), "player id")
	flag.StringVar(&address, "address", "172.26.160.19:10000", "server address")
	flag.Parse()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	defer conn.Close()

	c := game.NewGameClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	player := &game.Player{Id: playerID}
	r, err := c.NewGC(ctx, player)
	if err != nil {
		log.Fatalf("start error: %v", err)
	}
	if r.Code != int32(code.Code_OK) {
		log.Fatalf("start error: %s", r.Message)
	}
	fmt.Println(color.GreenString("Welcome, player %s", playerID))

	for {
		fmt.Print(color.BlueString("user input > "))
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalf("read from console error: %v", err)
		}

		switch strings.TrimSpace(command) {
		case "selfplay":
			for {
				r, err := c.GenMove(context.Background(), player)
				if err != nil {
					log.Fatalf("genmove error: %v", err)
				}
				fmt.Println(r.Board)
				if r.Resigned {
					if r.FinalScore > 0 {
						fmt.Printf("B+%.1f\n", r.FinalScore)
					} else {
						fmt.Printf("W+%.1f\n", -r.FinalScore)
					}
					break
				}
			}
		case "clear_board":
			r, err := c.ClearBoard(context.Background(), player)
			if err != nil {
				log.Fatalf("clear board error: %v", err)
			}
			if r.Code == int32(code.Code_OK) {
				fmt.Println("clear board ok")
			}
		case "play":
			// coordinateX := readCoordinate("X")
			// coordinateY := readCoordinate("Y")
			move := readMove()
			request := &game.Step{
				//X:      coordinateX,
				//Y:      coordinateY,
				Player: player,
				Move:   move,
			}
			r, err := c.Play(context.Background(), request)
			if err != nil {
				log.Fatalf("play error: %v", err)
			}
			if r.Status.Code != int32(code.Code_OK) {
				fmt.Printf("play error: %s\n", r.Status.Message)
				continue
			}
			fmt.Printf("Player placed a stone at move %s, next player is %s\n", r.LastMove, r.NextPlayer)
			fmt.Println(r.Board)
		case "genmove":
			r, err := c.GenMove(context.Background(), player)
			if err != nil {
				log.Fatalf("genmove error: %v", err)
			}
			fmt.Printf("OpenGO placed a stone at move %s, next player is %s\n", r.LastMove, r.NextPlayer)
			fmt.Println(r.Board)
		case "pass":
			r, err := c.Pass(context.Background(), player)
			if err != nil {
				log.Fatalf("pass error: %v", err)
			}
			fmt.Printf("Player passed this turn, next player is %s\n", r.NextPlayer)
			fmt.Println(r.Board)
		case "resign":
			r, err := c.Resign(context.Background(), player)
			if err != nil {
				log.Fatalf("resign error: %v", err)
			}
			if r.FinalScore > 0 {
				fmt.Printf("Player resigned this game, final score is B+%.1f\n", r.FinalScore)
			} else {
				fmt.Printf("Player resigned this game, final score is W+%.1f\n", r.FinalScore)
			}
		case "exit", "quit":
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// clear board before free gc
			r, err := c.ClearBoard(context.Background(), player)
			if err != nil {
				log.Fatalf("clear board error: %v", err)
			}
			if r.Code == int32(code.Code_OK) {
				fmt.Println("clear board ok")
			}

			r, err = c.FreeGC(ctx, player)
			if err != nil {
				log.Fatalf("stop error: %v", err)
			}
			if r.Code != int32(code.Code_OK) {
				log.Fatalf("stop error: %s", r.Message)
			}
			fmt.Println(color.GreenString("Bye, player %s", playerID))
			return
		}
	}
}
