package document

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hpft/api/models"
	gmodels "opensoach.com/models"
	pcconstants "opensoach.com/prodcore/constants"
	pcmodels "opensoach.com/prodcore/models"
	pcmgr "opensoach.com/prodcore/server/manager"
)

var SUB_MODULE_NAME = "HPFT.API.Document"

type DocumentService struct {
	ExeCtx *gmodels.ExecutionContext
}

type DeviceDocumentService struct {
	ExeCtx *gmodels.DeviceExecutionContext
}

func (service DocumentService) DocumentDownload(req lmodels.APIDocumentDownloadRequest) (bool, interface{}) {

	documentStoreDataModel := &pcmodels.DocumentStoreDataModel{}
	documentStoreDataModel.StorageType = pcconstants.DB_DOCUMENT_STORAGE_TYPE_FILE_SYSTEM
	documentStoreDataModel.DocumentID = req.Uuid
	documentStoreDataModel.DBContext = service.ExeCtx

	err, bytedata := pcmgr.DocumentStoreGet(documentStoreDataModel)
	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while downloading document.", err)
		return false, nil
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully downloaded document.")

	return true, bytedata
}

func (service DeviceDocumentService) DeviceDocumentDownload(req lmodels.APIDocumentDownloadRequest) (bool, interface{}) {

	documentStoreDataModel := &pcmodels.DocumentStoreDataModel{}
	documentStoreDataModel.StorageType = pcconstants.DB_DOCUMENT_STORAGE_TYPE_FILE_SYSTEM
	documentStoreDataModel.DocumentID = req.Uuid
	documentStoreDataModel.DBContext = service.ExeCtx

	err, documentdata := pcmgr.DocumentStoreGet(documentStoreDataModel)
	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while downloading document.", err)
		return false, nil
	}

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully downloaded document.")

	return true, documentdata
}

func (service DocumentService) DocumentUpload(pContext *gin.Context) (bool, interface{}) {

	form, err := pContext.MultipartForm()
	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while uploading document.", err)
		return false, nil
	}

	documentStoreDataModel := &pcmodels.DocumentStoreDataModel{}
	documentStoreDataModel.StorageType = pcconstants.DB_DOCUMENT_STORAGE_TYPE_FILE_SYSTEM
	documentStoreDataModel.FileData = form
	documentStoreDataModel.NestedPath = append(documentStoreDataModel.NestedPath, strconv.FormatInt(service.ExeCtx.SessionInfo.Product.CustProdID, 10))
	documentStoreDataModel.DBContext = service.ExeCtx

	err, UUID := pcmgr.DocumentStoreSave(documentStoreDataModel)
	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while uploading document.", err)
		return false, nil
	}

	uploadResponse := lmodels.APIDocumentUploadResponse{}
	uploadResponse.Uuid = UUID

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully uploaded document.")

	return true, uploadResponse
}

func (service DeviceDocumentService) DeviceDocumentUpload(pContext *gin.Context) (bool, interface{}) {

	form, err := pContext.MultipartForm()
	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while uploading document.", err)
		return false, nil
	}

	documentStoreDataModel := &pcmodels.DocumentStoreDataModel{}
	documentStoreDataModel.StorageType = pcconstants.DB_DOCUMENT_STORAGE_TYPE_FILE_SYSTEM
	documentStoreDataModel.FileData = form
	documentStoreDataModel.NestedPath = append(documentStoreDataModel.NestedPath, strconv.FormatInt(service.ExeCtx.DeviceSessionInfo.CpmID, 10))
	documentStoreDataModel.DBContext = service.ExeCtx
	documentStoreDataModel.DocumentID = form.Value["UUID"][0]

	err, UUID := pcmgr.DocumentStoreUpdate(documentStoreDataModel)
	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while uploading document.", err)
		return false, nil
	}

	uploadResponse := lmodels.APIDocumentUploadResponse{}
	uploadResponse.Uuid = UUID

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Successfully uploaded document.")

	return true, uploadResponse
}
