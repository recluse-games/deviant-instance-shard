package matchmaker

import (
	"fmt"
	"sync"
	"testing"
	"time"

	guuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_engine(t *testing.T) {
	// Total number of users
	numOfUsers := 10
	// Max number of users per pool
	maxUsers := 5
	// Required number of pools
	expectedNumOfPool := numOfUsers / maxUsers

	e := NewEngine(EngineOptions{MaxUsers: maxUsers})
	fmt.Printf("Engine: %v\n", e)
	for i := 0; i < numOfUsers; i++ {
		e.JoinPool(i)
	}

	numOfPools := e.getNumberOfPools()
	require.Equal(t, expectedNumOfPool, numOfPools)
}

func TestEngine_getNumberOfPools(t *testing.T) {
	var singlePool []*pool
	var severalPools []*pool

	singlePool = []*pool{newPool(guuid.New().String(), 5)}
	for i := 0; i < 10; i++ {
		severalPools = append(severalPools, newPool(guuid.New().String(), 5))
	}

	type fields struct {
		maxUsers     int
		waitPeriod   time.Duration
		mutex        sync.Mutex
		pools        []*pool
		expiredPools map[string]struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Single Pool",
			fields: fields{
				maxUsers:   100,
				waitPeriod: time.Second * 2,
				mutex:      sync.Mutex{},
				pools:      singlePool,
			},
			want: 1,
		},
		{
			name: "Several Pools",
			fields: fields{
				maxUsers:   100,
				waitPeriod: time.Second * 2,
				mutex:      sync.Mutex{},
				pools:      severalPools,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				maxUsers:     tt.fields.maxUsers,
				waitPeriod:   tt.fields.waitPeriod,
				mutex:        tt.fields.mutex,
				pools:        tt.fields.pools,
				expiredPools: tt.fields.expiredPools,
			}
			if got := e.getNumberOfPools(); got != tt.want {
				t.Errorf("Engine.GetNumberOfPools() = %v, want %v", got, tt.want)
			}
		})
	}
}
