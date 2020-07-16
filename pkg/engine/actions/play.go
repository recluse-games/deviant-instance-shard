package actions

import (
	"github.com/google/uuid"
	"github.com/recluse-games/deviant-instance-shard/pkg/engine/engineutil"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap"
)

func processAttack(damage int32, play *deviant.Play, board *deviant.Board) {
	if board.Entities.Entities[play.X].Entities[play.Y].Hp < damage {
		board.Entities.Entities[play.X].Entities[play.Y].Hp = 0
	} else {
		board.Entities.Entities[play.X].Entities[play.Y].Hp = board.Entities.Entities[play.X].Entities[play.Y].Hp - damage
	}
}

func processHeal(healing int32, play *deviant.Play, board *deviant.Board) {
	if board.Entities.Entities[play.X].Entities[play.Y].MaxHp <= board.Entities.Entities[play.X].Entities[play.Y].Hp+healing {
		board.Entities.Entities[play.X].Entities[play.Y].Hp = board.Entities.Entities[play.X].Entities[play.Y].MaxHp
	} else {
		board.Entities.Entities[play.X].Entities[play.Y].Hp = board.Entities.Entities[play.X].Entities[play.Y].Hp + healing
	}
}

func processBlock(hp int32, play *deviant.Play, board *deviant.Board) {
	// WARNING: This should be changed to something more specific coming from the card definition itself, get creative.
	wall := &deviant.Entity{
		Id:        uuid.New().String(),
		Name:      "Wall",
		Hp:        hp,
		MaxHp:     hp,
		Class:     deviant.Classes_WALL,
		State:     deviant.EntityStateNames_IDLE,
		Alignment: deviant.Alignment_NEUTRAL,
	}

	board.Entities.Entities[play.X].Entities[play.Y] = wall
}

// Play Processes a play.
func Play(encounter *deviant.Encounter, playAction *deviant.EntityPlayAction, logger *zap.SugaredLogger) bool {
	card, err := engineutil.GetCard(playAction.CardId, encounter.ActiveEntity.Hand.Cards)
	if err != nil {
		logger.Error("Card does not exist in hand.",
			zap.String("actionID", "Play"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)

		return false
	}

	entityLocation, err := engineutil.LocateEntity(encounter.ActiveEntity.Id, encounter.Board)
	if err != nil {
		logger.Error("Unable to locate entity",
			zap.String("actionID", "Play"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)

		return false
	}

	cardPlays := engineutil.GetCardPlays(card, entityLocation, encounter.Board)
	engineutil.RotatePlays(*entityLocation, *cardPlays, encounter.ActiveEntity.Rotation)

	for _, play := range *cardPlays {
		// This logic should maybe be migrated into another function, I'm not sure if I like conditionals this long here... alternatively this should be removed from the list when it's generated one time.
		if play.X <= int32(len(encounter.Board.Entities.Entities)-1) && 0 <= play.X && play.Y <= int32(len(encounter.Board.Entities.Entities[play.X].Entities)-1) && 0 <= play.Y {
			switch card.Type {
			case deviant.CardType_ATTACK:
				processAttack(card.Damage, play, encounter.Board)
			case deviant.CardType_HEAL:
				processHeal(card.Damage, play, encounter.Board)
			case deviant.CardType_BLOCK:
				processBlock(card.Damage, play, encounter.Board)
			}

			// HACK - This logic should be moved outside of this method and processed on every turn or something.
			if encounter.Board.Entities.Entities[play.X].Entities[play.Y].Hp <= 0 {
				encounter.ActiveEntityOrder, _ = engineutil.RemoveString(encounter.Board.Entities.Entities[play.X].Entities[play.Y].Id, encounter.ActiveEntityOrder)

				encounter.Board.Entities.Entities[play.X].Entities[play.Y] = &deviant.Entity{}
			}
		}
	}

	encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - card.Cost
	encounter.ActiveEntity.Discard.Cards = append(encounter.ActiveEntity.Discard.Cards, card)

	cardHandIndex, err := engineutil.LocateCard(card.InstanceId, encounter.ActiveEntity.Hand.Cards)
	if err != nil {
		logger.Error("Unable to get locate card in hand",
			zap.String("actionID", "Play"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)

		return false
	}
	encounter.ActiveEntity.Hand.Cards = engineutil.RemoveCard(*cardHandIndex, encounter.ActiveEntity.Hand.Cards)

	logger.Debug("Entity Play Processed",
		zap.String("actionID", "Play"),
		zap.String("entityID", encounter.ActiveEntity.Id),
	)

	return true
}
