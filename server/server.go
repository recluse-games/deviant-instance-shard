package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/collector"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/dispatcher"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             0,    // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true, // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     30 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      60 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

var streams = make(map[uuid.UUID]deviant.EncounterService_UpdateEncounterServer)
var sharedResponseQueue = make(chan *deviant.EncounterResponse)

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
		streamUUID, _ := uuid.NewRandom()
		streams[streamUUID] = stream

		in, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// Submit New Work Request from client to collector
		collector.IncomingCollector(in, sharedResponseQueue)

		for response := range sharedResponseQueue {

			for _, clientStream := range streams {
				if err := clientStream.Send(response); err != nil {
					log.Fatalf("%v", err)
				}
			}
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

	s := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep),
		grpc.KeepaliveParams(kasp))
	deviant.RegisterEncounterServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
