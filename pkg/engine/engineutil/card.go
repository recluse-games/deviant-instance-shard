package engineutil

import (
	"errors"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
)

// GetCard Returns a point to a card struct from a card slice based on ID.
func GetCard(id string, cards []*deviant.Card) (*deviant.Card, error) {
	for _, card := range cards {
		if card.InstanceId == id {
			return card, nil
		}
	}

	return nil, errors.New("Failed to locate card with provided ID in Hand")
}

// LocateCard Returns the index of a card in a slice of cards based on it's instanceID.
func LocateCard(instanceID string, cards []*deviant.Card) (*int, error) {
	for i, card := range cards {
		if card.InstanceId == instanceID {
			return &i, nil
		}
	}

	return nil, errors.New("Failed to locate card with provided instanceID in slice")
}

//RemoveCard Removes a card from a slice based on index.
func RemoveCard(s int, slice []*deviant.Card) []*deviant.Card {
	return append(slice[:s], slice[s+1:]...)
}

//actionToID Converts an action type to an ID of a tile.
func actionToID(status deviant.CardActionStatusTypes) string {
	switch status {
	case deviant.CardActionStatusTypes_EMPTY:
		return "0000"
	case deviant.CardActionStatusTypes_BLOCKED:
		return "0001"
	case deviant.CardActionStatusTypes_HIT:
		return "0002"
	}

	return "0000"
}

func traverseTargets(action *deviant.CardAction, tiles *[]*deviant.Tile, current *Location) {
	for _, up := range action.Up {
		clone := current.Clone()
		clone.Add(UP.Location())

		tile := &deviant.Tile{
			Id: actionToID(action.Status),
			X:  clone.X,
			Y:  clone.Y,
		}

		(*tiles) = append((*tiles), tile)
		traverseTargets(up, tiles, clone)
	}

	for _, down := range action.Down {
		clone := current.Clone()
		clone.Add(DOWN.Location())
		tile := &deviant.Tile{
			Id: actionToID(action.Status),
			X:  clone.X,
			Y:  clone.Y,
		}

		(*tiles) = append((*tiles), tile)
		traverseTargets(down, tiles, clone)
	}

	for _, left := range action.Left {
		clone := current.Clone()
		clone.Add(LEFT.Location())
		tile := &deviant.Tile{
			Id: actionToID(action.Status),
			X:  clone.X,
			Y:  clone.Y,
		}

		(*tiles) = append((*tiles), tile)
		traverseTargets(left, tiles, clone)
	}

	for _, right := range action.Right {
		clone := current.Clone()
		clone.Add(RIGHT.Location())
		tile := &deviant.Tile{
			Id: actionToID(action.Status),
			X:  clone.X,
			Y:  clone.Y,
		}

		(*tiles) = append((*tiles), tile)
		traverseTargets(right, tiles, clone)
	}

	for _, upLeft := range action.UpLeft {
		clone := current.Clone()
		clone.Add(UPLEFT.Location())
		tile := &deviant.Tile{
			Id: actionToID(action.Status),
			X:  clone.X,
			Y:  clone.Y,
		}

		(*tiles) = append((*tiles), tile)
		traverseTargets(upLeft, tiles, clone)
	}

	for _, upRight := range action.UpRight {
		clone := current.Clone()
		clone.Add(UPRIGHT.Location())
		tile := &deviant.Tile{
			Id: actionToID(action.Status),
			X:  clone.X,
			Y:  clone.Y,
		}

		(*tiles) = append((*tiles), tile)
		traverseTargets(upRight, tiles, clone)
	}

	for _, downLeft := range action.DownLeft {
		clone := current.Clone()
		clone.Add(DOWNLEFT.Location())
		tile := &deviant.Tile{
			Id: actionToID(action.Status),
			X:  clone.X,
			Y:  clone.Y,
		}

		(*tiles) = append((*tiles), tile)
		traverseTargets(downLeft, tiles, clone)
	}

	for _, downRight := range action.DownRight {
		clone := current.Clone()
		clone.Add(DOWNRIGHT.Location())
		tile := &deviant.Tile{
			Id: actionToID(action.Status),
			X:  clone.X,
			Y:  clone.Y,
		}

		(*tiles) = append((*tiles), tile)
		traverseTargets(downRight, tiles, clone)
	}
}

