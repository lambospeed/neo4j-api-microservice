package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/kobylyanskiy/dgraph-api/dgraph"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) AddAgent(ctx context.Context, in *pb.Agent) (*pb.Result, error) {
	return &pb.Result{Result: true, ErrorMessage: ""}, nil
}

func (s *server) AddOperation(ctx context.Context, in *pb.OperationParticipants) (*pb.Result, error) {
	return &pb.Result{Result: true, ErrorMessage: ""}, nil
}

func (s *server) GetOperations(ctx context.Context, in *pb.Agent) (*pb.GetOperationsResult, error) {
	return &pb.GetOperationsResult{
		Result: &pb.Result{
			Result:       true,
			ErrorMessage: "",
		},
		Operations: []*pb.Operation{
			&pb.Operation{Codename: "Codename"},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDgraphServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
