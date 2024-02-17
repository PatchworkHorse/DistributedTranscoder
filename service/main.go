package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/PatchworkHorse/DistributedTranscoder/api/rpc/generated"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEchoServiceServer
}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Listening for gRPC on port 8080")
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
