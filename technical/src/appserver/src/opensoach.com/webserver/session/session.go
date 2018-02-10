package session

import (
	"time"

	goCache "github.com/patrickmn/go-cache"
)

var MODULENAME = "Session"

var cache *goCache.Cache

func init() {
	cache = goCache.New(goCache.NoExpiration, 1)
}

func Init() bool {

	return true
}

func Get(key string) (bool, string) {
	data, isSuccess := cache.Get(key)

	if !isSuccess {
		return isSuccess, ""
	}

	return isSuccess, data.(string)
}

func Set(key string, value string, duration time.Duration) {
	cache.Set(key, value, duration)
}

func Replace(key string, value string, duration time.Duration) bool {
	err := cache.Replace(key, value, duration)

	if err != nil {
		//logger.Log(MODULENAME, logger.ERROR, "Unable to replace session for for key: ", key)
		return false
	}

	return true
}

func DeleteKey(key string) {
	cache.Delete(key)
}
