package main

import (
	"fmt"
	"gRPC-GO-learn/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("gRPC Server")
	lis, err := net.Listen("tcp", "0.0.0.0:9091")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	service := greetpb.UnimplementedGreetServiceServer{}
	grpcServer := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(grpcServer, service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen %v", err)
	}

}
