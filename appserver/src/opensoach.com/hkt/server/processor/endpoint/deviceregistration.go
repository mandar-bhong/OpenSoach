package endpoint

import (
	"fmt"

	//ghelper "opensoach.com/core/helper"
	lmodels "opensoach.com/hkt/server/models"
	gmodels "opensoach.com/models"
)

func ProcessDevReg(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	packetProcessingResult.IsSuccess = true

	de1 := &gmodels.DevicePacket{}
	de1.Header = gmodels.DeviceHeaderData{}
	de1.Header.Category = 3
	de1.Header.CommandID = 4

	de2 := &gmodels.DevicePacket{}
	de2.Header = gmodels.DeviceHeaderData{}
	de1.Header.Category = 5
	de1.Header.CommandID = 6

	fmt.Println("Assigning reslut")

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, de1)
	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, de2)

	fmt.Println("Returning ProcessDevReg")
}
