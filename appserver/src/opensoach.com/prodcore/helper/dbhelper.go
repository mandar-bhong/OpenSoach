package helper

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	dbmgr "opensoach.com/core/manager/db"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
)

func GetMasterConfiguration(config *gmodels.ConfigDB) (error, *[]gmodels.DBMasterConfigRowModel) {
	return GetConfiguration(config, pcconst.QUERY_GET_MASTER_CONFIGURATION)
}

func GetConfiguration(config *gmodels.ConfigDB, dbquery string) (error, *[]gmodels.DBMasterConfigRowModel) {

	configRows := &[]gmodels.DBMasterConfigRowModel{}

	selCtx := dbmgr.SelectContext{}
	selCtx.DBConnection = config.ConnectionString
	selCtx.Query = dbquery
	selCtx.QueryType = dbmgr.Query
	selCtx.Dest = configRows

	selErr := selCtx.SelectAll()

	if selErr != nil {
		fmt.Printf("DB Error %#+v \n", selErr.Error())
		return selErr, nil
	}

	return nil, configRows
}

func VerifyDBConnection(dbconfig *gmodels.ConfigDB) error {

	_, dbErr := sqlx.Connect(dbconfig.DBDriver, dbconfig.ConnectionString)

	if dbErr != nil {
		return dbErr
	}

	return nil
}
