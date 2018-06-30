package product

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	repo "opensoach.com/spl/api/repository"
	"opensoach.com/spl/api/webserver/product/dbaccess"
)

var SUB_MODULE_NAME = "SPL.Product"

type ProductService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (ProductService) GetProductList() (bool, interface{}) {

	dbErr, listData := dbaccess.DBSplMasterProductTableRowModelSelectAll(repo.Instance().Context.Master.DBConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched product table data.")

	return true, listData

}

func (ProductService) GetDbInstanceList() (bool, interface{}) {

	dbErr, listData := dbaccess.GetDbinstanceList(repo.Instance().Context.Master.DBConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched database instance table data.")

	return true, listData

}
