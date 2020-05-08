package play

import (
	"fmt"
	"log"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

//Play Applys a play action
func Play(encounter *deviant.Encounter, playAction *deviant.EntityPlayAction) bool {

	for _, card := range encounter.ActiveEntity.Hand.Cards {
		if card.Id == playAction.CardId {
			for _, playPair := range playAction.Plays {
				//CAUTION: HACK - This logic should be moved into rules
				if playPair.X >= 0 && playPair.Y >= 0 && playPair.X <= 7 && playPair.Y <= 7 {
					log.Output(1, "Dealing Damage")
					encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp = encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y].Hp - card.Damage

					var test = fmt.Sprintf("%v", encounter.Board.Entities.Entities[playPair.X].Entities[playPair.Y])

					log.Output(1, test)
				}
			}
		}
	}

	return true
}
