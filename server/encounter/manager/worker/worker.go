package worker

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/golang/glog"
	actions "github.com/recluse-games/deviant-instance-shard/server/actions/processor"
	"github.com/recluse-games/deviant-instance-shard/server/encounter/matchmaker"
	rules "github.com/recluse-games/deviant-instance-shard/server/rules/processor"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewCacheClient accss the cache
func NewCacheClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}

// IncomingWorker A worker to process work.
type IncomingWorker struct {
	ID int
}

// NewIncomingWorker creates, and returns a new Worker object.
func NewIncomingWorker(id int) IncomingWorker {
	// Create, and return the worker.
	worker := IncomingWorker{
		ID: id,
	}

	return worker
}

// StartIncoming Starts a working with an infinite loop.
func (w *IncomingWorker) ProcessWork(work *deviant.EncounterRequest) *deviant.EncounterResponse {
	var redisClient = NewCacheClient()

	// Setup Options for Marshalling and Unmarshalling JSON
	var marshalOptions = protojson.MarshalOptions{
		AllowPartial:    true,
		Multiline:       true,
		EmitUnpopulated: true,
	}

	var unmarshalOptions = protojson.UnmarshalOptions{
		AllowPartial: true,
	}

	// HACK: We should really just process our action message type here and switch on this rather then this crazy conditional logic.
	// Add ourselves into the worker queue.
	var actionResponse = &deviant.EncounterResponse{}
	var encounterFromDisk = &deviant.EncounterResponse{}

	in, err := redisClient.Get("encounter_0000").Result()

	if err == nil {
		unmarshalError := protojson.UnmarshalOptions(unmarshalOptions).Unmarshal([]byte(in), encounterFromDisk)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
	}

	if work.EncounterCreateAction != nil {
		actionResponse = matchmaker.GenerateMatch()
		actions.Process(actionResponse.Encounter, deviant.EntityActionNames_NOTHING, nil, nil)

		encounterFromDisk = actionResponse
	} else if work.EntityGetAction != nil {
		actionResponse = &deviant.EncounterResponse{
			PlayerId:  work.PlayerId,
			Encounter: encounterFromDisk.Encounter,
		}
	} else if work.EntityTargetAction != nil {
		// Update overlay tiles to the new tiles
		encounterFromDisk.Encounter.Board.OverlayTiles = work.EntityTargetAction.Tiles

		actionResponse = &deviant.EncounterResponse{
			PlayerId:  work.PlayerId,
			Encounter: encounterFromDisk.Encounter,
		}
	} else if work.EntityStateAction != nil {
		// Apply all state changes to entity in encounter as well as the activeEntity
		for outerIndex, outerValue := range encounterFromDisk.Encounter.Board.Entities.Entities {
			for innerIndex, innerValue := range outerValue.Entities {
				if innerValue.Id == work.EntityStateAction.Id {
					encounterFromDisk.Encounter.Board.Entities.Entities[outerIndex].Entities[innerIndex].State = work.EntityStateAction.State
				}
			}
		}

		if work.EntityStateAction.Id == encounterFromDisk.Encounter.ActiveEntity.Id {
			encounterFromDisk.Encounter.ActiveEntity.State = work.EntityStateAction.State
		}

		actionResponse = &deviant.EncounterResponse{
			PlayerId:  work.PlayerId,
			Encounter: encounterFromDisk.Encounter,
		}
	} else {
		// AuthZ the Player <- This should be migrated to a different layer of the codebase
		if work.PlayerId == encounterFromDisk.Encounter.ActiveEntity.OwnerId {
			isActionValid := rules.Process(encounterFromDisk.Encounter, work.EntityActionName, work.EntityMoveAction, work.EntityPlayAction)
			if isActionValid == true {
				actions.Process(encounterFromDisk.Encounter, work.EntityActionName, work.EntityMoveAction, work.EntityPlayAction)

				// Apply all state changes to entity in encounter as well as the activeEntity
				for outerIndex, outerValue := range encounterFromDisk.Encounter.Board.Entities.Entities {
					for innerIndex, innerValue := range outerValue.Entities {
						if innerValue.Id == encounterFromDisk.Encounter.ActiveEntity.Id {
							encounterFromDisk.Encounter.Board.Entities.Entities[outerIndex].Entities[innerIndex] = encounterFromDisk.Encounter.ActiveEntity
						}
					}
				}
			}
		}

		actionResponse = &deviant.EncounterResponse{
			PlayerId:  work.PlayerId,
			Encounter: encounterFromDisk.Encounter,
		}
	}

	result, marshalError := protojson.MarshalOptions(marshalOptions).Marshal(encounterFromDisk)
	if marshalError != nil {
		panic(marshalError)
	}

	writeErr := redisClient.Set("encounter_0000", string(result), 0).Err()
	if writeErr != nil {
		panic(writeErr)
	}

	message := fmt.Sprintf("Actions Processed: %s\n", work.EntityActionName)
	glog.Info(message)

	return actionResponse
}
