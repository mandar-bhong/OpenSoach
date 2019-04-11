package manager

import (
	"opensoach.com/vst/constants"

	"opensoach.com/vst/server/processor"
	"opensoach.com/vst/server/processor/endpoint"

	gmodels "opensoach.com/models"
	hktconst "opensoach.com/vst/constants"
)

func RegisterHandler(hkthandler map[string]interface{}) {
	hkthandler[constants.TASK_HANDLER_END_POINT_TO_SERVER_KEY] = ProcessEndPointReceivedPacket

	hkthandler[gmodels.TASK_HKT_EP_CONNECTED] = endpoint.ProcessDeviceConnected
	hkthandler[gmodels.TASK_HKT_EP_DISCONNECTED] = endpoint.ProcessDeviceDisConnected

	hkthandler[hktconst.TASK_HANDLER_HKT_API_CONTROLLER] = processor.APITaskController

}
