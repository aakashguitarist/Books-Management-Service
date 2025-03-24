package cache

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

// InitRedis initializes the Redis client with error handling
func InitRedis() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"), // Leave empty if no password
		DB:       0,                           // Use default DB
	})

	// Test Redis connection
	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Printf("[ERROR] Failed to connect to Redis: %v", err)
		rdb = nil // Reset rdb to nil to prevent usage of an invalid client
		return err
	}

	log.Println("[INFO] Connected to Redis successfully")
	return nil
}

// Delete removes a key from Redis safely
func Delete(key string) error {
	if rdb == nil {
		return fmt.Errorf("Redis client is not initialized")
	}

	err := rdb.Del(context.Background(), key).Err()
	if err != nil {
		log.Printf("[ERROR] Failed to delete key %s from Redis: %v", key, err)
		return err
	}

	log.Printf("[INFO] Key %s deleted successfully from Redis", key)
	return nil
}

// Get retrieves a value from Redis safely
func Get(key string) (string, error) {
	if rdb == nil {
		return "", fmt.Errorf("Redis client is not initialized")
	}

	val, err := rdb.Get(context.Background(), key).Result()
	if err == redis.Nil {
		log.Printf("[WARN] Key %s does not exist in Redis", key)
		return "", nil
	} else if err != nil {
		log.Printf("[ERROR] Failed to get key %s from Redis: %v", key, err)
		return "", err
	}

	log.Printf("[INFO] Key %s retrieved successfully from Redis", key)
	return val, nil
}

// Set stores a key-value pair in Redis safely
func Set(key string, value string) error {
	if rdb == nil {
		return fmt.Errorf("Redis client is not initialized")
	}

	err := rdb.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		log.Printf("[ERROR] Failed to set key %s in Redis: %v", key, err)
		return err
	}

	log.Printf("[INFO] Key %s set successfully in Redis", key)
	return nil
}
