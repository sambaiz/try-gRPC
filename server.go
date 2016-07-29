package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"log"
	pb "github.com/sambaiz/try-gRPC/protos"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) RetEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	return &pb.EchoReply{Ret: in.Say}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	log.Printf("server start localhost%s", port)
	s.Serve(lis)
}