func traversePlays(action *deviant.CardAction, plays *[]*deviant.Play, current *Location) {
	for _, up := range action.Up {
		clone := current.Clone()
		clone.Add(UP.Location())
		play := &deviant.Play{
			X:    clone.X,
			Y:    clone.Y,
			Type: action.Type,
		}

		// Only append this to plays if it's not blocked and keep traversing.
		if action.Status != deviant.CardActionStatusTypes_BLOCKED {
			(*plays) = append((*plays), play)
			traversePlays(up, plays, clone)
		}

		break
	}

	for _, down := range action.Down {
		clone := current.Clone()
		clone.Add(DOWN.Location())
		play := &deviant.Play{
			X:    clone.X,
			Y:    clone.Y,
			Type: action.Type,
		}

		if action.Status != deviant.CardActionStatusTypes_BLOCKED {
			(*plays) = append((*plays), play)
			traversePlays(down, plays, clone)
		}

		break
	}

	for _, left := range action.Left {
		clone := current.Clone()
		clone.Add(LEFT.Location())
		play := &deviant.Play{
			X:    clone.X,
			Y:    clone.Y,
			Type: action.Type,
		}

		if action.Status != deviant.CardActionStatusTypes_BLOCKED {
			(*plays) = append((*plays), play)
			traversePlays(left, plays, clone)
		}

		break
	}

	for _, right := range action.Right {
		clone := current.Clone()
		clone.Add(RIGHT.Location())
		play := &deviant.Play{
			X:    clone.X,
			Y:    clone.Y,
			Type: action.Type,
		}
		if action.Status != deviant.CardActionStatusTypes_BLOCKED {
			(*plays) = append((*plays), play)
			traversePlays(right, plays, clone)
		}

		break
	}

	for _, upLeft := range action.UpLeft {
		clone := current.Clone()
		clone.Add(UPLEFT.Location())
		play := &deviant.Play{
			X:    clone.X,
			Y:    clone.Y,
			Type: action.Type,
		}
		if action.Status != deviant.CardActionStatusTypes_BLOCKED {
			(*plays) = append((*plays), play)
			traversePlays(upLeft, plays, clone)
		}

		break
	}

	for _, upRight := range action.UpRight {
		clone := current.Clone()
		clone.Add(UPRIGHT.Location())
		play := &deviant.Play{
			X:    clone.X,
			Y:    clone.Y,
			Type: action.Type,
		}

		if action.Status != deviant.CardActionStatusTypes_BLOCKED {
			(*plays) = append((*plays), play)
			traversePlays(upRight, plays, clone)
		}

		break
	}

	for _, downLeft := range action.DownLeft {
		clone := current.Clone()
		clone.Add(DOWNLEFT.Location())
		play := &deviant.Play{
			X:    clone.X,
			Y:    clone.Y,
			Type: action.Type,
		}

		if action.Status != deviant.CardActionStatusTypes_BLOCKED {
			(*plays) = append((*plays), play)
			traversePlays(downLeft, plays, clone)
		}

		break
	}

	for _, downRight := range action.DownRight {
		clone := current.Clone()
		clone.Add(DOWNRIGHT.Location())
		play := &deviant.Play{
			X:    clone.X,
			Y:    clone.Y,
			Type: action.Type,
		}

		if action.Status != deviant.CardActionStatusTypes_BLOCKED {
			(*plays) = append((*plays), play)
			traversePlays(downRight, plays, clone)
		}

		break
	}
}

