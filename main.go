package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"./db"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dgraph-io/dgo"
	pb "github.com/kobylyanskiy/dgraph-api/dgraph"
)

const (
	port = ":50051"
)

type server struct{}

var client *dgo.Dgraph

func (s *server) AddAgent(ctx context.Context, in *pb.Agent) (*pb.Result, error) {

	p := db.Agent{
		Codename: "Agent008",
		Rating:   4.4,
		Operations: []db.Operation{{
			Codename: "Operation1",
		}, {
			Codename: "Operation2",
		}},
	}
	assigned, err := db.AddData(client, p)
	log.Println(assigned.Uids["blank-0"])

	resp, err := db.GetData(client, "Agent008")
	log.Println(resp)

	type Root struct {
		Me []db.Response `json:"request"`
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Me: %+v\n", r.Me)

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
	client = db.SetupDgraph()

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
