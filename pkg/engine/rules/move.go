package rules

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

func boundaryFill4(startx int32, starty int32, x int32, y int32, filledID string, blockedID string, limit int32, tiles []*[]*deviant.Tile) {
	if (*tiles[x])[y].Id != blockedID && (*tiles[x])[y].Id != filledID {
		var apCostX int32
		var apCostY int32

		if startx > x {
			apCostX = startx - x
		} else if startx < x {
			apCostX = x - startx
		} else {
			apCostX = 0
		}

		if starty > y {
			apCostY = starty - y
		} else if starty < y {
			apCostY = y - starty
		} else {
			apCostY = 0
		}

		newTile := &deviant.Tile{}
		newTile.X = int32(x)
		newTile.Y = int32(y)
		newTile.Id = filledID
		(*tiles[x])[y] = newTile

		if limit-apCostX-apCostY >= 0 {
			if x+1 <= 8 {
				boundaryFill4(startx, starty, x+1, y, filledID, blockedID, limit, tiles)
			}

			if y+1 <= 7 {
				boundaryFill4(startx, starty, x, y+1, filledID, blockedID, limit, tiles)
			}

			if x-1 >= 0 {
				boundaryFill4(startx, starty, x-1, y, filledID, blockedID, limit, tiles)
			}

			if y-1 >= 0 {
				boundaryFill4(startx, starty, x, y-1, filledID, blockedID, limit, tiles)
			}
		}
	}
}

func GeneratePermissableMoves(requestedMoveAction *deviant.EntityMoveAction, avaliableAp int32, entities *deviant.Entities) []*deviant.Tile {
	finalTiles := []*deviant.Tile{}
	moveTargetTiles := []*[]*deviant.Tile{}

	for y := 0; y < len(entities.Entities); y++ {
		newRow := []*deviant.Tile{}

		for x := 0; x < len(entities.Entities[y].Entities); x++ {
			newTile := deviant.Tile{}
			newTile.X = int32(y)
			newTile.Y = int32(x)

			if entities.Entities[y].Entities[x].Id != "" {
				newTile.Id = "select_0002"
			}

			if int32(y) == requestedMoveAction.StartXPosition && int32(x) == requestedMoveAction.StartYPosition {
				newTile.Id = "select_0001"
			}

			newRow = append(newRow, &newTile)
		}

		moveTargetTiles = append(moveTargetTiles, &newRow)
	}

	boundaryFill4(requestedMoveAction.StartXPosition, requestedMoveAction.StartYPosition, requestedMoveAction.StartXPosition, requestedMoveAction.StartYPosition, "select_0000", "select_0002", avaliableAp, moveTargetTiles)

	for _, row := range moveTargetTiles {
		for _, tile := range *row {
			if (*tile).Id == "select_0000" {
				finalTiles = append(finalTiles, tile)
			}
		}
	}

	return finalTiles
}

// ValidateMovePermissable Determines if the move is permissable using a flood fill algorithm and ap cost.
func ValidateMovePermissable(activeEntity *deviant.Entity, requestedMoveAction *deviant.EntityMoveAction, encounter *deviant.Encounter, logger *zap.Logger) bool {
	isRequestedMoveValid := false

	validTiles := GeneratePermissableMoves(requestedMoveAction, activeEntity.Ap, encounter.Board.Entities)

	for _, tile := range validTiles {
		if tile.X == requestedMoveAction.FinalXPosition && tile.Y == requestedMoveAction.FinalYPosition {
			if logger != nil {
				logger.Debug("Validated Move Permissable",
					zap.String("actionID", "ValidateMovePermissable"),
					zap.String("entityID", encounter.ActiveEntity.Id),
					zap.Bool("succeeded", true),
				)
			}

			isRequestedMoveValid = true
		}
	}

	if isRequestedMoveValid == false {
		if logger != nil {
			logger.Debug("Validated Move Permissable",
				zap.String("actionID", "ValidateMovePermissable"),
				zap.String("entityID", encounter.ActiveEntity.Id),
				zap.Bool("succeeded", false),
			)
		}
	}

	return isRequestedMoveValid
}

// ValidateMoveApCost Determines that the entity has the correct amount of AP to perform the requested move.
func ValidateMoveApCost(activeEntity *deviant.Entity, requestedMoveAction *deviant.EntityMoveAction, encounter *deviant.Encounter, logger *zap.Logger) bool {
	var totalApCost int32
	var apCostX int32
	var apCostY int32

	if requestedMoveAction.StartXPosition > requestedMoveAction.FinalXPosition {
		apCostX = requestedMoveAction.StartXPosition - requestedMoveAction.FinalXPosition
	} else if requestedMoveAction.StartXPosition < requestedMoveAction.FinalXPosition {
		apCostX = requestedMoveAction.FinalXPosition - requestedMoveAction.StartXPosition
	} else {
		apCostX = 0
	}

	if requestedMoveAction.StartYPosition > requestedMoveAction.FinalYPosition {
		apCostY = requestedMoveAction.StartYPosition - requestedMoveAction.FinalYPosition
	} else if requestedMoveAction.StartYPosition < requestedMoveAction.FinalYPosition {
		apCostY = requestedMoveAction.FinalYPosition - requestedMoveAction.StartYPosition
	} else {
		apCostY = 0
	}

	totalApCost = apCostX + apCostY

	if totalApCost <= activeEntity.Ap {
		if logger != nil {
			logger.Debug("Validated New Move AP Cost",
				zap.String("actionID", "ValidateMoveApCost"),
				zap.String("entityID", encounter.ActiveEntity.Id),
				zap.Bool("succeeded", true),
			)
		}

		return true
	}

	if logger != nil {
		logger.Debug("Validated New Move AP Cost",
			zap.String("actionID", "ValidateMoveApCost"),
			zap.String("entityID", encounter.ActiveEntity.Id),
			zap.Bool("succeeded", false),
		)
	}

	return false
}

// ValidateNewLocationEmpty Determines that the entity can actually move to this location on the board.
func ValidateNewLocationEmpty(activeEntity *deviant.Entity, requestedMoveAction *deviant.EntityMoveAction, encounter *deviant.Encounter, logger *zap.Logger) bool {
	// Validate that the new location is empty
	if encounter.Board.Entities.Entities[requestedMoveAction.FinalXPosition].Entities[requestedMoveAction.FinalYPosition].Id == "" {
		if logger != nil {
			logger.Debug("Validated New Location Empty",
				zap.String("actionID", "ValidateNewLocationEmpty"),
				zap.String("entityID", encounter.ActiveEntity.Id),
				zap.Bool("succeeded", true),
			)
		}

		return true
	}

	if logger != nil {
		logger.Debug("Validated New Location Empty",
			zap.String("actionID", "ValidateNewLocationEmpty"),
			zap.String("entityID", encounter.ActiveEntity.Id),
			zap.Bool("succeeded", false),
		)
	}

	return false
}
