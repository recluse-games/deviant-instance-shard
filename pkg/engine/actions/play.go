package actions

import (
	"errors"

	"github.com/google/uuid"
	"github.com/recluse-games/deviant-instance-shard/pkg/engine/engineutil"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

// GetCard Returns a point to a card struct from a card slice based on ID.
func getCard(id string, cards []*deviant.Card) (*deviant.Card, error) {
	for _, card := range cards {
		if card.InstanceId == id {
			return card, nil
		}
	}

	return nil, errors.New("Failed to locate card with provided ID in Hand")
}

// locateCard Returns the index of a card in a slice of cards based on it's instanceID.
func locateCard(instanceID string, cards []*deviant.Card) (*int, error) {
	for i, card := range cards {
		if card.InstanceId == instanceID {
			return &i, nil
		}
	}

	return nil, errors.New("Failed to locate card with provided instanceID in slice")
}

//remove Removes a string from a slice based on index.
func removeCard(s int, slice []*deviant.Card) []*deviant.Card {
	return append(slice[:s], slice[s+1:]...)
}

// getCardLocations Generates a list of locations where a card will be played from a location.
func getCardLocations(card *deviant.Card, start *engineutil.Location) []*engineutil.Location {
	cardActions := map[string]map[string][]*engineutil.Location{}
	cardActions[card.Action.Id] = map[string][]*engineutil.Location{}
	cardActions[card.Action.Id][engineutil.UP.String()] = []*engineutil.Location{}
	cardActions[card.Action.Id][engineutil.DOWN.String()] = []*engineutil.Location{}
	cardActions[card.Action.Id][engineutil.LEFT.String()] = []*engineutil.Location{}
	cardActions[card.Action.Id][engineutil.RIGHT.String()] = []*engineutil.Location{}

	for _, pattern := range card.Action.Pattern {
		offsetStart := engineutil.Location{X: 0, Y: 0}

		for _, offset := range pattern.Offset {
			direction, err := engineutil.ConvertToDirection(offset.Direction)
			if err != nil {
				panic("Invalid direction in card.")
			}
			offsetStart.Move(*direction, offset.Distance)
		}

		switch pattern.Direction {
		case deviant.Direction_UP:
			for i := int32(0); i < pattern.Distance; i++ {
				clone := start.Clone()
				clone.Add(offsetStart)

				cardActions[card.Action.Id][engineutil.UP.String()] = append(cardActions[card.Action.Id][engineutil.UP.String()], clone)
				offsetStart.Add(engineutil.UP.Location())
			}
		case deviant.Direction_DOWN:
			for i := int32(0); i < pattern.Distance; i++ {
				clone := start.Clone()
				clone.Add(offsetStart)

				cardActions[card.Action.Id][engineutil.DOWN.String()] = append(cardActions[card.Action.Id][engineutil.DOWN.String()], clone)
				offsetStart.Add(engineutil.DOWN.Location())
			}
		case deviant.Direction_LEFT:
			for i := int32(0); i < pattern.Distance; i++ {
				clone := start.Clone()
				clone.Add(offsetStart)

				cardActions[card.Action.Id][engineutil.LEFT.String()] = append(cardActions[card.Action.Id][engineutil.LEFT.String()], clone)
				offsetStart.Add(engineutil.LEFT.Location())
			}
		case deviant.Direction_RIGHT:
			for i := int32(0); i < pattern.Distance; i++ {
				clone := start.Clone()
				clone.Add(offsetStart)

				cardActions[card.Action.Id][engineutil.RIGHT.String()] = append(cardActions[card.Action.Id][engineutil.RIGHT.String()], clone)
				offsetStart.Add(engineutil.RIGHT.Location())
			}
		}
	}

	playLocations := []*engineutil.Location{}

	for _, direction := range cardActions {
		for _, locations := range direction {
			playLocations = append(playLocations, locations...)
		}
	}

	return playLocations
}

func processAttack(damage int32, target engineutil.Location, board *deviant.Board) {
	if board.Entities.Entities[target.X].Entities[target.Y].Hp < damage {
		board.Entities.Entities[target.X].Entities[target.Y].Hp = 0
	} else {
		board.Entities.Entities[target.X].Entities[target.Y].Hp = board.Entities.Entities[target.X].Entities[target.Y].Hp - damage
	}
}

func processHeal(healing int32, target engineutil.Location, board *deviant.Board) {
	if board.Entities.Entities[target.X].Entities[target.Y].MaxHp <= board.Entities.Entities[target.X].Entities[target.Y].Hp+healing {
		board.Entities.Entities[target.X].Entities[target.Y].Hp = board.Entities.Entities[target.X].Entities[target.Y].MaxHp
	} else {
		board.Entities.Entities[target.X].Entities[target.Y].Hp = board.Entities.Entities[target.X].Entities[target.Y].Hp + healing
	}
}

func processBlock(hp int32, target engineutil.Location, board *deviant.Board) {
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

	board.Entities.Entities[target.X].Entities[target.Y] = wall
}

// Play Processes a play.
func Play(encounter *deviant.Encounter, playAction *deviant.EntityPlayAction, logger *zap.Logger) bool {
	card, err := getCard(playAction.CardId, encounter.ActiveEntity.Hand.Cards)
	if err != nil {
		if logger != nil {
			logger.Error("Card does not exist in hand.",
				zap.String("actionID", "Play"),
				zap.String("entityID", encounter.ActiveEntity.Id),
			)
		}

		return false
	}

	entityLocation, err := engineutil.LocateEntity(encounter.ActiveEntity.Id, encounter.Board)
	if err != nil {
		if logger != nil {
			logger.Error("Unable to locate entity",
				zap.String("actionID", "Play"),
				zap.String("entityID", encounter.ActiveEntity.Id),
			)
		}

		return false
	}

	cardLocations := getCardLocations(card, entityLocation)
	engineutil.RotateLocations(*entityLocation, cardLocations, encounter.ActiveEntity.Rotation)

	for _, loc := range cardLocations {
		// This logic should maybe be migrated into another function, I'm not sure if I like conditionals this long here... alternatively this should be removed from the list when it's generated one time.
		if loc.X <= int32(len(encounter.Board.Entities.Entities)-1) && 0 <= loc.X && loc.Y <= int32(len(encounter.Board.Entities.Entities[loc.X].Entities)-1) && 0 <= loc.Y {
			switch card.Type {
			case deviant.CardType_ATTACK:
				processAttack(card.Damage, *loc, encounter.Board)
			case deviant.CardType_HEAL:
				processHeal(card.Damage, *loc, encounter.Board)
			case deviant.CardType_BLOCK:
				processBlock(card.Damage, *loc, encounter.Board)
			}

			// HACK - This logic should be moved outside of this method and processed on every turn or something.
			if encounter.Board.Entities.Entities[loc.X].Entities[loc.Y].Hp <= 0 {
				encounter.ActiveEntityOrder, _ = engineutil.RemoveEntityID(encounter.Board.Entities.Entities[loc.X].Entities[loc.Y].Id, encounter.ActiveEntityOrder)

				encounter.Board.Entities.Entities[loc.X].Entities[loc.Y] = &deviant.Entity{}
			}
		}
	}

	encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - card.Cost
	encounter.ActiveEntity.Discard.Cards = append(encounter.ActiveEntity.Discard.Cards, card)

	cardHandIndex, err := locateCard(card.InstanceId, encounter.ActiveEntity.Hand.Cards)
	if err != nil {
		if logger != nil {
			logger.Error("Unable to get locate card in hand",
				zap.String("actionID", "Play"),
				zap.String("entityID", encounter.ActiveEntity.Id),
			)
		}

		return false
	}
	encounter.ActiveEntity.Hand.Cards = removeCard(*cardHandIndex, encounter.ActiveEntity.Hand.Cards)

	if logger != nil {
		logger.Debug("Entity Play Processed",
			zap.String("actionID", "Play"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)
	}

	return true
}
