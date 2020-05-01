package board

import deviant "github.com/recluse-games/deviant-protobuf/genproto"

// CreateBoard Creates a new board
func CreateBoard(entities *deviant.Entities, tiles *deviant.Tiles) *deviant.Board {
	board := &deviant.Board{
		Entities: entities,
		Tiles:    tiles,
	}

	return board
}
