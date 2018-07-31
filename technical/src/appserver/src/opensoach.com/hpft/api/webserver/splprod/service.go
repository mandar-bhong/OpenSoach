package splprod

import (
	"opensoach.com/core/logger"
	repo "opensoach.com/hpft/api/repository"
	"opensoach.com/hpft/api/webserver/splprod/dbaccess"
	gmodels "opensoach.com/models"
)

var SUB_MODULE_NAME = "HPFT.API.SplProd"

type SplprodService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service SplprodService) GetBaseUrl() (bool, interface{}) {

	dbErr, baseUrl := dbaccess.GetSplBaseUrl(repo.Instance().Context.ProdMst.DBConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while validating user.", dbErr)

		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE
		return false, errModel
	}

	dbRecord := *baseUrl

	if len(dbRecord) < 1 {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Info, "Base Url Not Found")
		errModel := gmodels.APIResponseError{}
		errModel.Code = gmodels.MOD_OPER_ERR_DATABASE_RECORD_NOT_FOUND
		return false, errModel
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched hkt base url")
	return true, dbRecord[0]
}
