package dbaccess

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	dbmgr "opensoach.com/core/manager/db"

	lmodels "opensoach.com/spl/models"
)

func ValidateLogin(dbEngine *sqlx.DB, username string, password string) (error, *[]lmodels.DBMasterUserRowModel) {

	data := &[]lmodels.DBMasterUserRowModel{}
	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.Engine = dbEngine
	selDBCtx.Query = "sp_mst_chk_user_login"
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(username, password)

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}
