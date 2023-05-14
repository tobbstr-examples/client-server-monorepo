package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/tobbstr-examples/client-server-monorepo/shared/pb/v1"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 1997, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedAuditerServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) LastAccess(ctx context.Context, in *pb.LastAccessRequest) (*pb.LastAccessResponse, error) {
	log.Println("Received request")
	return &pb.LastAccessResponse{Person: &pb.Person{Name: "John Doe", Age: 65}}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuditerServer(s, &server{})
	log.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
