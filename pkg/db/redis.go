package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"technoCredits/internal/app/models"
	"technoCredits/internal/security"
	"technoCredits/pkg/logger"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	ctx         = context.Background()
)

func InitializeRedis(redisParams models.RedisParams) error {
	var addr string

	if redisParams.Host != "" {
		addr = fmt.Sprintf("%s:%d", redisParams.Host, redisParams.Port)
	} else {
		addr = ":6379"
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisParams.Password,
		DB:       redisParams.DB,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
		return err
	}

	return nil
}

func SetCache(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		logger.Error.Printf("Error marshaling value for Redis: %v", err)
		return err
	}

	err = RedisClient.Set(
		ctx,
		key,
		data,
		time.Duration(security.AppSettings.RedisParams.TTLMinutes)*time.Minute,
	).Err()

	if err != nil {
		logger.Error.Printf("Error setting cache in Redis: %v", err)
		return err
	}
	return nil
}

func GetCache(key string, dest interface{}) (bool, error) {
	data, err := RedisClient.Get(ctx, key).Bytes()
	if err == redis.Nil {
		logger.Error.Printf("[db.GetCache] Key %s does not exist", key)
		return false, nil
	} else if err != nil {
		logger.Error.Printf("[db.GetCache] Error getting key %s from redis: %v", key, err)
		return false, err
	}

	err = json.Unmarshal(data, dest)
	if err != nil {
		logger.Error.Printf("[db.GetCache] Error unmarshaling data for key %s: %v", key, err)
		return false, err
	}

	return true, nil
}

func DeleteCache(key string) error {
	err := RedisClient.Del(ctx, key).Err()
	if err != nil {
		logger.Error.Printf("[db.DeleteCache] Error deleting cache from Redis: %v", err)
		return err
	}
	return nil
}

func CloseRedisConnection() error {
	err := RedisClient.Close()
	if err != nil {
		return err
	}

	return nil
}
