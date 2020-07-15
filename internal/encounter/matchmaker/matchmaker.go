package matchmaker

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
	enginetest "github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewCacheClient Get a new cache client
func NewCacheClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Error attempting to ping Redis cache")
	}
	return client
}

// GenerateMatch Generates a new match
func GenerateMatch() *deviant.EncounterResponse {
	test := enginetest.GenerateMatchObject()

	var marshalOptions = protojson.MarshalOptions{
		AllowPartial:    true,
		EmitUnpopulated: true,
	}

	result, _ := protojson.MarshalOptions(marshalOptions).Marshal(test)
	redisClient := NewCacheClient()
	err := redisClient.Set("encounter_0000", string(result), 0).Err()
	if err != nil {
		panic(err)
	}

	return test
}
