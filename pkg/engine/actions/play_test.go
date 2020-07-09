package actions

import (
	"testing"

	"github.com/google/uuid"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestPlay(t *testing.T) {
	mockHand := &deviant.Hand{
		Cards: []*deviant.Card{
			{
				Id:          "attack_slash_0000",
				BackId:      "back_0000",
				InstanceId:  "12345",
				Cost:        2,
				Damage:      2,
				Title:       "Slash",
				Flavor:      "Downward Dog",
				Description: "A Simple Slash",
				Type:        deviant.CardType_ATTACK,
				Action: &deviant.CardAction{
					Id: uuid.New().String(),
					Pattern: []*deviant.Pattern{
						{
							Direction: deviant.Direction_DOWN,
							Distance:  3,
							Offset: []*deviant.Offset{
								{
									Direction: deviant.Direction_DOWN,
									Distance:  1,
								},
							},
						},
					},
				},
			},
		},
	}

	mockEntity := &deviant.Entity{
		Id:       "0001",
		Hp:       1,
		Ap:       2,
		Hand:     mockHand,
		Rotation: deviant.EntityRotationNames_EAST,
		Discard:  &deviant.Discard{},
	}

	mockedEnemy := &deviant.Entity{
		Id: "0002s",
		Hp: 2,
	}

	mockEncounter := &deviant.Encounter{
		ActiveEntity: mockEntity,
		Board: &deviant.Board{
			Entities: &deviant.Entities{
				Entities: []*deviant.EntitiesRow{
					{
						Entities: []*deviant.Entity{
							mockEntity,
							{},
							{},
							{},
							{},
							{},
							{},
							{},
							{},
						},
					},
					{
						Entities: []*deviant.Entity{
							mockedEnemy,
							{},
							{},
							{},
							{},
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
							{},
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
							{},
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

	mockPlayAction := &deviant.EntityPlayAction{
		Id:     "0001",
		CardId: "12345",
	}

	if Play(mockEncounter, mockPlayAction, nil) != true {
		t.Fail()
	}

	if mockEncounter.Board.Entities.Entities[1].Entities[0].Hp != 0 {
		t.Logf("%v", mockEncounter.Board.Entities)
		t.Fail()
	}
}
