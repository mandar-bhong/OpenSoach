package services

import (
	"errors"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
	pcmodels "opensoach.com/prodcore/models"
)

type PacketbldService struct {
	*ServiceContext
	NextHandler IHandler
}

//Handle Send Available to EP
func (r *PacketbldService) Handle(serctx *ServiceContext) error {

	packetbldheaderService := PacketbldheaderService{}
	packetbldheaderService.ServiceContext = r.ServiceContext

	header := packetbldheaderService.Build()

	packet := pcmodels.TaskEPPacketSendDataModel{}

	for _, token := range r.ServiceRuntime.Tokens {

		isSuccess, deviceTokenModel, _ := pchelper.CacheGetDeviceInfo(r.Repo.Context.Master.Cache, token)
		if isSuccess == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get device token model from cache.", nil)
			return errors.New("Failed to get device token model from cache.")
		}

		val, ok := r.ServiceRuntime.DeviceLocationMap[deviceTokenModel.DevID]
		if ok {
			for _, each := range val {
				devPacket := gmodels.DevicePacket{}
				devPacket.Header = header
				devPacket.Header.SPID = each
				devPacket.Payload = r.ServiceConfig.DestinationData
				_, jsonpack := ghelper.ConvertToJSON(devPacket)

				packet.Token = token
				packet.Packet = jsonpack

				r.ServiceResult.DestinationPackets = append(r.ServiceResult.DestinationPackets, packet)
			}
		}

	}

	if r.NextHandler != nil {
		err := r.NextHandler.Handle(r.ServiceContext)
		return err
	}

	return nil

}
