package repository

import (
	"sync"

	pcmodels "opensoach.com/prodcore/models"
)

var (
	r    *pcmodels.Repo
	once sync.Once
)


func Init() {
	once.Do(func() {
		r = &pcmodels.Repo{}
	})
}

func Instance() *pcmodels.Repo {
	return r
}
