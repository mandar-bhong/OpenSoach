package manager

import (
	pcconst "opensoach.com/prodcore/constants"
	epproc "opensoach.com/vst/endpoint/processor"
)

func RegisterTaskHandler(hkthandler map[string]interface{}) {

	hkthandler[pcconst.TASK_EP_SEND_PACKET] = epproc.SendEPPacketHandler
	hkthandler[pcconst.TASK_GET_ONLINE_DEVICES] = epproc.GetOnlineDeviceTokens

}
