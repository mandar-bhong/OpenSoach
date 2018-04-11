package core

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Context struct {
	Dynamic  ExecutionTime
	Master   DataStorage
	ProdMst  DataStorage
	ProdInst DataStorage
}

type ExecutionTime struct {
	Cache RedisContext
	DB    *sqlx.DB
}

type RedisContext struct {
	CacheAddress string
}

type DataStorage struct {
	Cache  RedisContext
	DBConn string
}
