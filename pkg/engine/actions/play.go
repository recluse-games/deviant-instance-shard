package actions

import (
	"math"

	"github.com/google/uuid"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

type gridLocation struct {
	X float64
	Y float64
}

func removeCardFromHand(cardID string, slice []*deviant.Card) []*deviant.Card {
	var newCards = []*deviant.Card{}

	for _, card := range slice {
		if card.InstanceId != cardID {
			newCards = append(newCards, card)
		}
	}

	return newCards
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

//Play Applys a play action
func Play(encounter *deviant.Encounter, playAction *deviant.EntityPlayAction, logger *zap.Logger) bool {
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
		if card.InstanceId == playAction.CardId {
			for _, playPair := range playAction.Plays {

				// CAUTION: HACK - This logic should be moved somewhere else to apply rotations directly to the cards themselves maybe?
				var rotationDegree = convertDirectionToDegree(encounter.ActiveEntity.Rotation)
				var rotatedPlayPair = rotateTilePatterns(activeEntityLocationPoint.X, activeEntityLocationPoint.Y, float64(playPair.X), float64(playPair.Y), rotationDegree)

				var x = int(math.RoundToEven(rotatedPlayPair.X))
				var y = int(math.RoundToEven(rotatedPlayPair.Y))

				//CAUTION: HACK - This logic should be moved into rules
				if x >= 0 && y >= 0 && x <= 7 && y <= 7 {
					switch card.Type {
					case deviant.CardType_ATTACK:
						if encounter.Board.Entities.Entities[x].Entities[y].Hp < card.Damage {
							encounter.Board.Entities.Entities[x].Entities[y].Hp = 0
						} else {
							encounter.Board.Entities.Entities[x].Entities[y].Hp = encounter.Board.Entities.Entities[x].Entities[y].Hp - card.Damage
						}
					case deviant.CardType_HEAL:
						if encounter.Board.Entities.Entities[x].Entities[y].MaxHp <= encounter.Board.Entities.Entities[x].Entities[y].Hp+card.Damage {
							encounter.Board.Entities.Entities[x].Entities[y].Hp = encounter.Board.Entities.Entities[x].Entities[y].MaxHp
						} else {
							encounter.Board.Entities.Entities[x].Entities[y].Hp = encounter.Board.Entities.Entities[x].Entities[y].Hp + card.Damage
						}
					case deviant.CardType_BLOCK:
						wall := &deviant.Entity{
							Id:        uuid.New().String(),
							Name:      "Wall",
							Hp:        1,
							MaxHp:     1,
							Class:     deviant.Classes_WALL,
							State:     deviant.EntityStateNames_IDLE,
							Alignment: deviant.Alignment_NEUTRAL,
						}

						encounter.Board.Entities.Entities[x].Entities[y] = wall
					}

					// HACK - This logic should be moved outside of this method and processed on every turn or something.
					if encounter.Board.Entities.Entities[x].Entities[y].Hp <= 0 {
						encounter.ActiveEntityOrder, _ = removeEntityID(encounter.Board.Entities.Entities[x].Entities[y].Id, encounter.ActiveEntityOrder)

						encounter.Board.Entities.Entities[x].Entities[y] = &deviant.Entity{}
					}
				}
			}

			// Pay Ap Cost for Card
			encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - card.Cost
			// Add Card To Discard Pile
			encounter.ActiveEntity.Discard.Cards = append(encounter.ActiveEntity.Discard.Cards, card)
		}
	}

	// Remove Card From Hand
	encounter.ActiveEntity.Hand.Cards = removeCardFromHand(playAction.CardId, encounter.ActiveEntity.Hand.Cards)

	if logger != nil {
		logger.Debug("Entity Play Processed",
			zap.String("actionID", "Play"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)
	}

	return true
}
