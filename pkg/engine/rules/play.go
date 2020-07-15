package rules

import (
	"github.com/recluse-games/deviant-instance-shard/pkg/engine/engineutil"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap"
)

// ValidatePlayApCost Determines that the entity has the correct amount of AP to perform the requested move.
func ValidatePlayApCost(activeEntity *deviant.Entity, requestedPlayAction *deviant.EntityPlayAction, encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
	status := false
	totalApCost := int32(0)

	for _, card := range encounter.ActiveEntity.Hand.Cards {
		if card.InstanceId == requestedPlayAction.CardId {
			totalApCost += card.Cost
		}
	}

	if totalApCost <= activeEntity.Ap {
		status = true
	}

	logger.Debug("Validated Play AP Cost",
		"actionID", "ValidatePlayApCost",
		"entityID", encounter.ActiveEntity.Id,
		"status", status,
	)

	return status
}

// ValidateCardInHand Determines if the request card is in the entities hand.
func ValidateCardInHand(activeEntity *deviant.Entity, requestedPlayAction *deviant.EntityPlayAction, encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
	status := false

	for _, card := range encounter.ActiveEntity.Hand.Cards {
		if card.InstanceId == requestedPlayAction.CardId {
			status = true
		}
	}

	logger.Debug("ValidateCardInHand Card is not in hand",
		"actionID", "ValidatePlayApCost",
		"entityID", encounter.ActiveEntity.Id,
		"status", status,
	)

	return status
}

// ValidateCardConstraints Validates specific sub rules for particular types of cards I.E. Block/Attack/Heal.
func ValidateCardConstraints(activeEntity *deviant.Entity, requestedPlayAction *deviant.EntityPlayAction, encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
	activeEntityLocationPoint, err := engineutil.LocateEntity(encounter.ActiveEntity.Id, encounter.Board)
	if err != nil {
		logger.Debug("Validated Card Type Specific Constraints",
			"actionID", "ValidateCardTypeSpecificConstraints",
			"entityID", encounter.ActiveEntity.Id,
			"status", false,
		)
		return false
	}

	for _, card := range encounter.ActiveEntity.Hand.Cards {
		if card.InstanceId == requestedPlayAction.CardId {
			for _, playPair := range requestedPlayAction.Plays {
				location := &engineutil.Location{X: playPair.X, Y: playPair.Y}
				degree := engineutil.GetDegree(encounter.ActiveEntity.Rotation)

				location.Rotate(*activeEntityLocationPoint, degree)

				if location.X >= 0 && location.Y >= 0 && location.X <= 7 && location.Y <= 8 {
					switch card.Type {
					case deviant.CardType_BLOCK:
						if encounter.Board.Entities.Entities[location.X].Entities[location.Y].Id != "" {
							return false
						}
					}
				}
			}
		}
	}

	logger.Debug("Validated Card Type Specific Constraints",
		"actionID", "ValidateCardTypeSpecificConstraints",
		"entityID", encounter.ActiveEntity.Id,
	)

	return true
}
