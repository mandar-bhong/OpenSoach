package api

import (
	lmodels "opensoach.com/hpft/models"
	repo "opensoach.com/hpft/server/repository"
	gmodels "opensoach.com/models"
	pcconstants "opensoach.com/prodcore/constants"
	pcmodels "opensoach.com/prodcore/models"
	pcservices "opensoach.com/prodcore/services"
)

func NotifyDBChanges(ctx *pcmodels.APITaskExecutionCtx) (error, *pcmodels.APITaskProcessorResultModel) {

	apiTaskProcessorResultModel := &pcmodels.APITaskProcessorResultModel{}

	taskDBChangesModel := ctx.TaskData.(*lmodels.TaskDBChangesModel)

	storeSyncModel := pcmodels.StoreSyncModel{}
	storeSyncModel.StoreName = taskDBChangesModel.StoreName

	devPacket := &gmodels.DevicePacket{}
	devPacket.Header.Category = pcconstants.DEVICE_CMD_CAT_DATA
	devPacket.Header.CommandID = pcconstants.DEVICE_CMD_STORE_APPLY_SYNC

	serviceCtx := &pcservices.ServiceContext{}
	serviceCtx.Repo = *repo.Instance()
	serviceCtx.ServiceConfig.CPMID = taskDBChangesModel.CpmId
	serviceCtx.ServiceConfig.SourceToken = ""
	serviceCtx.ServiceConfig.SourcePacket = devPacket
	serviceCtx.ServiceConfig.DestinationData = storeSyncModel

	err := notifyCPMID(serviceCtx)
	if err != nil {
		return err, apiTaskProcessorResultModel
	}

	apiTaskProcessorResultModel.IsEPSync = true

	return nil, apiTaskProcessorResultModel
}

func notifyCPMID(serviceCtx *pcservices.ServiceContext) error {

	dtcollOnlineDevicesService := &pcservices.DtcollOnlineDevicesService{}
	dtcollOnlineDevicesService.ServiceContext = serviceCtx

	filterCPMIDService := &pcservices.FilterCPMIDService{}
	filterCPMIDService.ServiceContext = serviceCtx

	packetbldService := &pcservices.PacketbldService{}
	packetbldService.ServiceContext = serviceCtx

	senderService := &pcservices.SenderService{}
	senderService.ServiceContext = serviceCtx

	dttfGetDevMapService := &pcservices.DttfGetDevMapService{}
	dttfGetDevMapService.ServiceContext = serviceCtx

	dtcollOnlineDevicesService.NextHandler = filterCPMIDService
	filterCPMIDService.NextHandler = dttfGetDevMapService
	dttfGetDevMapService.NextHandler = packetbldService
	packetbldService.NextHandler = senderService

	err := dtcollOnlineDevicesService.Handle(serviceCtx)
	if err != nil {
		return err
	}

	return nil

}
