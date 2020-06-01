package move

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestValidateApCost(t *testing.T) {
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

	isApCostValid := ValidateApCost(entity, moveAction, encounter)

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

	isMoveLocationValid := ValidateNewLocationEmpty(entity, moveAction, encounter)

	if isMoveLocationValid == false {
		t.Fail()
	}
}
