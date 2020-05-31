package play

import (
	"log"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// ValidateApCost Determines that the entity has the correct amount of AP to perform the requested move.
func ValidateApCost(activeEntity *deviant.Entity, requestedPlayAction *deviant.EntityPlayAction, encounter *deviant.Encounter) bool {
	var totalApCost = int32(0)

	log.Output(0, requestedPlayAction.CardId)
	for _, card := range encounter.ActiveEntity.Hand.Cards {
		if card.InstanceId == requestedPlayAction.CardId {
			totalApCost += card.Cost
		}
	}

	if totalApCost <= activeEntity.Ap {
		return true
	}

	return false
}
