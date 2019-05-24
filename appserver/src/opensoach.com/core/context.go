package core

type Context struct {
	Master   DataStorage
	ProdMst  DataStorage
	ProdInst CacheContext
}

type DataStorage struct {
	Cache  CacheContext
	DBConn string
}

type CacheContext struct {
	CacheAddress string
}
