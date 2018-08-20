#!/bin/bash

cd protos/

protoc -I ./ --go_out=plugins=grpc:. game.proto
python -m grpc_tools.protoc -I ./ --python_out=. --grpc_python_out=. game.proto

mv game.pb.go ../client/proto
mv game_pb2.py game_pb2_grpc.py ../server/
