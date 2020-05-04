package dispatcher

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/collector"
	model "github.com/recluse-games/deviant-instance-shard/server/encounter/manager/model"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/worker"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// WorkerQueue holds all work requests
var IncomingWorkerQueue chan chan *model.DeviantRequestResponse
var OutgoingWorkerQueue chan chan *deviant.EncounterResponse

// StartDispatcher Dispatches to the worker queues.
func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	IncomingWorkerQueue = make(chan chan *model.DeviantRequestResponse, nworkers)
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
				glog.Info("Received incoming work request")
				go func() {
					worker := <-IncomingWorkerQueue

					glog.Info("Dispatching work request")
					worker <- incomingWork
				}()
			case outgoingWork := <-collector.OutgoingWorkQueue:
				glog.Info("Received incoming work request")
				go func() {
					worker := <-OutgoingWorkerQueue

					glog.Info("Dispatching work request")
					worker <- outgoingWork
				}()
			}
		}
	}()
}
