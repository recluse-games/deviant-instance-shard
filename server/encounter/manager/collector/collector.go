package collector

import (
	model "github.com/recluse-games/deviant-instance-shard/server/encounter/manager/model"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// IncomingWorkQueue manages workers for outgoing messages.
var IncomingWorkQueue = make(chan *model.DeviantRequestResponse, 100)

// OutgoingWorkQueue manages workers for outgoing messages.
var OutgoingWorkQueue = make(chan *deviant.EncounterResponse, 100)

//IncomingCollector Collects incoming messages and Creates Work Queues for Them
func IncomingCollector(encounterRequest *deviant.EncounterRequest, responseChannel chan *deviant.EncounterResponse) {
	// Push the work onto the queue.
	IncomingWorkQueue <- &model.DeviantRequestResponse{Request: encounterRequest, ResponseChannel: responseChannel}

	return
}

//OutgoingCollector Collects outgoing messages and Creates Work Queues for Them
func OutgoingCollector(encounterResponse *deviant.EncounterResponse) {
	// Push the work onto the queue.
	OutgoingWorkQueue <- encounterResponse

	return
}
