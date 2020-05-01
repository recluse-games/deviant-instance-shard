package turn

import deviant "github.com/recluse-games/deviant-protobuf/genproto"

// GrantAp grants some entity a default AP value of 5
func GrantAp(entity *deviant.Entity) bool {
	startingApValue := int32(5)
	entity.Ap = startingApValue

	return true
}
