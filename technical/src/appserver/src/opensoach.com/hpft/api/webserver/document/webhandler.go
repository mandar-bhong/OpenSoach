package document

import (
	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	"opensoach.com/hpft/api/constants"
	lhelper "opensoach.com/hpft/api/helper"
	lmodels "opensoach.com/hpft/api/models"
	repo "opensoach.com/hpft/api/repository"
	gmodels "opensoach.com/models"
)

func registerRouters(router *gin.RouterGroup) {
	router.GET(constants.API_DOCUMENT_DOWNLOAD, func(c *gin.Context) { lhelper.FileDownloadHandler(c, requestHandler) })
	router.POST(constants.API_DOCUMENT_UPLOAD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.POST(constants.API_DEVICE_DOCUMENT_UPLOAD, func(c *gin.Context) { lhelper.CommonWebRequestHandler(c, requestHandler) })
	router.GET(constants.API_DEVICE_DOCUMENT_DOWNLOAD, func(c *gin.Context) { lhelper.FileDownloadHandler(c, requestHandler) })
}

func requestHandler(pContext *gin.Context) (bool, interface{}) {
	var resultData interface{}
	isSuccess := false

	logger.Context().WithField("Request: ", pContext.Request.URL.Path).LogDebug(SUB_MODULE_NAME, logger.Normal, "API request received.")

	switch pContext.Request.URL.Path {

	case constants.API_DOCUMENT_DOWNLOAD:

		req := lmodels.APIDocumentDownloadRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionReqData(repo.Instance().Context, pContext, &req)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DocumentService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.DocumentDownload(req)

		break

	case constants.API_DOCUMENT_UPLOAD:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DocumentService{
			ExeCtx: successErrorData.(*gmodels.ExecutionContext),
		}.DocumentUpload(pContext)

		break

	case constants.API_DEVICE_DOCUMENT_UPLOAD:

		isPrepareExeSuccess, successErrorData := lhelper.PrepareDeviceExecutionData(repo.Instance().Context, pContext)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceDocumentService{
			ExeCtx: successErrorData.(*gmodels.DeviceExecutionContext),
		}.DeviceDocumentUpload(pContext)

		break

	case constants.API_DEVICE_DOCUMENT_DOWNLOAD:

		req := lmodels.APIDocumentDownloadRequest{}

		isPrepareExeSuccess, successErrorData := lhelper.PrepareDeviceExecutionReqData(repo.Instance().Context, pContext, &req)

		if isPrepareExeSuccess == false {
			logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while preparing execution data.")
			return false, successErrorData
		}

		isSuccess, resultData = DeviceDocumentService{
			ExeCtx: successErrorData.(*gmodels.DeviceExecutionContext),
		}.DeviceDocumentDownload(req)

		break

	}

	return isSuccess, resultData
}
