package main

import (
	"log"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "github.com/sambaiz/try-gRPC/protos"

)

const (
	address     = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	r, err := c.RetEcho(context.Background(), &pb.EchoRequest{Say: "hello"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Return: %s", r.Ret)
}
