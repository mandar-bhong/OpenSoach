package dbaccess

import (
	dbmgr "opensoach.com/core/manager/db"

	lmodels "opensoach.com/spl/models"
)

var dbDriverName = "mysql"

func GetUserProducts(dbConn string, userid int64) (error, *[]lmodels.DBProductBriefRowModel) {

	data := &[]lmodels.DBProductBriefRowModel{}
	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = "sp_mst_get_usr_products"
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(userid)

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}

func GetProductDB(dbConn string, cpmid int64) (error, *[]lmodels.DBProductBriefRowModel) {

	data := &[]lmodels.DBProductBriefRowModel{}
	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = "sp_mst_get_usr_products"
	selDBCtx.Dest = data

	selErr := selDBCtx.Select(cpmid)

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}

func GetCustomerProductDetails(dbConn string, cpmid int64) (error, *lmodels.DBProductBriefRowModel) {
	data := &lmodels.DBProductBriefRowModel{}
	selDBCtx := dbmgr.SelectContext{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = QUERY_SELECT_CPM_DETAILS
	selDBCtx.Dest = data
	selDBCtx.QueryType = dbmgr.Query

	selErr := selDBCtx.Select(cpmid)

	if selErr != nil {
		return selErr, nil
	}

	return nil, data
}
