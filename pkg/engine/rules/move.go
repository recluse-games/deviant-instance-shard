package rules

import (
	"github.com/recluse-games/deviant-instance-shard/pkg/engine/engineutil"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap"
)

// GeneratePermissableMoves Generates a slice of tiles that represent moves that are perm
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

	engineutil.FloodFill(requestedMoveAction.StartXPosition, requestedMoveAction.StartYPosition, requestedMoveAction.StartXPosition, requestedMoveAction.StartYPosition, "select_0000", "select_0002", avaliableAp, moveTargetTiles)

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
func ValidateMovePermissable(activeEntity *deviant.Entity, requestedMoveAction *deviant.EntityMoveAction, encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
	isRequestedMoveValid := false

	validTiles := GeneratePermissableMoves(requestedMoveAction, activeEntity.Ap, encounter.Board.Entities)

	for _, tile := range validTiles {
		if tile.X == requestedMoveAction.FinalXPosition && tile.Y == requestedMoveAction.FinalYPosition {
			isRequestedMoveValid = true
		}
	}

	logger.Debug("Validated Move Permissable",
		"actionID", "ValidateMovePermissable",
		"entityID", encounter.ActiveEntity.Id,
		"succeeded", isRequestedMoveValid,
	)

	return isRequestedMoveValid
}

// ValidateMoveApCost Determines that the entity has the correct amount of AP to perform the requested move.
func ValidateMoveApCost(activeEntity *deviant.Entity, requestedMoveAction *deviant.EntityMoveAction, encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
	status := false

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
		status = true
	}

	logger.Debugw("Validated New Move AP Cost",
		"actionID", "ValidateMoveApCost",
		"entityID", encounter.ActiveEntity.Id,
		"succeeded", status,
	)

	return status
}

// ValidateNewLocationEmpty Determines that the entity can actually move to this location on the board.
func ValidateNewLocationEmpty(activeEntity *deviant.Entity, requestedMoveAction *deviant.EntityMoveAction, encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
	status := false

	// Validate that the new location is empty
	if encounter.Board.Entities.Entities[requestedMoveAction.FinalXPosition].Entities[requestedMoveAction.FinalYPosition].Id == "" {
		status = true
	}

	logger.Debug("Validated New Location Empty",
		"actionID", "ValidateNewLocationEmpty",
		"entityID", encounter.ActiveEntity.Id,
		"succeeded", status,
	)

	return status
}
