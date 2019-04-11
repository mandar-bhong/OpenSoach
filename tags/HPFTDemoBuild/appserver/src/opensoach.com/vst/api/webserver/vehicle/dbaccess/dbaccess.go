package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/vst/constants"
	hktmodels "opensoach.com/vst/models"
)

var SUB_MODULE_NAME = "VST.API.Vehicle.DB"

func Insert(dbConn string, insrtStruct *hktmodels.DBVehicleInsertRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing vehicle insert.")

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = *insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_VST_VEHICLE_MASTER_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func VehicleTableSelectByID(dbConn string, vehicleid int64) (error, *[]hktmodels.DBSplVstVehicleMasterTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing VehicleTableSelectByID")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplVstVehicleMasterTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.QueryType = dbmgr.AutoQuery
	selDBCtx.TableName = constants.DB_SPL_VST_VEHICLE_MASTER_TBL
	selDBCtx.Dest = data
	selErr := selDBCtx.SelectById(vehicleid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func UpdateByFilter(dbConn string, updtStruct *hktmodels.DBVehicleUpdateRowModel) (error, int64) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing vehicle UpdateByFilter")

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = *updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_VST_VEHICLE_MASTER_TBL
	updateErr := updateCtx.UpdateByFilter("VehicleId", "CpmId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}
