package play

import (
	"fmt"
	"log"

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

//Play Applys a play action
func Play(encounter *deviant.Encounter, playAction *deviant.EntityPlayAction) bool {

	for _, card := range encounter.ActiveEntity.Hand.Cards {
		if card.Id == playAction.CardId {
			for _, playPair := range playAction.Plays {
				//CAUTION: HACK - This logic should be moved into rules
				if playPair.X >= 0 && playPair.Y >= 0 && playPair.X <= 7 && playPair.Y <= 7 {
					var xycoord = fmt.Sprintf("%d,%d", playPair.X, playPair.Y)
					log.Output(1, "Dealing Damage"+xycoord)
					if encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp < card.Damage {
						encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp = 0
					} else {
						encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp = encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp - card.Damage
					}

					if encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp == 0 {
						encounter.ActiveEntityOrder = removeEntityFromOrder(encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Id, encounter.ActiveEntityOrder)
						encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y] = &deviant.Entity{}
					}
				}
			}

			// Pay Ap Cost for Card
			encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - card.Cost
		}
	}

	return true
}
