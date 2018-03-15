package dbaccess

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	dbmgr "opensoach.com/core/manager/db"

	lmodels "opensoach.com/spl/models"
)

var dbDriverName = "mysql"

func GetUserProducts(conn string, userid int64) (error, *[]lmodels.DBProductBriefRowModel) {

	dbEngine, connErr := sqlx.Connect(lmodels.DB_DRIVER_NAME, conn)

	if connErr != nil {
		return connErr, nil
	}

	data := &[]lmodels.DBProductBriefRowModel{}
	selDBCtx := dbmgr.SelectProcContext{}
	selDBCtx.Engine = dbEngine
	selDBCtx.SPName = "sp_mst_get_usr_products"
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(userid)

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}
