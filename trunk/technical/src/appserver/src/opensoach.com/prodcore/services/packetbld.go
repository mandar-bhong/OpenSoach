package services

import (
	"errors"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
	pchelper "opensoach.com/prodcore/helper"
	pcmodels "opensoach.com/prodcore/models"
)

type PacketbldService struct {
	*ServiceContext
	NextHandler IHandler
}

//Handle Send Available to EP
func (r *PacketbldService) Handle(serctx *ServiceContext) error {

	packet := pcmodels.TaskEPPacketSendDataModel{}

	for _, token := range r.ServiceRuntime.Tokens {

		if token != r.ServiceConfig.SourceToken {

			isSuccess, deviceTokenModel, _ := pchelper.CacheGetDeviceInfo(r.Repo.Context.Master.Cache, token)
			if isSuccess == false {
				logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Failed to get device token model from cache.", nil)
				return errors.New("Failed to get device token model from cache.")
			}

			val, ok := r.ServiceRuntime.DeviceLocationMap[deviceTokenModel.DevID]
			if ok {
				for _, each := range val {
					devPacket := gmodels.DevicePacket{}
					devPacket.Header = gmodels.DeviceHeaderData{}
					devPacket.Header.Category = pcconst.DEVICE_CMD_STORE_APPLY_SYNC
					devPacket.Header.SPID = each
					devPacket.Payload = r.ServiceConfig.DestinationData
					_, jsonpack := ghelper.ConvertToJSON(devPacket)

					packet.Token = token
					packet.Packet = jsonpack

					r.ServiceResult.DestinationPackets = append(r.ServiceResult.DestinationPackets, packet)
				}
			}
		}
	}

	if r.NextHandler != nil {
		err := r.NextHandler.Handle(r.ServiceContext)
		return err
	}

	return nil

}
