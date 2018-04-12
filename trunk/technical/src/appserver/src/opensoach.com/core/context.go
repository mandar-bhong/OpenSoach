package core

type Context struct {
	Master   DataStorage
	ProdMst  DataStorage
	ProdInst DataStorage
}

type CacheContext struct {
	CacheAddress string
}

type DataStorage struct {
	Cache  CacheContext
	DBConn string
}
