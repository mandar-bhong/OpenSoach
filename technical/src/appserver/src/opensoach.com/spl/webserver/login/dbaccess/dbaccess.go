package dbaccess

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	dbmgr "opensoach.com/core/manager/db"

	lmodels "opensoach.com/spl/models"
)

var SUB_MODULE_NAME = "SPL.Login.DB"

func ValidateAuth(dbEngine *sqlx.DB, username, password string) (error, *[]lmodels.DBMasterUserRowModel) {
	filter := lmodels.AuthRequest{}
	filter.UserName = username
	filter.Password = password
	data := &[]lmodels.DBMasterUserRowModel{}
	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.Engine = dbEngine
	selDBCtx.Query = QUERY_MUST_CHECK_USER_LOGIN
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selDBCtx.TableName = lmodels.DB_TABLE_USER_TBL

	selErr := selDBCtx.SelectByFilter(filter, "usr_name", "usr_password")

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}

func GetUserAuthInfo(dbEngine *sqlx.DB, prodcode string) (error, *[]lmodels.DBUserAuthInfo) {
	filter := lmodels.DBUserAuthInfo{}
	filter.ProdCode = prodcode
	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBUserAuthInfo{}
	selDBCtx.Engine = dbEngine
	selDBCtx.Query = QUERY_GET_USER_AUTH_INFO
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selDBCtx.TableName = lmodels.DB_TABLE_PRODUCT_TBL
	selErr := selDBCtx.SelectByFilter(filter, "prod_code")
	if selErr != nil {
		return selErr, &[]lmodels.DBUserAuthInfo{}
	}
	return nil, data
}
