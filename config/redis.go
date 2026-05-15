package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var RedisClient *redis.Client

func ConnectRedis() error {

	// This code works fine

	
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // default DB
	})

	// Test Connection
	res, err := RedisClient.Ping(Ctx).Result()

	if err != nil {
		panic(err)
	}
	_ = RedisClient.Set(context.TODO(), "ay7ush", "akshit",time.Minute * 10).Err()
	fmt.Println("Result:",res)
	fmt.Println("Redis Connected Successfully")
	

	// Code to panic the system, fetal

	/*
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Check Redis Connection
	_, err := RedisClient.Ping(Ctx).Result()

	if err != nil {
		return err
	}

	// Test Redis Write
	err = RedisClient.Set(
		context.TODO(),
		"ay7ush",
		"akshit",
		time.Minute*10,
		
	).Err()

	if err != nil {
		return err
	}

	fmt.Println("Redis Connected Successfully")

	*/

	return nil	

}
