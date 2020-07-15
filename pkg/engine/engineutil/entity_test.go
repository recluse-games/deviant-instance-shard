package engineutil

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
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

func TestRemoveString(t *testing.T) {
	entityID := "0001"
	entityList := []string{"0001", "0002", "0003"}

	newSlice, err := RemoveString(entityID, entityList)
	if err != nil {
		t.Logf("Failed to remove entity ID from slice of entity IDs")
		t.Fail()
	}

	if newSlice[0] != "0002" {
		t.Logf("Failed to remove entity ID from slice of entity IDs")
		t.Fail()
	}

	newSlice, err = RemoveString(entityID, entityList)
	if err == nil {
		t.Logf("Failed to raise error when ID is missing from slice of IDs")
		t.Fail()
	}
}

func TestLocateEntity(t *testing.T) {
	entityID := "0001"

	mockEntity := &deviant.Entity{
		Id: "0001",
	}

	mockBoard := &deviant.Board{
		Entities: &deviant.Entities{
			Entities: []*deviant.EntitiesRow{
				{
					Entities: []*deviant.Entity{
						{},
						mockEntity,
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

	location, err := LocateEntity(entityID, mockBoard)
	if err != nil {
		t.Logf("Failed to generate entity location from board")
		t.Fail()
	}

	emptyBoard := &deviant.Board{}
	location, err = LocateEntity(entityID, emptyBoard)
	if err == nil {
		t.Logf("Failed to generate error when entity is not on board")
		t.Fail()
	}

	if location != nil {
		t.Logf("Failed to return a nil value when entity isn't on board")
		t.Fail()
	}
}

func TestIndexString(t *testing.T) {
	stringList := []string{"0001", "0002", "0003"}

	i, err := IndexString(stringList, "0001")
	if err != nil {
		t.Logf("Failed to locate string index in slice of strings")
		t.Fail()
	}

	if *i != 0 {
		t.Logf("Located incorrect index for string in slice")
		t.Fail()
	}

	stringList = []string{"0002", "0003"}
	i, err = IndexString(stringList, "0001")
	if err == nil {
		t.Logf("Failed to throw error when string is missing from slice of strings")
		t.Fail()
	}

	if i != nil {
		t.Logf("Failed to return nil when string not found in slice of strings")
		t.Fail()
	}
}