func processAction(location *Location, action *deviant.CardAction, parentAction *deviant.CardAction, board *deviant.Board) {
	for _, up := range action.Up {
		clone := location.Clone()
		clone.Add(UP.Location())

		// If the parent action is blocked or a hit then block the current action.
		if parentAction != nil {
			if parentAction.Status == deviant.CardActionStatusTypes_BLOCKED || parentAction.Status == deviant.CardActionStatusTypes_HIT {
				action.Status = deviant.CardActionStatusTypes_BLOCKED
			}
		}

		// If the current action hits something on the board.
		if int32(len(board.Entities.Entities)) <= clone.X && int32(len(board.Entities.Entities[clone.X].Entities)) < clone.Y {
			if board.Entities.Entities[clone.X].Entities[clone.Y].Id != "" {
				action.Status = deviant.CardActionStatusTypes_HIT
			}
		}

		processAction(clone, up, action, board)
	}

	for _, down := range action.Down {
		clone := location.Clone()
		clone.Add(DOWN.Location())

		// If the parent action is blocked or a hit then block the current action.
		if parentAction != nil {
			if parentAction.Status == deviant.CardActionStatusTypes_BLOCKED || parentAction.Status == deviant.CardActionStatusTypes_HIT {
				action.Status = deviant.CardActionStatusTypes_BLOCKED
			}
		}

		// If the current action hits something on the board.
		if int32(len(board.Entities.Entities)) <= clone.X && int32(len(board.Entities.Entities[clone.X].Entities)) < clone.Y {
			if board.Entities.Entities[clone.X].Entities[clone.Y].Id != "" {
				action.Status = deviant.CardActionStatusTypes_HIT
			}
		}

		processAction(clone, down, action, board)
	}

	for _, left := range action.Left {
		clone := location.Clone()
		clone.Add(LEFT.Location())
		// If the parent action is blocked or a hit then block the current action.
		if parentAction != nil {
			if parentAction.Status == deviant.CardActionStatusTypes_BLOCKED || parentAction.Status == deviant.CardActionStatusTypes_HIT {
				action.Status = deviant.CardActionStatusTypes_BLOCKED
			}
		}

		// If the current action hits something on the board.
		if int32(len(board.Entities.Entities)) <= clone.X && int32(len(board.Entities.Entities[clone.X].Entities)) < clone.Y {
			if board.Entities.Entities[clone.X].Entities[clone.Y].Id != "" {
				action.Status = deviant.CardActionStatusTypes_HIT
			}
		}

		processAction(clone, left, action, board)
	}

	for _, right := range action.Right {
		clone := location.Clone()
		clone.Add(RIGHT.Location())
		// If the parent action is blocked or a hit then block the current action.
		if parentAction != nil {
			if parentAction.Status == deviant.CardActionStatusTypes_BLOCKED || parentAction.Status == deviant.CardActionStatusTypes_HIT {
				action.Status = deviant.CardActionStatusTypes_BLOCKED
			}
		}

		// If the current action hits something on the board.
		if int32(len(board.Entities.Entities)) <= clone.X && int32(len(board.Entities.Entities[clone.X].Entities)) < clone.Y {
			if board.Entities.Entities[clone.X].Entities[clone.Y].Id != "" {
				action.Status = deviant.CardActionStatusTypes_HIT
			}
		}

		processAction(clone, right, action, board)
	}

	for _, upLeft := range action.UpLeft {
		clone := location.Clone()
		clone.Add(UPLEFT.Location())
		// If the parent action is blocked or a hit then block the current action.
		if parentAction != nil {
			if parentAction.Status == deviant.CardActionStatusTypes_BLOCKED || parentAction.Status == deviant.CardActionStatusTypes_HIT {
				action.Status = deviant.CardActionStatusTypes_BLOCKED
			}
		}

		// If the current action hits something on the board.
		if int32(len(board.Entities.Entities)) <= clone.X && int32(len(board.Entities.Entities[clone.X].Entities)) < clone.Y {
			if board.Entities.Entities[clone.X].Entities[clone.Y].Id != "" {
				action.Status = deviant.CardActionStatusTypes_HIT
			}
		}

		processAction(clone, upLeft, action, board)
	}

	for _, upRight := range action.UpRight {
		clone := location.Clone()
		clone.Add(UPRIGHT.Location())
		// If the parent action is blocked or a hit then block the current action.
		if parentAction != nil {
			if parentAction.Status == deviant.CardActionStatusTypes_BLOCKED || parentAction.Status == deviant.CardActionStatusTypes_HIT {
				action.Status = deviant.CardActionStatusTypes_BLOCKED
			}
		}

		// If the current action hits something on the board.
		if int32(len(board.Entities.Entities)) <= clone.X && int32(len(board.Entities.Entities[clone.X].Entities)) < clone.Y {
			if board.Entities.Entities[clone.X].Entities[clone.Y].Id != "" {
				action.Status = deviant.CardActionStatusTypes_HIT
			}
		}
		processAction(clone, upRight, action, board)
	}

	for _, downLeft := range action.DownLeft {
		clone := location.Clone()
		clone.Add(DOWNLEFT.Location())
		// If the parent action is blocked or a hit then block the current action.
		if parentAction != nil {
			if parentAction.Status == deviant.CardActionStatusTypes_BLOCKED || parentAction.Status == deviant.CardActionStatusTypes_HIT {
				action.Status = deviant.CardActionStatusTypes_BLOCKED
			}
		}

		// If the current action hits something on the board.
		if int32(len(board.Entities.Entities)) <= clone.X && int32(len(board.Entities.Entities[clone.X].Entities)) < clone.Y {
			if board.Entities.Entities[clone.X].Entities[clone.Y].Id != "" {
				action.Status = deviant.CardActionStatusTypes_HIT
			}
		}

		processAction(clone, downLeft, action, board)
	}

	for _, downRight := range action.DownRight {
		clone := location.Clone()
		clone.Add(DOWNRIGHT.Location())
		// If the parent action is blocked or a hit then block the current action.
		if parentAction != nil {
			if parentAction.Status == deviant.CardActionStatusTypes_BLOCKED || parentAction.Status == deviant.CardActionStatusTypes_HIT {
				action.Status = deviant.CardActionStatusTypes_BLOCKED
			}
		}

		// If the current action hits something on the board.
		if int32(len(board.Entities.Entities)) <= clone.X && int32(len(board.Entities.Entities[clone.X].Entities)) < clone.Y {
			if board.Entities.Entities[clone.X].Entities[clone.Y].Id != "" {
				action.Status = deviant.CardActionStatusTypes_HIT
			}
		}

		processAction(clone, downRight, action, board)
	}
}

func calculateBlocks(origin *Location, card *deviant.Card, board *deviant.Board) {
	for _, action := range card.Actions {
		processAction(origin, action, nil, board)
	}
}

// GetCardPlays Generates a list of plays where a card will be played.
func GetCardPlays(card *deviant.Card, start *Location, board *deviant.Board) *[]*deviant.Play {
	origin := Location{X: 0, Y: 0}
	plays := &[]*deviant.Play{}

	for _, action := range card.Actions {
		processAction(&origin, action, nil, board)
		traversePlays(action, plays, &origin)
	}

	return plays
}

// GetCardTargets Generates a list of tiles that a particular card would target
func GetCardTargets(card *deviant.Card, start *Location, board *deviant.Board) *[]*deviant.Tile {
	origin := Location{X: 0, Y: 0}
	targets := &[]*deviant.Tile{}

	for _, action := range card.Actions {
		processAction(&origin, action, nil, board)
		traverseTargets(action, targets, &origin)
	}

	return targets
}
