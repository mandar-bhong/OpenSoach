package dbaccess

import (
	dbmgr "opensoach.com/core/manager/db"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/splserver/constants"
	lmodels "opensoach.com/splserver/models"
)

func UpdateEPConnectionState(dbconn string, deviceID int64, state int) error {

	dbDeviceConnectionStateUpdateRowModel := lmodels.DBDeviceConnectionStateUpdateRowModel{}
	dbDeviceConnectionStateUpdateRowModel.DeviceID = deviceID
	dbDeviceConnectionStateUpdateRowModel.ConnectionState = state
	dbDeviceConnectionStateUpdateRowModel.StateSince = ghelper.GetCurrentTime()

	updateDeleteContext := dbmgr.UpdateDeleteContext{}
	updateDeleteContext.DBConnection = dbconn
	updateDeleteContext.QueryType = dbmgr.AutoQuery
	updateDeleteContext.TableName = constants.DB_TABLE_SPL_MASTER_DEVICE_STATUS_TBL
	updateDeleteContext.Args = dbDeviceConnectionStateUpdateRowModel

	return updateDeleteContext.Update()
}
