package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/collector"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/dispatcher"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/matchmaker"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	server struct{}

	// Map client connections to a potential pool
	ctxMap struct {
		streams map[string][]*deviant.EncounterService_FindEncounterServer
		m       sync.Mutex
	}
)

var (
	cm       = ctxMap{streams: make(map[string][]*deviant.EncounterService_FindEncounterServer)}
	maxUsers = 2
	mm       = matchmaker.NewEngine(matchmaker.EngineOptions{MaxUsers: maxUsers, WaitPeriod: time.Duration(time.Second * 10)})
	wg       sync.WaitGroup
)

func (s *server) FindEncounter(req *deviant.FindEncounterRequest, stream deviant.EncounterService_FindEncounterServer) error {
	p := mm.JoinPool(req.GetPlayerID())

	wg.Add(1)
	go func() {
		select {
		case pool := <-p:
			if !pool.IsFull {
				cm.m.Lock()
				// Add the context to the map until the pool is full or timed out
				cm.streams[pool.PoolID] = append(cm.streams[pool.PoolID], &stream)
				cm.m.Unlock()
			}
			if pool.IsFull {
				log.Printf("Filled pool: %v, with users: %v", pool.PoolID, pool.Users)
				cm.m.Lock()
				cm.streams[pool.PoolID] = append(cm.streams[pool.PoolID], &stream)

				// Push a message to all members of the full pool
				for _, st := range cm.streams[pool.PoolID] {
					res := &deviant.FindEncounterResponse{PoolID: pool.PoolID}
					stream := *st
					if err := stream.Send(res); err != nil {
						// Delete the pool from the map once full
						delete(cm.streams, pool.PoolID)
					}
				}
				cm.m.Unlock()
			}
			if pool.TimedOut {
				log.Printf("Timed out attempting to fill pool: %v | Users in pool: %v", pool.PoolID, pool.Users)
				// Push messages to all members of the full pool
				for _, st := range cm.streams[pool.PoolID] {
					res := &deviant.FindEncounterResponse{PoolID: "null"}
					stream := *st
					if err := stream.Send(res); err != nil {
						// Delete the timed out pool
						delete(cm.streams, pool.PoolID)
					}
				}
			}
		}
	}()
	wg.Wait()

	return nil
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
	fmt.Printf("starting deviant-instance-shard on %v\n", socket)

	protocol := "tcp"
	lis, err := net.Listen(protocol, socket)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	deviant.RegisterEncounterServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
