package endpoint

import (
	"errors"
	"strings"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	repo "opensoach.com/hpft/server/repository"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
	"opensoach.com/prodcore/constants/dbquery"
	pchelper "opensoach.com/prodcore/helper"
	pcmodels "opensoach.com/prodcore/models"
	pcstoresync "opensoach.com/prodcore/server/storesync"
	pcservices "opensoach.com/prodcore/services"
)

func ProcessGetStoreSync(ctx *pcmodels.DevicePacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	packetProcessingResult.IsSuccess = true
	deviceCommandAck := gmodels.DeviceCommandAck{}
	deviceCommandAck.Ack = true

	reqModel := pcmodels.StoreSyncGetRequestModel{}
	reqModel.FilterHandler = AttachServerFilter
	reqModel.QueryHandler = AttachQueryHandler

	devPacket := &gmodels.DevicePacket{}
	devPacket.Payload = &reqModel

	deviceType := ctx.GetDeviceContextType()
	if deviceType == pcconst.DEVICE_TYPE_NONE {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting device context type", nil)
	}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devPacket)
	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json", convErr)
		deviceCommandAck.Ack = false
		packetProcessingResult.IsSuccess = false
	} else {

		dbConnections := make(map[int]string)

		dbConnections[gmodels.DB_CONNECTION_MASTER] = repo.Instance().Context.Master.DBConn
		dbConnections[gmodels.DB_CONNECTION_NODE] = ctx.InstanceDBConn

		err, data := pcstoresync.GetChanges(ctx, dbConnections, reqModel)
		if err != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting db changes", err)
			deviceCommandAck.Ack = false
			packetProcessingResult.IsSuccess = false
		}

		deviceCommandAck.Data = data
	}

	serviceCtx := &pcservices.ServiceContext{}
	serviceCtx.Repo = *repo.Instance()
	serviceCtx.ServiceConfig.SourcePacket = devPacket
	serviceCtx.ServiceConfig.AckData = deviceCommandAck
	serviceCtx.ServiceConfig.SourceToken = ctx.Token

	packetbldAckSourceService := &pcservices.PacketbldAckSourceService{}
	packetbldAckSourceService.ServiceContext = serviceCtx

	senderService := &pcservices.SenderService{}
	senderService.ServiceContext = serviceCtx

	packetbldAckSourceService.NextHandler = senderService

	serviceErr := packetbldAckSourceService.Handle(serviceCtx)
	if serviceErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while implementing services", serviceErr)
		packetProcessingResult.IsSuccess = false
		return
	}

}

func ProcessApplyStoreSync(ctx *pcmodels.DevicePacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	packetProcessingResult.IsSuccess = true
	deviceCommandAck := gmodels.DeviceCommandAck{}
	deviceCommandAck.Ack = true

	deviceType := ctx.GetDeviceContextType()
	if deviceType == pcconst.DEVICE_TYPE_NONE {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting device context type", nil)
	}

	convErr, reqModel, devPacket := pchelper.GetStoreTableStruct(ctx.DevicePacket, pcmodels.StoreConfigModel{})
	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting store struct", convErr)
		packetProcessingResult.IsSuccess = false
		deviceCommandAck.Ack = false

		serviceCtx := &pcservices.ServiceContext{}
		serviceCtx.Repo = *repo.Instance()
		serviceCtx.ServiceConfig.SourcePacket = devPacket
		serviceCtx.ServiceConfig.SourceToken = ctx.Token
		serviceCtx.ServiceConfig.AckData = deviceCommandAck
		notifyErr := pcstoresync.NotifyAck(serviceCtx)
		if notifyErr != nil {
			logger.Context().WithField("Service Context", serviceCtx).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to notify apply sync changes.", notifyErr)
		}

		return
	}

	// add cpmid in request data
	for i := 0; i < len(reqModel.Data.([]map[string]interface{})); i++ {
		reqModel.Data.([]map[string]interface{})[i]["cpm_id_fk"] = ctx.GetCPMID()
	}

	err, _ := pcstoresync.ApplyChangesNotify(ctx.InstanceDBConn, reqModel, devPacket, ctx.Token, *repo.Instance())

	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while apply sync changes", err)
		packetProcessingResult.IsSuccess = false
		return
	}

}

