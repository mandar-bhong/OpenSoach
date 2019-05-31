package master

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	repo "opensoach.com/vst/api/repository"
	"opensoach.com/vst/api/webserver/master/dbaccess"
)

var SUB_MODULE_NAME = "VST.API.Master"

type MasterService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service MasterService) GetSpcTaskLib(id int64) (bool, interface{}) {

	dbErr, dbData := dbaccess.GetHktMasterSpcTaskLibTableById(repo.Instance().Context.ProdMst.DBConn, id)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting spc task lib info.", dbErr)

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
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting task lib info.", dbErr)

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully get Task info")
	return true, dbRecord[0]
}

func (service MasterService) GetServConfType(id int64) (bool, interface{}) {

	dbErr, dbData := dbaccess.GetSplProdMasterServConfTypeTableById(repo.Instance().Context.ProdMst.DBConn, id)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting serv conf type info.", dbErr)

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully get ServConfType info")
	return true, dbRecord[0]
}

func (service MasterService) GetSpCategory(id int64) (bool, interface{}) {

	dbErr, dbData := dbaccess.GetSplProdMasterSpCategoryTableById(repo.Instance().Context.ProdMst.DBConn, id)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting sp cateqory info.", dbErr)

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully get Sp Category info")
	return true, dbRecord[0]
}
