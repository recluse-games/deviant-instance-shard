package matchmaker

import (
	"fmt"
	"testing"

	guuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_pool(t *testing.T) {
	var (
		id       = guuid.New().String()
		maxUsers = 5
	)

	// Mock user queue
	queuedUsers := []string{"ian", "cameron", "chris", "matt", "zach"}
	// Create a pool
	pool := newPool(id, maxUsers)
	// Add users to the pool
	for i := 0; i < maxUsers+1; i++ {
		if pool.ableToJoin() {
			pool.add(queuedUsers[i])
		}
	}

	// Get the users in the pool
	users := pool.users

	fmt.Printf("Pool: %v\n", pool)
	require.Equal(t, maxUsers, len(users))
}
