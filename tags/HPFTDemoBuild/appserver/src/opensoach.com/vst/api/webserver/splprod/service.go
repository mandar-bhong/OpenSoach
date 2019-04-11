package splprod

import (
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	repo "opensoach.com/vst/api/repository"
	"opensoach.com/vst/api/webserver/splprod/dbaccess"
)

var SUB_MODULE_NAME = "VST.API.SplProd"

type SplprodService struct {
	ExeCtx *gmodels.ExecutionContext
}

func (service SplprodService) GetBaseUrl() (bool, interface{}) {

	dbErr, baseUrl := dbaccess.GetSplBaseUrl(repo.Instance().Context.ProdMst.DBConn)
	if dbErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Database error occured while getting base url.", dbErr)

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

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully fetched base url")
	return true, dbRecord[0]
}