func AttachServerFilter(ctx *pcmodels.DevicePacketProccessExecution, filterModel *pcmodels.SyncConfigModel, request *pcmodels.StoreSyncGetRequestModel) error {

	queryDataModel := pcmodels.QueryDataModel{}
	isSuccess := ghelper.ConvertFromJSONString(filterModel.QueryData, &queryDataModel)
	if isSuccess == false {
		logger.Context().WithField("DB Server Filter", filterModel.QueryData).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to convert query data json.", nil)
		return errors.New("Unable to server parse query parameter form json data")
	}

	if request.QueryParams == nil {
		request.QueryParams = make(map[string]interface{})
	}

	deviceType := ctx.GetDeviceContextType()
	if deviceType == pcconst.DEVICE_TYPE_NONE {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting device context type", nil)
	}

	filterList := []string{}

	if deviceType == pcconst.DEVICE_TYPE_SHARED_DEVICE {
		filterList = queryDataModel.SharedDevice.Filters.Select
	} else if deviceType == pcconst.DEVICE_TYPE_USER_DEVICE {
		filterList = queryDataModel.UserDevice.Filters.Select
	}

	if len(filterList) > 0 {
		for _, each := range filterList {
			switch each {
			case "cpm":
				request.QueryParams["cpmid"] = ctx.GetCPMID()
				break
			case "updatedon":
				request.QueryParams["updatedon"] = request.UpdatedOn
				break
			case "usrid":
				_, usrid := ctx.GetDeviceUserID()
				request.QueryParams["usrid"] = usrid
				break
			}
		}
	}

	return nil
}

func AttachQueryHandler(ctx *pcmodels.DevicePacketProccessExecution, syncConfigData *pcmodels.SyncConfigModel, syncReq *pcmodels.StoreSyncGetRequestModel) (error, pcmodels.QueryModel) {

	queryModel := pcmodels.QueryModel{}

	deviceType := ctx.GetDeviceContextType()
	if deviceType == pcconst.DEVICE_TYPE_NONE {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting device context type", nil)
	}

	selectQueryDataModel := pcmodels.SelectQueryDataModel{}
	isSuccess := ghelper.ConvertFromJSONString(syncConfigData.SelectQry, &selectQueryDataModel)
	if isSuccess == false {
		logger.Context().WithField("select query:", syncConfigData.SelectQry).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to convert query data json.", nil)
		return errors.New("Unable to parse select query string from json data"), queryModel
	}

	switch deviceType {
	case pcconst.DEVICE_TYPE_USER_DEVICE:
		switch syncReq.StoreName {
		case pcconst.SYNC_STORE_PATIENT_CONF,
			pcconst.SYNC_STORE_PERSONAL_DETAILS, pcconst.SYNC_STORE_MEDICAL_DETAILS,
			pcconst.SYNC_STORE_ACTION_TXN, pcconst.SYNC_STORE_DOCTORS_ORDERS,
			pcconst.SYNC_STORE_TREATMENT, pcconst.SYNC_STORE_TREATMENT_DOC,
			pcconst.SYNC_STORE_PATHOLOGY, pcconst.SYNC_STORE_PATHOLOGY_DOC,
			pcconst.SYNC_STORE_ACTION, pcconst.SYNC_STORE_PATIENT_ADMISSION:
			queryModel.SelectQuery = strings.Replace(dbquery.QUERY_SELECT_SYNC_STORE_USER_DEVICE_QUERY, "$SyncQuery$", selectQueryDataModel.UserDeviceSelectQuery, 1)
			queryModel.SelectCountQuery = strings.Replace(dbquery.QUERY_SELECT_SYNC_STORE_USER_DEVICE_COUNT_QUERY, "$SyncQuery$", selectQueryDataModel.UserDeviceSelectQuery, 1)
			break
		case pcconst.SYNC_STORE_PATIENT_MONITOR_MAPPING_VIEW:
			queryModel.SelectQuery = strings.Replace(dbquery.QUERY_SELECT_SYNC_STORE_PATIENT_MONITOR_MAPPING_USER_DEVICE_QUERY, "$SyncQuery$", selectQueryDataModel.UserDeviceSelectQuery, 1)
			queryModel.SelectCountQuery = strings.Replace(dbquery.QUERY_SELECT_SYNC_STORE_PATIENT_MONITOR_MAPPING_USER_DEVICE_COUNT_QUERY, "$SyncQuery$", selectQueryDataModel.UserDeviceSelectQuery, 1)
			break
		default:
			queryModel.SelectQuery = selectQueryDataModel.UserDeviceSelectQuery
			queryModel.SelectCountQuery = syncConfigData.SelectCountQry
			break
		}
		break
	case pcconst.DEVICE_TYPE_SHARED_DEVICE:
		queryModel.SelectQuery = selectQueryDataModel.SharedDeviceSelectQuery
		queryModel.SelectCountQuery = strings.Replace(dbquery.QUERY_SELECT_SYNC_STORE_COUNT_QUERY, "$SyncQuery$", selectQueryDataModel.SharedDeviceSelectQuery, 1)
		break

	}

	return nil, queryModel
}
