package main

import (
	"context"
	"log"
	"net"

	pb "todolist/protobuf/todo"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type RegisterTodolistServer struct {
	pb.UnimplementedTodolistServer
}

func (s *RegisterTodolistServer) GetFirstItem(ctx context.Context, input *pb.FirstItemRequest) (*pb.FirstItemResponse, error) {
	log.Printf("Requesting first item with no input")
	return &pb.FirstItemResponse{Todo: &pb.Todo{Id: 1, Todo: "Learn Go, gRPC, and Docker!", Completed: true}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodolistServer(s, &RegisterTodolistServer{})
	log.Printf("Server listening at : %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
