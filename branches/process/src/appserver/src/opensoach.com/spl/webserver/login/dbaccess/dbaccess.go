package dbaccess

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	dbmgr "opensoach.com/core/manager/db"

	lmodels "opensoach.com/spl/models"
)

func ValidateLogin(conn string, username string, password string) (error, *[]lmodels.DBMasterUserRowModel) {

	dbEngine, connErr := sqlx.Connect(lmodels.DB_DRIVER_NAME, conn)

	if connErr != nil {
		return connErr, nil
	}

	data := &[]lmodels.DBMasterUserRowModel{}
	selDBCtx := dbmgr.SelectProcContext{}
	selDBCtx.Engine = dbEngine
	selDBCtx.SPName = "sp_mst_chk_user_login"
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(username, password)

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}
