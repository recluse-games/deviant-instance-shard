package rules

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestValidateMoveApCost(t *testing.T) {
	entity := &deviant.Entity{
		Ap: 5,
	}

	moveAction := &deviant.EntityMoveAction{
		StartXPosition: 0,
		StartYPosition: 0,
		FinalXPosition: 1,
		FinalYPosition: 3,
	}

	encounter := &deviant.Encounter{}

	isApCostValid := ValidateMoveApCost(entity, moveAction, encounter, nil)

	if isApCostValid != true {
		t.Fail()
	}
}

func TestValidateNewLocationEmpty(t *testing.T) {
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

	isMoveLocationValid := ValidateNewLocationEmpty(entity, moveAction, encounter, nil)

	if isMoveLocationValid == false {
		t.Fail()
	}
}

func TestValidateMovePermissable(t *testing.T) {
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

	result := ValidateMovePermissable(entity, moveAction, encounter, nil)

	if result == false {
		t.Fail()
	}
}
