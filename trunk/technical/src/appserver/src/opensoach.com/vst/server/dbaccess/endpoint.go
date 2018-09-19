package dbaccess

import (
	"github.com/jmoiron/sqlx"
	dbmgr "opensoach.com/core/manager/db"
	hktmodels "opensoach.com/vst/models"
	"opensoach.com/vst/server/constants"
	"opensoach.com/vst/server/constants/dbquery"
	lmodels "opensoach.com/vst/server/models"
)

func GetDBTransaction(dbconn string) (error, *sqlx.Tx) {
	ctx := dbmgr.InsertTxContext{}
	return ctx.GetTransaction(dbconn)
}

func EPGetInstanceDB(dbConn string, cpmid, deviceid int64) (error, string) {
	selDBCtx := dbmgr.SelectContext{}
	data := ""
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_DB_INST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = &data
	selErr := selDBCtx.Get(cpmid)
	if selErr != nil {
		return selErr, ""
	}

	return nil, data
}

func EPGetDeviceServicePoints(dbConn string, cpmid int64, deviceid int64) (error, *[]hktmodels.DBEPSPDataModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBEPSPDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_DEVICE_SP
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(cpmid, deviceid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPGetSPAuthCodes(dbConn string, cpmid int64, spid int64) (error, *[]string) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]string{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_SP_OPERATOR
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(cpmid, spid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPGetSPServConf(dbConn string, cpmid int64, spid int64) (error, *[]hktmodels.DBEPSPServConfDataModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBEPSPServConfDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_SP_SERV_CONF
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(cpmid, spid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPInsertServiceInstanceData(dbConn string,
	dbSerInstDataRowModel hktmodels.DBServiceInstanceTxDataRowModel) error {

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = dbSerInstDataRowModel
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_NODE_SERVICE_INST_TXN_TBL
	insertErr := insDBCtx.Insert()

	return insertErr
}

func EPInsertComplaintData(dbConn string,
	dbComplaintInsertRowModel hktmodels.DBComplaintInsertRowModel) error {

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = dbComplaintInsertRowModel
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_HKT_SP_COMPLAINT_TBL
	insertErr := insDBCtx.Insert()

	return insertErr
}

func EPInsertFeedbackData(dbConn string,
	dbFeedbackInsertRowModel hktmodels.DBFeedbackInsertRowModel) error {

	insDBCtx := dbmgr.InsertContext{}
	insDBCtx.DBConnection = dbConn
	insDBCtx.Args = dbFeedbackInsertRowModel
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_NODE_FEEDBACK_TBL
	insertErr := insDBCtx.Insert()

	return insertErr
}

func EPUpdateDeviceBatteryLevelData(dbConn string,
	dbDevStatusBatteryLevelUpdateDataModel hktmodels.DBDevStatusBatteryLevelUpdateDataModel) error {

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = dbDevStatusBatteryLevelUpdateDataModel
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_NODE_DEV_STATUS_TBL
	updateErr := updateCtx.Update()

	return updateErr
}

func EPUpdateDeviceConnectionStatusData(dbConn string,
	dbDevStatusConnectionStateUpdateDataModel hktmodels.DBDevStatusConnectionStateUpdateDataModel) error {

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = dbDevStatusConnectionStateUpdateDataModel
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_NODE_DEV_STATUS_TBL
	updateErr := updateCtx.Update()

	return updateErr
}

func EPGetVehicleId(dbConn string, vehicleNo string) (error, *[]hktmodels.DBSplVstVehicleMasterTableRowModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplVstVehicleMasterTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_VEHICLE_ID_BY_VHL_NO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(vehicleNo)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPInsertVehicleData(tx *sqlx.Tx, insrtStruct hktmodels.DBVehicleInsertRowModel) (error, int64) {

	insDBCtx := dbmgr.InsertTxContext{}
	insDBCtx.Tx = tx
	insDBCtx.Args = insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_VST_VEHICLE_MASTER_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func EPInsertVstTokenData(tx *sqlx.Tx, insrtStruct hktmodels.DBTokenInsertRowModel) (error, int64) {

	insDBCtx := dbmgr.InsertTxContext{}
	insDBCtx.Tx = tx
	insDBCtx.Args = insrtStruct
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_SPL_VST_TOKEN_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func EPGetLastVhlTokenRecord(dbConn string) (error, *[]hktmodels.DBSplVstTokenRowModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplVstTokenRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_LAST_VEHICLE_RECORD
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPUpdateVehicleDetailsData(dbConn string, updtStruct hktmodels.DBVehicleDetailsUpdateModel) (error, int64) {

	updateCtx := dbmgr.UpdateDeleteContext{}
	updateCtx.DBConnection = dbConn
	updateCtx.Args = updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_VST_VEHICLE_MASTER_TBL
	updateErr := updateCtx.UpdateByFilter("VehicleNo")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func EPInsertServiceInstanceTxnData(tx *sqlx.Tx, dbSerInstDataRowModel hktmodels.DBServiceInstanceTxDataRowModel) (error, int64) {

	insDBCtx := dbmgr.InsertTxContext{}
	insDBCtx.Tx = tx
	insDBCtx.Args = dbSerInstDataRowModel
	insDBCtx.QueryType = dbmgr.AutoQuery
	insDBCtx.TableName = constants.DB_TABLE_NODE_SERVICE_INST_TXN_TBL
	insertErr := insDBCtx.Insert()
	if insertErr != nil {
		return insertErr, 0
	}
	return nil, insDBCtx.InsertID
}

func EPUpdateTokenMappingDetailsData(tx *sqlx.Tx, updtStruct hktmodels.DBTokenMappingDetailsUpdateModel) (error, int64) {

	updateCtx := dbmgr.UpdateDeleteTxContext{}
	updateCtx.Tx = tx
	updateCtx.Args = updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_VST_TOKEN_TBL
	updateErr := updateCtx.UpdateByFilter("TokenId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func EPGetTokenMappingDetailsData(dbConn string, tokenID int64) (error, *[]hktmodels.DBSplVstTokenRowModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplVstTokenRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_TOKEN_MAPPING_DETAILS_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(tokenID)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPGetTokenList(dbConn string) (error, *[]hktmodels.DBEPSPVhlTokenDataModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBEPSPVhlTokenDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_VHL_TOKEN_LIST
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPUpdateTokenStateData(tx *sqlx.Tx, updtStruct hktmodels.DBTokenStateUpdateModel) (error, int64) {

	updateCtx := dbmgr.UpdateDeleteTxContext{}
	updateCtx.Tx = tx
	updateCtx.Args = updtStruct
	updateCtx.QueryType = dbmgr.AutoQuery
	updateCtx.TableName = constants.DB_SPL_VST_TOKEN_TBL
	updateErr := updateCtx.UpdateByFilter("TokenId")
	if updateErr != nil {
		return updateErr, 0
	}
	return nil, updateCtx.AffectedRows
}

func EPGetTokenDataById(dbConn string, tokenID int64) (error, *[]hktmodels.DBEPSPVhlTokenDataModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBEPSPVhlTokenDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_VHL_TOKEN_BY_TOKEN_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(tokenID)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPGetConfigListByTokenId(dbConn string, tokenid int64) (error, *[]hktmodels.DBTokenConfigBriefDataModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBTokenConfigBriefDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_CONFIG_LIST_BY_TOKEN_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(tokenid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}

func EPGetVehicleDetailsDataByVhlNo(dbConn string, vehicleNo string) (error, *[]lmodels.PacketVehicleDetailsData) {

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.PacketVehicleDetailsData{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_EP_PROC_GET_VEHICLE_DETAILS_BY_VHL_NO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(vehicleNo)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}
