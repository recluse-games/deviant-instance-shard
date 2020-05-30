package collector

import (
	model "github.com/recluse-games/deviant-instance-shard/server/encounter/manager/model"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// IncomingWorkQueue manages workers for outgoing messages.
var IncomingWorkQueue = make(chan *model.DeviantRequestResponse, 1000)

//IncomingCollector Collects incoming messages and Creates Work Queues for Them
func IncomingCollector(encounterRequest *deviant.EncounterRequest, responseChannel chan *deviant.EncounterResponse) {
	IncomingWorkQueue <- &model.DeviantRequestResponse{Request: encounterRequest, ResponseChannel: responseChannel}
	return
}
