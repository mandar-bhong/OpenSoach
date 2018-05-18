package manager

import (
	"opensoach.com/hkt/constants"

	"opensoach.com/hkt/server/processor/endpoint"

	gmodels "opensoach.com/models"
)

func RegisterHandler(hkthandler map[string]interface{}) {
	hkthandler[constants.TASK_HANDLER_END_POINT_TO_SERVER_KEY] = ProcessEndPointReceivedPacket

	hkthandler[gmodels.TASK_HKT_EP_CONNECTED] = endpoint.ProcessDeviceConnected
	hkthandler[gmodels.TASK_HKT_EP_DISCONNECTED] = endpoint.ProcessDeviceDisConnected
}
