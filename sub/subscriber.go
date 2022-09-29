package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	for {
		time.Sleep(time.Second)
		go func() {
			id := rand.Int()
			c := int(rand.Float32()*5) + 1
			rds := getRedis()
			defer rds.Close()
			ps := getPubSub(rds)
			defer ps.Close()
			key := fmt.Sprintf("%d", c)
			val := popFromList(rds, key)
			for val == "" { // empty list
				// wait for subscription
				fmt.Println(time.Now(), "wait for sub with id", id, " and key ", key)
				_, err := ps.ReceiveMessage()
				if err != nil {
					return
				}
				val = popFromList(rds, key)
			}
			fmt.Println("the end", " For key ", key, getLengthList(rds, key), " with id ", id)
		}()
	}
}

func popFromList(rds *redis.Client, key string) string {
	cmd := rds.LPop(key)
	val := cmd.Val()
	return val
}

func getLengthList(rds *redis.Client, key string) int64 {
	cmd := rds.LLen(key)
	i := cmd.Val()
	return i
}

func getRedis() *redis.Client {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:7979",
		Password: "",
		DB:       0,
	})
	_, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Println("error: ", err)
	}
	return redisdb
}

func getPubSub(client *redis.Client) *redis.PubSub {
	return client.Subscribe("1", "2", "3", "4", "5")
}
