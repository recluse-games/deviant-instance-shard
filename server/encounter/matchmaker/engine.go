package matchmaker

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Engine struct hold required data for the matchmaker engine
type Engine struct {
	maxUsers     int
	waitPeriod   time.Duration
	mutex        sync.Mutex
	pools        []*pool
	expiredPools map[string]struct{}
}

// EngineOptions struct declare engine options
type EngineOptions struct {
	MaxUsers   int
	WaitPeriod time.Duration
}

// NewEngine function factory for Engines
func NewEngine(opt EngineOptions) *Engine {
	return &Engine{
		maxUsers:     opt.MaxUsers,
		waitPeriod:   time.Second * 2,
		expiredPools: make(map[string]struct{}),
	}
}

// GetNumberOfPools return number of pools
func (e *Engine) getNumberOfPools() int {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return len(e.pools)
}

// Return an available pool
func (e *Engine) getAvailablePool() *pool {
	numOfPools := e.getNumberOfPools()
	id := uuid.New().ID()
	sID := fmt.Sprintf("%d", id)

	// Create a pool if none exist
	if numOfPools == 0 {
		return e.createPool(sID)
	}

	// Loop through the pools on the engine and return the first available
	for _, v := range e.pools {
		if v.ableToJoin() {
			return v
		}
	}

	// Default to creating a new pool
	return e.createPool(sID)
}

// Create a pool for users to join
func (e *Engine) createPool(id string) *pool {
	pool := newPool(id, e.maxUsers)

	e.mutex.Lock()
	e.pools = append(e.pools, pool)
	e.mutex.Unlock()

	return pool
}

// JoinPool function add the given user into an available pool.
// Returns a channel of PoolResp that will notify when pool is full.
func (e *Engine) JoinPool(user interface{}) chan PoolResp {
	p := e.getAvailablePool()
	timer := time.NewTimer(e.waitPeriod)

	log.Printf("User: %v attempting to join pool: %v", user, p)
	go func() {
		select {
		case <-timer.C:
			e.mutex.Lock()

			// Make sure there is space in the pool
			if p.ableToJoin() {
				p.expireWaitCount++

				// Check if the pool is present in the expired pools map
				if _, present := e.expiredPools[p.id]; !present {
					p.respChan <- PoolResp{
						PoolID:   p.id,
						TimedOut: true,
						Users:    p.users,
					}
					e.expiredPools[p.id] = struct{}{}
				}

				// Destory the pool and reap the users
				if p.expireWaitCount == len(p.users) {
					// Remove users from pool
					p.users = nil
					// Remove pool from expired pools map
					delete(e.expiredPools, p.id)
				}
			}

			e.mutex.Unlock()
			break
		}
	}()

	return p.add(user)
}
