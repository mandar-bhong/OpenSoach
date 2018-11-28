package storesync

import (
	pcservices "opensoach.com/prodcore/services"
)

func NotifyCPMID(serviceCtx *pcservices.ServiceContext) error {

	packetbldAckSourceService := &pcservices.PacketbldAckSourceService{}
	packetbldAckSourceService.ServiceContext = serviceCtx

	dtcollOnlineDevicesService := &pcservices.DtcollOnlineDevicesService{}
	dtcollOnlineDevicesService.ServiceContext = serviceCtx

	filterCPMIDService := &pcservices.FilterCPMIDService{}
	filterCPMIDService.ServiceContext = serviceCtx

	packetbldService := &pcservices.PacketbldService{}
	packetbldService.ServiceContext = serviceCtx

	senderService := &pcservices.SenderService{}
	senderService.ServiceContext = serviceCtx

	dttfGetDevMapService := &pcservices.DttfGetDevMapService{}
	dttfGetDevMapService.ServiceContext = serviceCtx

	dtcollOnlineDevicesService.NextHandler = packetbldAckSourceService
	packetbldAckSourceService.NextHandler = filterCPMIDService
	filterCPMIDService.NextHandler = dttfGetDevMapService
	dttfGetDevMapService.NextHandler = packetbldService
	packetbldService.NextHandler = senderService

	err := dtcollOnlineDevicesService.Handle(serviceCtx)
	if err != nil {
		return err
	}

	return nil

}
