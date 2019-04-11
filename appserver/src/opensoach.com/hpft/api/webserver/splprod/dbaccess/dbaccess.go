package dbaccess

import (
	"opensoach.com/core/logger"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/hpft/constants/dbquery"
	hktmodels "opensoach.com/hpft/models"
)

var SUB_MODULE_NAME = "HPFT.API.SplProd.DB"

func GetSplBaseUrl(dbConn string) (error, *[]hktmodels.DBSplBaseUrlDataModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetHktBaseUrl")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]hktmodels.DBSplBaseUrlDataModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_SPL_BASE_URL
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select()
	if selErr != nil {
		return selErr, nil
	}
	return nil, data
}
