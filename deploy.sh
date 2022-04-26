#!/bin/bash
docker build -t grpc-server -f server.Dockerfile .   
docker build -t grpc-client -f client.Dockerfile .  
docker run -d --name grpc-server -p 50051:50051 grpc-server
docker run -d --name grpc-client -p 8080:8080 -e USER_ENDPOINT=grpc-server:50051 --link=grpc-server grpc-client