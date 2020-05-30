package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/worker"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var streams = make(map[string]deviant.EncounterService_UpdateEncounterServer)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             0,    // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true, // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     2 * time.Hour,    // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      2 * time.Hour,    // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 60 * time.Second, // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  25 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               25 * time.Second, // Wait 1 second for the ping ack before assuming the connection is dead
}

type server struct {
}

func (s *server) UpdateEncounter(stream deviant.EncounterService_UpdateEncounterServer) error {
	incomingWorker := worker.NewIncomingWorker(1)

	for {
		in, err := stream.Recv()
		log.Output(0, in.String())
		streams[in.PlayerId] = stream

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		response := incomingWorker.ProcessWork(in)

		for _, stream := range streams {
			if err := stream.Send(response); err != nil {
				return err
			}
		}
	}
}

// Start Starts the Server.
func Start() {

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
