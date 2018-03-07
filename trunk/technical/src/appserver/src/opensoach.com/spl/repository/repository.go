package repository

import (
	"sync"

	"opensoach.com/core"
)

var (
	r    *repo
	once sync.Once
)

type repo struct {
	MasterDBConnection string
	Context            *core.Context
}

func Init(mstdbConn string) {
	once.Do(func() {
		r = &repo{
			MasterDBConnection: mstdbConn,
			Context:            &core.Context{},
		}
	})
}

func Instance() *repo {
	return r
}
