package worker

import (
	"fmt"
	"io/ioutil"

	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/collector"
	actions "github.com/recluse-games/deviant-instance-shard/server/rules/processor"
	rules "github.com/recluse-games/deviant-instance-shard/server/rules/processor"
	deviant "github.com/recluse-games/deviant-protobuf/genproto"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewIncomingWorker creates, and returns a new Worker object.
func NewIncomingWorker(id int, workerQueue chan chan *deviant.EncounterRequest) IncomingWorker {
	// Create, and return the worker.
	worker := IncomingWorker{
		ID:          id,
		Work:        make(chan *deviant.EncounterRequest),
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
	Work        chan *deviant.EncounterRequest
	WorkerQueue chan chan *deviant.EncounterRequest
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
				// Receive a work request.
				validatedAction := rules.Process(work.Encounter.Turn.Phase, work.Encounter.ActiveEntity, work.ActionName)

				if validatedAction {
					actionProcessed := actions.Process(work.Encounter.Turn.Phase, work.Encounter.ActiveEntity, work.ActionName)
					if actionProcessed {
						fmt.Printf("Action Processed\n")
						actionResponse := &deviant.EncounterResponse{
							PlayerId:  "0000",
							Encounter: work.Encounter,
						}
						collector.OutgoingCollector(actionResponse)
					}
				}
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
