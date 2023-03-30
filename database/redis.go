package database

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/meta-byte/rocketeer-discord-bot/api"
	"github.com/meta-byte/rocketeer-discord-bot/types"
)

type Database interface {
	Cache()
	CacheInBackground(interval time.Duration)
	FetchLaunches() []types.Launch
}

type Redis struct {
	client *redis.Client
}

func NewRedis(redisAddress string) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return &Redis{client: client}
}

func (db *Redis) FlushDB() error {
	return db.client.FlushDB().Err()
}

func (db *Redis) Cache() {
	results, err := api.GetLaunchResults()
	if err != nil {
		fmt.Println("Error querying API:", err)
	}

	for _, v := range results.Results {
		serializedLaunch, err := json.Marshal(v)
		if err != nil {
			fmt.Println("Error serializing struct:", err)
			return
		}
		key := "launch:" + v.ID
		err = db.client.Set(key, serializedLaunch, 60*time.Minute).Err()
		fmt.Printf("cached: %v \n", key)
		if err != nil {
			fmt.Println("Error caching API response:", err)
		}
	}
}

func (db *Redis) CacheInBackground(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			db.Cache()
		}
	}
}

func (db *Redis) FetchLaunches() []types.Launch {
	launches := []types.Launch{}
	var launch types.Launch

	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = db.client.Scan(cursor, "launch:*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, v := range keys {
			serializedLaunch, err := db.client.Get(v).Bytes()
			if err != nil {
				if err == redis.Nil {
					fmt.Println("Key not found in Redis")
				} else {
					fmt.Println("Error getting data from Redis:", err)
				}
				continue
			}
			err = json.Unmarshal(serializedLaunch, &launch)
			log.Printf("Retrieved Launch: %v\n", launch.ID)
			launches = append(launches, launch)
			if err != nil {
				fmt.Println("Error deserializing data:", err)
				continue
			}
		}

		if cursor == 0 {
			break
		}
	}

	return launches
}
