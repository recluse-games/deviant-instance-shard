package rules

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap/zaptest"
)

func TestValidateMoveApCost(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entity := &deviant.Entity{
		Id: "0001",
		Ap: 5,
	}

	moveAction := &deviant.EntityMoveAction{
		StartXPosition: 0,
		StartYPosition: 0,
		FinalXPosition: 1,
		FinalYPosition: 3,
	}

	encounter := &deviant.Encounter{
		ActiveEntity: entity,
	}

	isApCostValid := ValidateMoveApCost(entity, moveAction, encounter, logger.Sugar())

	if isApCostValid != true {
		t.Fail()
	}
}

func TestValidateNewLocationEmpty(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entity := &deviant.Entity{
		Ap: 5,
	}

	moveAction := &deviant.EntityMoveAction{
		StartXPosition: 0,
		StartYPosition: 0,
		FinalXPosition: 1,
		FinalYPosition: 3,
	}

	encounter := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{Id: "0001"},
		Board: &deviant.Board{
			Entities: &deviant.Entities{
				Entities: []*deviant.EntitiesRow{
					{
						Entities: []*deviant.Entity{
							entity,
							{},
							{},
							{},
						},
					},
					{
						Entities: []*deviant.Entity{
							{},
							{},
							{},
							{},
						},
					},
					{
						Entities: []*deviant.Entity{
							{},
							{},
							{},
							{},
						},
					},
				},
			},
		},
	}

	isMoveLocationValid := ValidateNewLocationEmpty(entity, moveAction, encounter, logger.Sugar())

	if isMoveLocationValid == false {
		t.Fail()
	}
}

func TestValidateMovePermissable(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entity := &deviant.Entity{
		Ap: 5,
	}

	moveAction := &deviant.EntityMoveAction{
		StartXPosition: 0,
		StartYPosition: 0,
		FinalXPosition: 1,
		FinalYPosition: 2,
	}

	encounter := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{Id: "0001"},
		Board: &deviant.Board{
			Entities: &deviant.Entities{
				Entities: []*deviant.EntitiesRow{
					{
						Entities: []*deviant.Entity{
							entity,
							{},
							{},
							{},
						},
					},
					{
						Entities: []*deviant.Entity{
							{},
							{},
							{},
							{},
						},
					},
					{
						Entities: []*deviant.Entity{
							{},
							{},
							{},
							{},
						},
					},
				},
			},
		},
	}

	result := ValidateMovePermissable(entity, moveAction, encounter, logger.Sugar())

	if result == false {
		t.Fail()
	}
}
