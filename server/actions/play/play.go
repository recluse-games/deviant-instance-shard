package play

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func removeEntityFromOrder(entityID string, slice []string) []string {
	var entityIDIndex = find(slice, entityID)

	log.Output(1, fmt.Sprintf("%d", entityIDIndex))

	if len(slice) > entityIDIndex+1 {
		return slice[:entityIDIndex+copy(slice[entityIDIndex:], slice[entityIDIndex+1:])]
	} else if len(slice) == entityIDIndex {
		return slice[:entityIDIndex+copy(slice[entityIDIndex:], slice[entityIDIndex-1:])]
	}
	if 0 > entityIDIndex-1 {
		return slice[:0+copy(slice[(entityIDIndex-1):], slice[entityIDIndex:])]
	} else if len(slice) > 1 {
		return slice[:len(slice)-1]
	} else {
		return []string{}
	}
}

func removeCardFromHand(cardID string, slice []*deviant.Card) []*deviant.Card {
	var newCards = []*deviant.Card{}

	for _, card := range slice {
		log.Output(0, card.InstanceId)
		log.Output(0, cardID)
		if card.InstanceId != cardID {
			newCards = append(newCards, card)
		}
	}

	return newCards
}

//Play Applys a play action
func Play(encounter *deviant.Encounter, playAction *deviant.EntityPlayAction) bool {

	for _, card := range encounter.ActiveEntity.Hand.Cards {
		if card.InstanceId == playAction.CardId {
			for _, playPair := range playAction.Plays {
				//CAUTION: HACK - This logic should be moved into rules
				if playPair.X >= 0 && playPair.Y >= 0 && playPair.X <= 7 && playPair.Y <= 8 {
					switch card.Type {
					case deviant.CardType_ATTACK:
						if encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp < card.Damage {
							encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp = 0
						} else {
							encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp = encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp - card.Damage
						}
					case deviant.CardType_HEAL:
						if encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].MaxHp <= encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp+card.Damage {
							encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp = encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].MaxHp
						} else {
							encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp = encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp + card.Damage
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

						encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y] = wall
					}

					// HACK - This logic should be moved outside of this method and processed on every turn or something.
					if encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp == 0 {
						encounter.ActiveEntityOrder = removeEntityFromOrder(encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Id, encounter.ActiveEntityOrder)
						encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y] = &deviant.Entity{}
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

	return true
}
