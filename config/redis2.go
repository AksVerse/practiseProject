package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var Cntx = context.Background()

var RedisCli *redis.Client

func connectRedis() {
	RedisCli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	res, err := RedisCli.Ping(Cntx).Result()

	if err != nil {
		fmt.Println("Error occured")
		return 
	}

	_ = RedisCli.Set(context.TODO(),"lastName","Verma",5*time.Second).Err()

	fmt.Println("Result: ",res)
	fmt.Println("Connected Successfully")
}