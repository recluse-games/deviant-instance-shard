package dispatcher

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/collector"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/worker"
)

// StartDispatcher Dispatches to the worker queues.
func StartDispatcher(nworkers int) {
	fmt.Println("Starting worker", 1)
	incomingWorker := worker.NewIncomingWorker(1)

	go func() {
		for {
			select {
			case incomingWork := <-collector.IncomingWorkQueue:
				incomingWorker.ProcessWork(incomingWork)

				glog.Info("Dispatching work request")
			}
		}
	}()
}
