package actions

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap/zaptest"
)

func TestProcessWinConditions(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entity := &deviant.Entity{
		Ap:        0,
		MaxAp:     5,
		Alignment: deviant.Alignment_FRIENDLY,
	}

	encounter := &deviant.Encounter{
		Board: &deviant.Board{
			Entities: &deviant.Entities{
				Entities: []*deviant.EntitiesRow{
					{
						Entities: []*deviant.Entity{
							entity,
						},
					},
				},
			},
		},
	}

	isEncounterComplete := ProcessWinConditions(encounter, logger.Sugar())

	if isEncounterComplete == false {
		t.Fail()
	}

	if encounter.WinningAlignment != deviant.Alignment_FRIENDLY {
		t.Logf("%v", encounter)
		t.Fail()
	}
}
