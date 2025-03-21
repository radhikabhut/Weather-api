package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

// RedisClient struct holds the Redis client

// NewRedisClient initializes a new Redis client
func NewRedisClient(addr, password string, db int) error {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // Empty string means no password
		DB:       db,       // Default DB is 0
	})

	// Check connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}
	return err
}

// Set sets a key-value pair with an optional expiration time
func Set(key string, value interface{}, exp time.Duration) error {
	err := client.Set(context.Background(), key, value, exp).Err()
	if err != nil {
		return fmt.Errorf("failed to set key: %v", err)
	}
	return nil
}

// Get retrieves a value by key
func Get(key string) (string, error) {
	value, err := client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("failed to get key: %v", err)
	}
	return value, nil
}

// Delete removes a key from Redis
func Delete(key string) error {
	_, err := client.Del(context.Background(), key).Result()
	if err != nil {
		return fmt.Errorf("failed to delete key: %v", err)
	}
	return nil
}

// Exists checks if a key exists
func Exists(key string) (bool, error) {
	val, err := client.Exists(context.Background(), key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check key existence: %v", err)
	}
	return val > 0, nil
}

// Close closes the Redis connection
func Close() error {
	return client.Close()
}
