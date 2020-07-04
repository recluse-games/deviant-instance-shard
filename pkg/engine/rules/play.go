package rules

import (
	"math"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

type gridLocation struct {
	X float64
	Y float64
}

func convertDirectionToDegree(characterRotation deviant.EntityRotationNames) float64 {
	switch characterRotation {
	case deviant.EntityRotationNames_NORTH:
		return 180.00
	case deviant.EntityRotationNames_SOUTH:
		return 0.00
	case deviant.EntityRotationNames_EAST:
		return 270.00
	case deviant.EntityRotationNames_WEST:
		return 90.00
	}

	return 0.00
}

func rotateTilePatterns(ocx float64, ocy float64, px float64, py float64, rotationAngle float64) *gridLocation {
	var radians = (math.Pi / 180) * rotationAngle
	var s = math.Sin(radians)
	var c = math.Cos(radians)

	// translate point back to origin:
	px -= ocx
	py -= ocy

	// rotate point
	var xnew = px*c - py*s
	var ynew = px*s + py*c

	// translate point back:
	px = xnew + ocx
	py = ynew + ocy

	return &gridLocation{px, py}
}

// ValidateApCost Determines that the entity has the correct amount of AP to perform the requested move.
func ValidatePlayApCost(activeEntity *deviant.Entity, requestedPlayAction *deviant.EntityPlayAction, encounter *deviant.Encounter) bool {
	var totalApCost = int32(0)

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

// ValidateApCost Validates specific sub rules for particular types of cards I.E. Block/Attack/Heal.
func ValidateCardTypeSpecificConstraints(activeEntity *deviant.Entity, requestedPlayAction *deviant.EntityPlayAction, encounter *deviant.Encounter) bool {
	var activeEntityLocationPoint = &gridLocation{}

	for y, entitiesRow := range encounter.Board.Entities.Entities {
		for x, entity := range entitiesRow.Entities {
			if entity.Id == encounter.ActiveEntity.Id {
				activeEntityLocationPoint.X = float64(y)
				activeEntityLocationPoint.Y = float64(x)
			}
		}
	}
	for _, card := range encounter.ActiveEntity.Hand.Cards {
		if card.InstanceId == requestedPlayAction.CardId {
			for _, playPair := range requestedPlayAction.Plays {

				// CAUTION: HACK - This logic should be moved somewhere else to apply rotations directly to the cards themselves maybe?
				var rotationDegree = convertDirectionToDegree(encounter.ActiveEntity.Rotation)
				var rotatedPlayPair = rotateTilePatterns(activeEntityLocationPoint.X, activeEntityLocationPoint.Y, float64(playPair.X), float64(playPair.Y), rotationDegree)

				var x = int(math.RoundToEven(rotatedPlayPair.X))
				var y = int(math.RoundToEven(rotatedPlayPair.Y))

				//CAUTION: HACK - This logic should be moved into rules
				if x >= 0 && y >= 0 && x <= 7 && y <= 8 {
					switch card.Type {
					case deviant.CardType_BLOCK:
						if encounter.Board.Entities.Entities[x].Entities[y].Id != "" {
							return false
						}
					}
				}
			}
		}
	}

	return true
}
