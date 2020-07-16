package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/recluse-games/deviant-instance-shard/internal/encounter/manager/worker"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var streams = make(map[string]deviant.EncounterService_UpdateEncounterServer)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5,
	PermitWithoutStream: true,
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     2 * time.Hour,
	MaxConnectionAge:      2 * time.Hour,
	MaxConnectionAgeGrace: 60 * time.Second,
	Time:                  60 * time.Second,
	Timeout:               60 * time.Second,
}

type server struct {
}

// UpdateEncounter This is the main server implementation for Deviant it handles all forms of requests in an async duplex fashion utilizing GRPC.
func (s *server) UpdateEncounter(stream deviant.EncounterService_UpdateEncounterServer) error {
	incomingWorker := worker.NewIncomingWorker(1)

	for {
		in, err := stream.Recv()
		streams[in.PlayerId] = stream

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		response := incomingWorker.ProcessWork(in)

		for id, stream := range streams {
			if err := stream.Send(response); err != nil {
				delete(streams, id)
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
