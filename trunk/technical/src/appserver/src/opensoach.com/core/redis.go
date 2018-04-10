package core

import (
	"time"

	"github.com/go-redis/redis"
)

func (r *RedisContext) Connect(options *redis.Options) {
	r.RedisClient = redis.NewClient(options)
}

func (r RedisContext) Get(key string) (bool, string) {

	if r.RedisClient == nil {
		panic("Redis client is nil")
		return false, ""
	}

	value, err := r.RedisClient.Get(key).Result()

	if err != nil {
		return false, ""
	}

	return true, value

}

func (r RedisContext) Set(key string, value interface{}, t time.Duration) bool {
	if r.RedisClient == nil {
		panic("Redis client is nil")
		return false
	}

	err := r.RedisClient.Set(key, value, t).Err()

	if err != nil {
		return false
	}

	return true

}

func (r RedisContext) Update(key string, t time.Duration) bool {
	if r.RedisClient == nil {
		panic("Redis client is nil")
		return false
	}

	err := r.RedisClient.Expire(key, t).Err()

	if err != nil {
		return false
	}

	return true
}

func (r RedisContext) Remove(key string) bool {
	if r.RedisClient == nil {
		panic("Redis client is nil")
		return false
	}

	err := r.RedisClient.Expire(key, 0).Err()

	if err != nil {
		return false
	}

	return true
}
