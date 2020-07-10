package rules

import (
	"github.com/recluse-games/deviant-instance-shard/pkg/engine/engineutil"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

// ValidatePlayApCost Determines that the entity has the correct amount of AP to perform the requested move.
func ValidatePlayApCost(activeEntity *deviant.Entity, requestedPlayAction *deviant.EntityPlayAction, encounter *deviant.Encounter, logger *zap.Logger) bool {
	var totalApCost = int32(0)

	for _, card := range encounter.ActiveEntity.Hand.Cards {
		if card.InstanceId == requestedPlayAction.CardId {
			totalApCost += card.Cost
		}
	}

	if totalApCost <= activeEntity.Ap {
		return true
	}

	if logger != nil {
		logger.Debug("Validated Play AP Cost",
			zap.String("actionID", "ValidatePlayApCost"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)
	}

	return false
}

// ValidateCardInHand Determines if the request card is in the entities hand.
func ValidateCardInHand(activeEntity *deviant.Entity, requestedPlayAction *deviant.EntityPlayAction, encounter *deviant.Encounter, logger *zap.Logger) bool {
	for _, card := range encounter.ActiveEntity.Hand.Cards {
		if card.InstanceId == requestedPlayAction.CardId {
			return true
		}
	}

	if logger != nil {
		logger.Debug("ValidateCardInHand Card is not in hand",
			zap.String("actionID", "ValidatePlayApCost"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)
	}

	return false
}

// ValidateCardConstraints Validates specific sub rules for particular types of cards I.E. Block/Attack/Heal.
func ValidateCardConstraints(activeEntity *deviant.Entity, requestedPlayAction *deviant.EntityPlayAction, encounter *deviant.Encounter, logger *zap.Logger) bool {
	activeEntityLocationPoint, err := engineutil.LocateEntity(encounter.ActiveEntity.Id, encounter.Board)
	if err != nil {
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

	if logger != nil {
		logger.Debug("Validated Card Type Specific Constraints",
			zap.String("actionID", "ValidateCardTypeSpecificConstraints"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)
	}

	return true
}
