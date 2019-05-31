package manager

import (
	epproc "opensoach.com/hpft/endpoint/processor"
	pcconst "opensoach.com/prodcore/constants"
)

func RegisterTaskHandler(hkthandler map[string]interface{}) {

	hkthandler[pcconst.TASK_EP_SEND_PACKET] = epproc.SendEPPacketHandler
	hkthandler[pcconst.TASK_GET_ONLINE_DEVICES] = epproc.GetOnlineDeviceTokens

}
