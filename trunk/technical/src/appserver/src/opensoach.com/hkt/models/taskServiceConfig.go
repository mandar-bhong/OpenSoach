package models

type TaskSPDevAsscociatedModel struct {
	DeviceSPId int64 `json:"devspid"`
	CpmId      int64 `json:"cpmid"`
	DevId      int64 `json:"devid"`
	SpId       int64 `json:"spid"`
}

type TaskSerConfigAddedOnSPModel struct {
	CpmId      int64 `json:"cpmid"`
	DeviceSPId int64 `json:"devspid"`
	DevId      int64 `json:"devid"`
	SpId       int64 `json:"spid"`
}
