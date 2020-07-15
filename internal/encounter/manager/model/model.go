package model

import deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"

// DeviantRequestResponse RR LifeCycle Object for Workers
type DeviantRequestResponse struct {
	Request         *deviant.EncounterRequest
	ResponseChannel chan *deviant.EncounterResponse
}
