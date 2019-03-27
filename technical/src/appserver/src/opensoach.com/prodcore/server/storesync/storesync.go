package storesync

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
	"opensoach.com/prodcore/server/dbaccess"
	pcservices "opensoach.com/prodcore/services"
)

var SUB_MODULE_NAME = "ProdCore.Server.StoreSync"

func GetChanges(ctx *pcmodels.DevicePacketProccessExecution, dbConnections map[int]string, syncReq pcmodels.StoreSyncGetRequestModel) (error, *pcmodels.StoreSyncGetResponseModel) {

	dbConn := dbConnections[gmodels.DB_CONNECTION_NODE]

	dbErr, syncConfigData := dbaccess.GetSyncConfig(dbConn, syncReq.StoreName)
	if dbErr != nil {
		logger.Context().WithField("Sync Config Request", syncReq).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get sync config.", dbErr)
		return dbErr, nil
	}

	switch syncConfigData.DataSourceType {
	case gmodels.DB_CONNECTION_MASTER:
		dbConn = dbConnections[gmodels.DB_CONNECTION_MASTER]
		break
	case gmodels.DB_CONNECTION_NODE:
		dbConn = dbConnections[gmodels.DB_CONNECTION_NODE]
		break
	}

	if syncReq.FilterHandler != nil {
		err := syncReq.FilterHandler(ctx, syncConfigData, &syncReq)
		if err != nil {
			logger.Context().WithField("DB server filter:", syncConfigData).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to apply server filter.", err)
			return err, nil
		}
	}

	dbErr, tableData := dbaccess.GetTableData(dbConn, syncConfigData.SelectQry, syncReq.QueryParams)
	if dbErr != nil {
		logger.Context().WithField("Table Data Request", syncConfigData).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get table data.", dbErr)
		return dbErr, nil
	}

	dbErr, count := dbaccess.GetTableDataCount(dbConn, syncConfigData.SelectCountQry, syncReq.QueryParams)
	if dbErr != nil {
		logger.Context().WithField("Table Data Request", syncConfigData).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get table data count", dbErr)
		return dbErr, nil
	}

	countData := *count

	storeSyncGetResponseModel := &pcmodels.StoreSyncGetResponseModel{}
	storeSyncGetResponseModel.Data = tableData
	storeSyncGetResponseModel.Count = countData.Count
	storeSyncGetResponseModel.UpdatedOn = count.MaxUpdatedOn
	storeSyncGetResponseModel.StoreName = syncReq.StoreName

	return nil, storeSyncGetResponseModel
}

//TODO Add notification logic
func ApplyChanges(dbConn string, syncReq pcmodels.StoreSyncApplyRequestModel) (error, *pcmodels.StoreSyncApplyResponseModel) {

	dbErr, syncConfigData := dbaccess.GetSyncConfig(dbConn, syncReq.StoreName)
	if dbErr != nil {
		logger.Context().WithField("Sync Config Request", syncReq).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get sync config", dbErr)
		return dbErr, nil
	}

	//To Do - handle req data type if []map/[]IStoreSync
	err, list := syncReq.GetDataItems()
	if err != nil {
		logger.Context().WithField("Sync Request", syncReq).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to IStoreSync list", err)
		return err, nil
	}

	for _, each := range list {

		dbErr, count := dbaccess.GetTableDataByUuid(dbConn, syncConfigData.HasQuery, each["uuid"].(string))
		if dbErr != nil {
			logger.Context().WithField("Uuid", each["uuid"]).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get uuid  count", dbErr)
			return dbErr, nil
		}

		if count > 0 {
			dbErr, _ := dbaccess.UpdateTableData(dbConn, syncConfigData.UpdateQry, each)
			if dbErr != nil {
				logger.Context().
					WithField("Update table data", each).
					WithField("Update query", syncConfigData.UpdateQry).
					LogError(SUB_MODULE_NAME, logger.Normal, "Failed to update table data ", dbErr)
				return dbErr, nil
			}
		} else {
			dbErr := dbaccess.InsertTableData(dbConn, syncConfigData.InsertQry, each)
			if dbErr != nil {
				logger.Context().
					WithField("Insert table data", each).
					WithField("Update query", syncConfigData.InsertQry).
					LogError(SUB_MODULE_NAME, logger.Normal, "Failed to insert table data ", dbErr)
				return dbErr, nil
			}
		}
	}

	storeSyncApplyResponseModel := &pcmodels.StoreSyncApplyResponseModel{}

	return nil, storeSyncApplyResponseModel
}

func ApplyChangesNotify(dbConn string, syncReq pcmodels.StoreSyncApplyRequestModel, devPacket *gmodels.DevicePacket, Token string,
	repo pcmodels.Repo) (error, *pcmodels.StoreSyncApplyResponseModel) {

	deviceCommandAck := gmodels.DeviceCommandAck{}
	deviceCommandAck.Ack = true

	serviceCtx := &pcservices.ServiceContext{}
	serviceCtx.Repo = repo
	serviceCtx.ServiceConfig.SourcePacket = devPacket
	serviceCtx.ServiceConfig.SourceToken = Token

	err, resp := ApplyChanges(dbConn, syncReq)
	if err != nil {
		logger.Context().WithField("Sync Req", syncReq).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to apply sync changes.", err)
		deviceCommandAck.Ack = false

		serviceCtx.ServiceConfig.AckData = deviceCommandAck
		notifyErr := NotifyAck(serviceCtx)
		if notifyErr != nil {
			logger.Context().WithField("Service Context", serviceCtx).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to notify apply sync changes.", notifyErr)
		}
		return err, resp
	} else {
		storeSyncModel := pcmodels.StoreSyncModel{}
		storeSyncModel.StoreName = syncReq.StoreName
		serviceCtx.ServiceConfig.AckData = deviceCommandAck
		serviceCtx.ServiceConfig.DestinationData = storeSyncModel

		err = NotifyCPMID(serviceCtx)
		if err != nil {
			logger.Context().WithField("Service Context", serviceCtx).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to notify apply sync changes.", err)
			return err, resp
		}
	}

	return err, resp

}
