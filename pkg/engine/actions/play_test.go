package actions

import (
	"testing"

	"github.com/google/uuid"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap/zaptest"
)

func TestPlay(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

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
				Actions: []*deviant.CardAction{
					{
						Id:     uuid.New().String(),
						Status: deviant.CardActionStatusTypes_EMPTY,
						Down: []*deviant.CardAction{
							{
								Id:            uuid.New().String(),
								Type:          deviant.CardType_ATTACK,
								Status:        deviant.CardActionStatusTypes_UNBLOCKED,
								TargetingType: deviant.CardActionTargetingTypes_GROUND,
								Origin:        true,
								Down: []*deviant.CardAction{
									{
										Id:            uuid.New().String(),
										Type:          deviant.CardType_ATTACK,
										Status:        deviant.CardActionStatusTypes_UNBLOCKED,
										TargetingType: deviant.CardActionTargetingTypes_GROUND,
										Down: []*deviant.CardAction{
											{
												Id:            uuid.New().String(),
												Type:          deviant.CardType_ATTACK,
												Status:        deviant.CardActionStatusTypes_UNBLOCKED,
												TargetingType: deviant.CardActionTargetingTypes_GROUND,
											},
										},
									},
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
		Rotation: deviant.EntityRotationNames_NORTH,
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
				},
			},
		},
	}

	mockPlayAction := &deviant.EntityPlayAction{
		Id:     "0001",
		CardId: "12345",
	}

	if Play(mockEncounter, mockPlayAction, logger.Sugar()) != true {
		t.Fail()
	}

	if mockEncounter.Board.Entities.Entities[1].Entities[0].Hp != 0 {
		t.Logf("%v", mockEncounter.Board.Entities)
		t.Fail()
	}
}
