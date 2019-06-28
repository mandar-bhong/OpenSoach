package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/hpft/constants/dbquery"
	lmodels "opensoach.com/hpft/models"
)

var SUB_MODULE_NAME = "HPFT.API.User.DB"

func GetDoctorUsers(dbConn string, cpmid int64) (error, *[]lmodels.DBUserInfoModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetDoctorUsers")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBUserInfoModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_SELECT_DOCTOR_USERS
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(cpmid)
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}
