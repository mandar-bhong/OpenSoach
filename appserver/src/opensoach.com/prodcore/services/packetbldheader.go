package services

import (
	gmodels "opensoach.com/models"
	pcconstants "opensoach.com/prodcore/constants"
)

type PacketbldheaderService struct {
	*ServiceContext
}

func (r *PacketbldheaderService) Build() gmodels.DeviceHeaderData {

	srccat := r.ServiceConfig.SourcePacket.Header.Category
	srccmd := r.ServiceConfig.SourcePacket.Header.CommandID

	header := gmodels.DeviceHeaderData{}

	switch {
	case srccat == pcconstants.DEVICE_CMD_CAT_DATA && srccmd == pcconstants.DEVICE_CMD_STORE_GET_SYNC:
		header.Category = pcconstants.DEVICE_CMD_CAT_NOTIFICATION
		header.CommandID = pcconstants.DEVICE_CMD_STORE_GET_SYNC
	case srccat == pcconstants.DEVICE_CMD_CAT_DATA && srccmd == pcconstants.DEVICE_CMD_STORE_APPLY_SYNC:
		header.Category = pcconstants.DEVICE_CMD_CAT_NOTIFICATION
		header.CommandID = pcconstants.DEVICE_CMD_STORE_APPLY_SYNC
	}

	return header
}
