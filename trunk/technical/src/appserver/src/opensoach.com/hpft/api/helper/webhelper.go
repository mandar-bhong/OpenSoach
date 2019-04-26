package helper

import (
	"github.com/gin-gonic/gin"
	gcore "opensoach.com/core"
	pchelper "opensoach.com/prodcore/helper"
)

func PrepareExecutionData(osContext *gcore.Context, ginContext *gin.Context) (bool, interface{}) {
	return pchelper.PrepareExecutionData(osContext, ginContext)
}

func PrepareExecutionReqData(osContext *gcore.Context, ginContext *gin.Context, pClientReq interface{}) (bool, interface{}) {
	return pchelper.PrepareExecutionReqData(osContext, ginContext, pClientReq)
}

func PrepareDeviceExecutionData(osContext *gcore.Context, ginContext *gin.Context) (bool, interface{}) {
	return pchelper.PrepareDeviceExecutionData(osContext, ginContext)
}

func PrepareDeviceExecutionReqData(osContext *gcore.Context, ginContext *gin.Context, pClientReq interface{}) (bool, interface{}) {
	return pchelper.PrepareDeviceExecutionReqData(osContext, ginContext, pClientReq)
}

func PrepareDeviceUserExecutionData(osContext *gcore.Context, ginContext *gin.Context) (bool, interface{}) {
	return pchelper.PrepareDeviceUserExecutionData(osContext, ginContext)
}

func PrepareDeviceUserExecutionReqData(osContext *gcore.Context, ginContext *gin.Context, pClientReq interface{}) (bool, interface{}) {
	return pchelper.PrepareDeviceUserExecutionReqData(osContext, ginContext, pClientReq)
}

func CommonWebRequestHandler(pContext *gin.Context, requestHandlerFunc pchelper.RequestHandler) {
	pchelper.CommonWebRequestHandler(pContext, requestHandlerFunc)
}

func FileDownloadHandler(pContext *gin.Context, requestHandlerFunc pchelper.RequestHandler) {
	pchelper.FileDownloadHandler(pContext, requestHandlerFunc)
}
