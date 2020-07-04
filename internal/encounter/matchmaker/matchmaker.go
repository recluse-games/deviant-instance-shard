package matchmaker

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	enginetest "github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewCacheClient Get a new cache client
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
