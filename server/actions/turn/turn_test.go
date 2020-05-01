package turn

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto"
)

func TestGrantAp(t *testing.T) {
	entityWithNoAp := &deviant.Entity{
		Ap: int32(0),
	}

	if GrantAp(entityWithNoAp) != true {
		t.Fail()
	}

	if entityWithNoAp.Ap != 5 {
		t.Fail()
	}
}
