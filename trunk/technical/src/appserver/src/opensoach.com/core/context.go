package core

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Context struct {
	Dynamic ExecutionTime
}

type ExecutionTime struct {
	Cache    RedisContext
	DB       *sqlx.DB
	ModCache RedisContext
}

type RedisContext struct {
	RedisClient *redis.Client
}
