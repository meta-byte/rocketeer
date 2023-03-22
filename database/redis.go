package database

import (
	"fmt"

	"github.com/go-redis/redis"
)

// type MyStruct struct {
// 	Data    string   `json:"data"`
// 	UserIDs []string `json:"user_ids"`
// }

func InitRedis(redisAddress string) {
	fmt.Println("Testing Golang Redis")

	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	// create a new struct
	// inputStruct := MyStruct{
	// 	Data:    "Test Data",
	// 	UserIDs: []string{"user1", "user2", "user3"},
	// }

	// encode the struct as JSON
	// jsonBytes, err := json.Marshal(inputStruct)
	// if err != nil {
	// 	panic(err)
	// }
	// jsonString := string(jsonBytes)

	// // set the struct in Redis
	// err = client.Set("myStruct1", jsonString, 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// err = client.Set("myStruct2", jsonString, 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// // read the struct and each user ID associated with it
	// jsonString, err = client.Get("myStruct2").Result()
	// if err != nil {
	// 	panic(err)
	// }

	// outputStruct := MyStruct{}
	// err = json.Unmarshal([]byte(jsonString), &outputStruct)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Retrieved struct with ID %s: %+v\n", "mystruct2", outputStruct)

	// jsonString, err = client.Get("myStruct1").Result()
	// if err != nil {
	// 	panic(err)
	// }

	// err = json.Unmarshal([]byte(jsonString), &outputStruct)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Retrieved struct with ID %s: %+v\n", "mystruct1", outputStruct)

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
