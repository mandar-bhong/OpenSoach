package manager

import (
	gmodels "opensoach.com/models"
	"opensoach.com/splserver/processor"
)

var SUB_MODULE_NAME = "SPL.Server.Manager"

func RegisterHandler(handler map[string]interface{}) {
	handler[gmodels.TASK_HANDLER_API_SPL_CONTROLLER_KEY] = processor.APITaskController

	handler[gmodels.TASK_SPL_EP_CONNECTED] = processor.EndPointHandlerOnConnection
	handler[gmodels.TASK_SPL_EP_DISCONNECTED] = processor.EndPointHandlerOnDisConnection
}
