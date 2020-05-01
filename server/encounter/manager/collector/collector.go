package collector

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto"
)

// IncomingWorkQueue manages workers for outgoing messages.
var IncomingWorkQueue = make(chan *deviant.EncounterRequest, 100)

// OutgoingWorkQueue manages workers for outgoing messages.
var OutgoingWorkQueue = make(chan *deviant.EncounterResponse, 100)

//IncomingCollector Collects incoming messages and Creates Work Queues for Them
func IncomingCollector(encounterRequest *deviant.EncounterRequest) {
	// Push the work onto the queue.
	IncomingWorkQueue <- encounterRequest

	return
}

//OutgoingCollector Collects outgoing messages and Creates Work Queues for Them
func OutgoingCollector(encounterResponse *deviant.EncounterResponse) {
	// Push the work onto the queue.
	OutgoingWorkQueue <- encounterResponse

	return
}
