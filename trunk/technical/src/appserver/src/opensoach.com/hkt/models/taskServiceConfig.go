package models

type TaskSPDevAsscociatedModel struct {
	CpmId int64 `json:"cpmid"`
	DevId int64 `json:"devid"`
	SpId  int64 `json:"spid"`
}

type TaskSerConfigAddedOnSPModel struct {
	CpmId          int64 `json:"cpmid"`
	ServInstConfID int64 `json:"servconfinid"`
}

type TaskServConfigUpdatedModel struct {
	CpmId      int64 `json:"cpmid"`
	ServConfId int64 `json:"servconfid"`
}
