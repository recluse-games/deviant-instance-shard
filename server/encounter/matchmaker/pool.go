package matchmaker

import (
	"sync"
)

type (
	// Pool of users that will be commited to the same in game match
	Pool struct {
		expireWaitCount int
		id              string
		users           []interface{}
		mutex           sync.Mutex
		maxUsers        int
		respChan        chan PoolResp
	}

	// PoolResp function response for when a user is joining the pool
	PoolResp struct {
		IsFull   bool
		Users    []interface{}
		PoolID   string
		TimedOut bool
	}
)

// NewPool func create new pool
func newPool(id string, maxUsers int) *Pool {
	return &Pool{
		id:       id,
		maxUsers: maxUsers,
		respChan: make(chan PoolResp, maxUsers),
	}
}

// Check is a user is able to join a pool
func (p *Pool) ableToJoin() bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return len(p.users) < p.maxUsers
}

// Add a user to the pool
func (p *Pool) add(user interface{}) chan PoolResp {
	p.mutex.Lock()
	defer func() {
		if len(p.users) == p.maxUsers {
			p.respChan <- PoolResp{
				PoolID: p.id,
				IsFull: true,
				Users:  p.users,
			}
		} else {
			p.respChan <- PoolResp{
				PoolID: p.id,
				IsFull: false,
				Users:  p.users,
			}
		}
		p.mutex.Unlock()
	}()

	p.users = append(p.users, user)

	return p.respChan
}
