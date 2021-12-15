package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "instructor", "Gaurav", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "instructor").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("instructor:", val)

	// err = rdb.LPush(ctx, "participants", "Ajeet", "Eswar").Err()

	// if err != nil {
	// 	log.Println("Unable to push:", err)
	// 	return
	// }

	participants := make([]string, 10)

	// participants, err := rdb.LRange(ctx, "participants", 0, -1).Result()
	err = rdb.LRange(ctx, "participants", 0, -1).ScanSlice(&participants)

	if err != nil {
		log.Println("Unable to retrieve:", err)
		return
	}

	log.Println("Participants:", participants)
}
