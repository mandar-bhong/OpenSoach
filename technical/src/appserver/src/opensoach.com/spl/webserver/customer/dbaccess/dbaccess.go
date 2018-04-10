package dbaccess

import (
	"github.com/jmoiron/sqlx"
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/spl/constants"
	"opensoach.com/spl/constants/dbquery"
	lmodels "opensoach.com/spl/models"
)

func GetCustomerById(dbEngine *sqlx.DB, customerId int64) (error, *lmodels.DBCustomerInfoDataModel) {
	selDBCtx := dbmgr.SelectContext{}
	data := &lmodels.DBCustomerInfoDataModel{}
	selDBCtx.Engine = dbEngine
	selDBCtx.Query = dbquery.QUERY_GET_CUSTOMER_INFO_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selDBCtx.TableName = constants.DB_TABLE_CUSTOMER_TBL
	selErr := selDBCtx.Get(customerId)
	if selErr != nil {
		return selErr, &lmodels.DBCustomerInfoDataModel{}
	}
	return nil, data
}
