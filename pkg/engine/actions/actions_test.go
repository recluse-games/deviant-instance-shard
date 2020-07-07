package actions

import (
	"testing"
)

func TestRemoveEntityID(t *testing.T) {
	entityID1 := "0001"
	entityID2 := "0002"
	entityIDs := []string{entityID1, entityID2}

	updatedEntityIDs, _ := removeEntityID(entityID1, entityIDs)
	if updatedEntityIDs[0] != "0002" {
		t.Logf("EntityID removal failed for non-zero slice.")
		t.Fail()
	}

	updatedEntityIDs, _ = removeEntityID(entityID2, updatedEntityIDs)
	if updatedEntityIDs != nil {
		t.Logf("EntityID removal failed for 1 length slice.")
		t.Fail()
	}
}
