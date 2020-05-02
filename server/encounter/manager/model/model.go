package model

import deviant "github.com/recluse-games/deviant-protobuf/genproto"

// DeviantRequestResponse RR LifeCycle Object for Workers
type DeviantRequestResponse struct {
	Request         *deviant.EncounterRequest
	ResponseChannel chan *deviant.EncounterResponse
}
