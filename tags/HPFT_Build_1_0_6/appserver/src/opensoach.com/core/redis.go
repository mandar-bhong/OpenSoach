package core

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
)

var cacheAddressClient map[string]*redis.Client

func init() {
	cacheAddressClient = make(map[string]*redis.Client, 0)
}

func (r CacheContext) getClient() (bool, *redis.Client) {

	value, found := cacheAddressClient[r.CacheAddress]

	if found == true {
		return true, value
	}

	caheAddress := gmodels.ConfigCacheAddress{}
	ghelper.ConvertFromJSONString(r.CacheAddress, &caheAddress)

	client := redis.NewClient(&redis.Options{
		Addr:     caheAddress.Address + ":" + strconv.Itoa(caheAddress.Port),
		Password: caheAddress.Password,
		DB:       caheAddress.DB,
	})

	_, redisMstErr := client.Ping().Result()

	if redisMstErr != nil {
		logger.Context().
			WithField("Redis.Address", caheAddress.Address).
			WithField("Redis.Password", caheAddress.Password).
			LogError(SUB_MODULE_NAME, logger.Server, "Error occured while connecting to redis server. ", redisMstErr)

		return false, nil
	}

	cacheAddressClient[r.CacheAddress] = client

	return true, client

}

func (r CacheContext) Get(key string) (bool, string) {

	isClientGetSuccess, redisClient := r.getClient()

	if isClientGetSuccess == false {
		return false, ""
	}

	value, err := redisClient.Get(key).Result()

	if err != nil {
		logger.Context().WithField("Redis Key", key).LogError("Core.Redis", logger.Server, "Error occured while getting data from redis", err)
		return false, ""
	}

	return true, value

}

func (r CacheContext) Set(key string, value interface{}, t time.Duration) bool {

	isClientGetSuccess, redisClient := r.getClient()

	if isClientGetSuccess == false {
		return false
	}

	err := redisClient.Set(key, value, t).Err()

	if err != nil {
		logger.Context().WithField("Redis Key", key).LogError("Core.Redis", logger.Server, "Error occured while setting data to redis", err)
		return false
	}

	return true

}

func (r CacheContext) Update(key string, t time.Duration) bool {

	isClientGetSuccess, redisClient := r.getClient()

	if isClientGetSuccess == false {
		return false
	}

	err := redisClient.Expire(key, t).Err()

	if err != nil {
		logger.Context().WithField("Redis Key", key).LogError("Core.Redis", logger.Server, "Error occured while updating data to redis", err)
		return false
	}

	return true
}

func (r CacheContext) Remove(key string) bool {

	isClientGetSuccess, redisClient := r.getClient()

	if isClientGetSuccess == false {
		return false
	}

	err := redisClient.Expire(key, 0).Err()

	if err != nil {
		logger.Context().WithField("Redis Key", key).LogError("Core.Redis", logger.Server, "Error occured while removing key from redis", err)
		return false
	}

	return true
}
