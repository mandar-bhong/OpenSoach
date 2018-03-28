package dbaccess

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	dbmgr "opensoach.com/core/manager/db"

	lmodels "opensoach.com/spl/models"
)

var dbDriverName = "mysql"

func GetUserProducts(dbEngine *sqlx.DB, userid int64) (error, *[]lmodels.DBProductBriefRowModel) {

	data := &[]lmodels.DBProductBriefRowModel{}
	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.Engine = dbEngine
	selDBCtx.Query = "sp_mst_get_usr_products"
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(userid)

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}

func GetProductDB(dbEngine *sqlx.DB, cpmid int64) (error, *[]lmodels.DBProductBriefRowModel) {

	data := &[]lmodels.DBProductBriefRowModel{}
	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.Engine = dbEngine
	selDBCtx.Query = "sp_mst_get_usr_products"
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(cpmid)

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}
