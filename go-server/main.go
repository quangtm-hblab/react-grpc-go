package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/quangtm-hblab/react-grpc-go/protobuf"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Receiverd : %v", in.GetName())
	return &pb.HelloReply{Message: "Hello" + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("err when create lis: %v", err)

	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)

	}
}
