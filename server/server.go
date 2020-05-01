package server

import (
	"fmt"

	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/collector"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/manager/dispatcher"
	deviant "github.com/recluse-games/deviant-protobuf/genproto"
)

// Start Starts the Server.
func Start() {
	// Start the message dispatcher.
	fmt.Println("Starting the dispatcher")
	dispatcher.StartDispatcher(100)

	test := &deviant.EncounterRequest{
		PlayerId: "0000",
		Encounter: &deviant.Encounter{
			Turn: &deviant.Turn{
				Id:    "000",
				Phase: deviant.TurnPhaseNames_PHASE_POINT,
			},
		},
		ActionName: deviant.EntityActionNames_NOTHING,
	}

	for {
		collector.IncomingCollector(test)
	}
}
