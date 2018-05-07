package master

import (
	"opensoach.com/core/logger"
	repo "opensoach.com/hkt/api/repository"
	"opensoach.com/hkt/api/webserver/master/dbaccess"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HKT.Master"

type MasterService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service MasterService) GetSpcTaskLib(id int64) (bool, interface{}) {

	dbErr, dbData := dbaccess.GetHktMasterSpcTaskLibTableById(repo.Instance().Context.ProdMst.DBConn, id)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *dbData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully get SpcTask info")
	return true, dbRecord[0]
}

func (service MasterService) GetTaskLib(id int64) (bool, interface{}) {

	dbErr, dbData := dbaccess.GetHktMasterTaskLibTableById(repo.Instance().Context.ProdMst.DBConn, id)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *dbData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully get SpcTask info")
	return true, dbRecord[0]
}

func (service MasterService) GetServConfType(id int64) (bool, interface{}) {

	dbErr, dbData := dbaccess.GetSplProdMasterServConfTypeTableById(repo.Instance().Context.ProdMst.DBConn, id)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *dbData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully get SpcTask info")
	return true, dbRecord[0]
}

func (service MasterService) GetSpCategory(id int64) (bool, interface{}) {

	dbErr, dbData := dbaccess.GetSplProdMasterSpCategoryTableById(repo.Instance().Context.ProdMst.DBConn, id)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *dbData

	if len(dbRecord) < 1 {
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully get SpcTask info")
	return true, dbRecord[0]
}
