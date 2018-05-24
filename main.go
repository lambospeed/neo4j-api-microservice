package main

import (
	"context"
	"log"
	"net"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/kobylyanskiy/dgraph-api/dgraph"
)

func newClient() *dgo.Dgraph {
	d, err := grpc.Dial("dgraph-server-public:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

const (
	port = ":50051"
)

type server struct{}

func (s *server) AddAgent(ctx context.Context, in *pb.Agent) (*pb.Result, error) {
	return &pb.Result{Result: "Hello " + in.Codename}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDgraphServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
