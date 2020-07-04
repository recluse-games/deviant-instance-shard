package tiles

import (
	"testing"
)

func TestProcess(t *testing.T) {
	tileIds := [][]string{{"grass_0000", "grass_0001"}, {"grass_0000", "grass_0001"}}
	createdTiles := CreateTiles(tileIds)

	// Validate That Tiles is being generated properly with expected data.
	if createdTiles.Tiles[0].Tiles[0].Id != "grass_0000" {
		t.Fail()
	}

	// Validate Out Tiles size is accurate.
	if len(createdTiles.Tiles) != 2 {
		t.Fail()
	}

	// Validate TilesRow size is accurate.
	if len(createdTiles.Tiles[0].Tiles) != 2 {
		t.Fail()
	}
}
