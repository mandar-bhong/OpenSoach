package services

import (
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

type IHandler interface {
	Handle(*ServiceContext) error
}

type ServiceContext struct {
	Repo           pcmodels.Repo
	ServiceRuntime ServiceRuntimeModel
	ServiceConfig  ServiceConfigModel
	ServiceResult  ServiceResultModel
}

type ServiceConfigModel struct {
	SourceToken     string
	SourcePacket    *gmodels.DevicePacket
	DestinationData interface{}
	InstDBConn      string //instance db connection - GetOperatorLocService
	ID              int64  // for operator id - GetOperatorLocService
	AckData         interface{}
	CPMID           int64
	// InputData       interface{}
}

type ServiceResultModel struct {
	DestinationPackets []pcmodels.TaskEPPacketSendDataModel
	AckPacket          []pcmodels.TaskEPPacketSendDataModel
}

type ServiceRuntimeModel struct {
	Tokens               []string
	CPMID                []int64
	DeviceID             []int64
	Location             []int64
	UserID               int64
	DeviceTokenModelList []gmodels.DeviceTokenModel
	//DestinationPackets []*pcmodels.TaskEPPacketSendDataModel
	ExecutionData     interface{}
	DeviceLocationMap map[int64][]int64
}

type OperatorData struct {
	CpmId int64 `json:"cpmid"`
	FopId int64 `json:"fopid"`
	SpId  int64 `json:"spid"`
}
