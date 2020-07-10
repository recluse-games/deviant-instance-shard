package engineutil

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestGetEntity(t *testing.T) {
	mockEntity := &deviant.Entity{
		Id: "0001",
	}

	mockBoard := &deviant.Board{
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
			},
		},
	}

	entity, err := GetEntity("0001", mockBoard)
	if err != nil {
		t.Logf("Failed to locate entity on board by ID.")
		t.Fail()
	}

	if entity == nil {
		t.Logf("Failed to add a location to another location.")
		t.Fail()
	}

	entity, err = GetEntity("0002", mockBoard)
	if err == nil {
		t.Logf("Failed to properly raise an error if the entity isn't found.")
		t.Fail()
	}
}
