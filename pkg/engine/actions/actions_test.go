package actions

import (
	"testing"

	"github.com/recluse-games/deviant-instance-shard/pkg/engine/engineutil"
)

func TestRemoveEntityID(t *testing.T) {
	entityID1 := "0001"
	entityID2 := "0002"
	entityIDs := []string{entityID1, entityID2}

	updatedEntityIDs, _ := engineutil.RemoveString(entityID1, entityIDs)
	if updatedEntityIDs[0] != "0002" {
		t.Logf("EntityID removal failed for non-zero slice.")
		t.Fail()
	}

	updatedEntityIDs, _ = engineutil.RemoveString(entityID2, updatedEntityIDs)
	if updatedEntityIDs != nil {
		t.Logf("EntityID removal failed for 1 length slice.")
		t.Fail()
	}
}
