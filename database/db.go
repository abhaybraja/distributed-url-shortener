package database

import (
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

var RedisClient *redis.Client

func InitDB(host string, password string, db string) {

	DB, nerr := strconv.Atoi(db)
	if nerr != nil {
		fmt.Println("Error:", nerr)
	}

	// Redis connection
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // No password by default
		DB:       DB,       // Default DB
	})

	// Check Redis connection
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("❌ Redis connection failed:", err)
	} else {
		fmt.Println("✅ Connected to Redis")
	}

}
