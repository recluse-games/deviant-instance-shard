package server

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/collector"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/dispatcher"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (s *server) StartEncounter(stream deviant.EncounterService_StartEncounterServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		responseQueue := make(chan *deviant.EncounterResponse)

		// Submit New Work Request from client to collector
		collector.IncomingCollector(in, responseQueue)

		response := <-responseQueue

		if err := stream.Send(response); err != nil {
			return err
		}
	}
}

func (s *server) UpdateEncounter(stream deviant.EncounterService_UpdateEncounterServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		responseQueue := make(chan *deviant.EncounterResponse)

		// Submit New Work Request from client to collector
		collector.IncomingCollector(in, responseQueue)

		response := <-responseQueue

		if err := stream.Send(response); err != nil {
			return err
		}
	}
}

// Start Starts the Server.
func Start() {
	// Start the message dispatcher.
	fmt.Println("Starting the dispatcher")
	dispatcher.StartDispatcher(10)

	socket := "0.0.0.0:50051"
	fmt.Printf("starting deviant-instance-shard on %v", socket)

	protocol := "tcp"
	lis, err := net.Listen(protocol, socket)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	deviant.RegisterEncounterServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
