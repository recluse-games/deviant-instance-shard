package engineutil

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
)

func TestAdd(t *testing.T) {
	location := Location{X: 0, Y: 0}
	locationToAdd := Location{X: 1, Y: 0}

	location.Add(locationToAdd)

	if location.X != 1 && location.Y != 0 {
		t.Logf("Failed to add a location to another location.")
		t.Fail()
	}
}

func TestClone(t *testing.T) {
	location := Location{X: 0, Y: 0}
	locationClone := location.Clone()

	location.Add(Location{X: 1, Y: 0})

	if locationClone.X != 0 && locationClone.Y != 0 {
		t.Logf("Clone failed to create a unique copy.")
		t.Fail()
	}
}

func TestMove(t *testing.T) {
	location := Location{X: 0, Y: 0}

	location.Move(UP, 1)
	if location.X != 1 && location.Y != 0 {
		t.Logf("Failed to move UP.")
		t.Fail()
	}

	location.Move(DOWN, 1)
	if location.X != 0 && location.Y != 0 {
		t.Logf("Failed to move DOWN")
		t.Fail()
	}

	location.Move(LEFT, 1)
	if location.X != 0 && location.Y != -1 {
		t.Logf("Failed to move LEFT")
		t.Fail()
	}

	location.Move(RIGHT, 1)
	if location.X != 0 && location.Y != 0 {
		t.Logf("Failed to move RIGHT")
		t.Fail()
	}
}

func TestRotate(t *testing.T) {
	origin := Location{X: 0, Y: 0}
	location := Location{X: 3, Y: 3}

	location.Rotate(origin, 90)

	if location.X != -3 && location.Y != -3 {
		t.Logf("Failed to rotate Location 90*")
		t.Fail()
	}
}

func TestConvertLocation(t *testing.T) {
	north := 180.00
	if GetDegree(deviant.EntityRotationNames_NORTH) != north {
		t.Logf("Failed to get the proper degree for north")
		t.Fail()
	}

	south := 0.00
	if GetDegree(deviant.EntityRotationNames_SOUTH) != south {
		t.Logf("Failed to get the proper degree for south")
		t.Fail()
	}

	east := 270.00
	if GetDegree(deviant.EntityRotationNames_EAST) != east {
		t.Logf("Failed to get the proper degree for east")
		t.Fail()
	}

	west := 90.00
	if GetDegree(deviant.EntityRotationNames_WEST) != west {
		t.Logf("Failed to get the proper degree for west")
		t.Fail()
	}
}

func TestRotateLocations(t *testing.T) {
	origin := Location{X: 0, Y: 0}
	locations := []*Location{}
	rotation := deviant.EntityRotationNames_NORTH

	for i := 1; i < 4; i++ {
		location := Location{X: int32(i), Y: 0}
		locations = append(locations, &location)
	}

	RotateLocations(origin, locations, rotation)

	for i := 0; i < len(locations); i++ {
		if locations[i].X != int32(-(i + 1)) {
			t.Logf("Failed to rotate location in locations")
			t.Fail()
		}
	}
}
