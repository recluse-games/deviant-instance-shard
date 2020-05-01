package dispatcher

import (
	"fmt"

	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/collector"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/worker"
	deviant "github.com/recluse-games/deviant-protobuf/genproto"
)

// WorkerQueue holds all work requests
var IncomingWorkerQueue chan chan *deviant.EncounterRequest
var OutgoingWorkerQueue chan chan *deviant.EncounterResponse

// StartDispatcher Dispatches to the worker queues.
func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	IncomingWorkerQueue = make(chan chan *deviant.EncounterRequest, nworkers)
	OutgoingWorkerQueue = make(chan chan *deviant.EncounterResponse, nworkers)

	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		incomingWorker := worker.NewIncomingWorker(i+1, IncomingWorkerQueue)
		outgoingWorker := worker.NewOutgoingWorker(i+1, OutgoingWorkerQueue)
		incomingWorker.StartIncoming()
		outgoingWorker.StartOutgoing()
	}

	go func() {
		for {
			select {
			case incomingWork := <-collector.IncomingWorkQueue:
				fmt.Println("Received incoming work requeust")
				go func() {
					worker := <-IncomingWorkerQueue

					fmt.Println("Dispatching work request")
					worker <- incomingWork
				}()
			case outgoingWork := <-collector.OutgoingWorkQueue:
				fmt.Println("Received outgoing work requeust")
				go func() {
					worker := <-OutgoingWorkerQueue

					fmt.Println("Dispatching work request")
					worker <- outgoingWork
				}()
			}
		}
	}()
}
