package main

import (
	"log"
	"net"

	"github.com/huyntsgs/grpc-bi-directional-stream/api"
	"google.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()
	protobuf.RegisterMathServer(grpcServer, &MathService{})

	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	log.Println("Listening grpc on 8888")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Error start grpc server: %v", err)
	}

}
