package worker

import (
	"fmt"
	"io/ioutil"

	model "github.com/recluse-games/deviant-instance-shard/server/encounter/manager/model"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/matchmaker"
	deviant "github.com/recluse-games/deviant-protobuf/genproto"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewIncomingWorker creates, and returns a new Worker object.
func NewIncomingWorker(id int, workerQueue chan chan *model.DeviantRequestResponse) IncomingWorker {
	// Create, and return the worker.
	worker := IncomingWorker{
		ID:          id,
		Work:        make(chan *model.DeviantRequestResponse),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

// NewOutgoingWorker creates, and returns a new Worker object.
func NewOutgoingWorker(id int, workerQueue chan chan *deviant.EncounterResponse) OutgoingWorker {
	// Create, and return the worker.
	worker := OutgoingWorker{
		ID:          id,
		Work:        make(chan *deviant.EncounterResponse),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

// IncomingWorker A worker to process work.
type IncomingWorker struct {
	ID          int
	Work        chan *model.DeviantRequestResponse
	WorkerQueue chan chan *model.DeviantRequestResponse
	QuitChan    chan bool
}

// OutgoingWorker A worker to process work.
type OutgoingWorker struct {
	ID          int
	Work        chan *deviant.EncounterResponse
	WorkerQueue chan chan *deviant.EncounterResponse
	QuitChan    chan bool
}

// StartIncoming Starts a working with an infinite loop.
func (w *IncomingWorker) StartIncoming() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				// Implement Rules Engine and Matchingmaking integration here.

				fmt.Printf("Action Processed\n")
				actionResponse := matchmaker.GenerateMatch()

				work.ResponseChannel <- actionResponse

			case <-w.QuitChan:
				// We have been asked to stop.
				fmt.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}

// StartOutgoing Starts a working with an infinite loop.
func (w *OutgoingWorker) StartOutgoing() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				jsonResponse, _ := protojson.Marshal(work)
				ioutil.WriteFile("test", jsonResponse, 0755)
				// Receive a work request.
				fmt.Printf("Recieved outgoing", work)
			case <-w.QuitChan:
				// We have been asked to stop.
				fmt.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}

// StopIncoming tells the worker to stop listening for work requests.
func (w *IncomingWorker) StopIncoming() {
	go func() {
		w.QuitChan <- true
	}()
}

// StopOutgoing tells the worker to stop listening for work requests.
func (w *IncomingWorker) StopOutgoing() {
	go func() {
		w.QuitChan <- true
	}()
}
