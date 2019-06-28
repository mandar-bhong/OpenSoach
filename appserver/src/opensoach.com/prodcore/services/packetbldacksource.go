package services

import (
	//pcmodels "opensoach.com/prodcore/models"
	"errors"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	//ghelper "opensoach.com/core/helper"
	pcconst "opensoach.com/prodcore/constants"
	pcmodels "opensoach.com/prodcore/models"
	//"opensoach.com/core/logger"
)

//PacketbldAckSourceService This service get online devices token
type PacketbldAckSourceService struct {
	*ServiceContext
	NextHandler IHandler
}

//Handle Creation of packet to ack source
func (r *PacketbldAckSourceService) Handle(serctx *ServiceContext) error {

	epPacket := pcmodels.TaskEPPacketSendDataModel{}
	epPacket.Token = r.ServiceContext.ServiceConfig.SourceToken

	devPacket := &gmodels.DevicePacket{}
	devPacket.Header = gmodels.DeviceHeaderData{}
	devPacket.Header.SeqID = serctx.ServiceConfig.SourcePacket.Header.SeqID
	devPacket.Header.Category = pcconst.DEVICE_CMD_CAT_ACK
	devPacket.Payload = r.ServiceConfig.AckData

	isSuccess, packetJson := ghelper.ConvertToJSON(devPacket)

	if isSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while convert json packet in acksource", nil)
		return errors.New("Error occured while converto to json packet in acksource")
	}

	epPacket.Packet = packetJson

	r.ServiceResult.AckPacket = append(r.ServiceResult.AckPacket, epPacket)

	if r.NextHandler != nil {
		err := r.NextHandler.Handle(serctx)
		return err
	}

	return nil
}
