package helper

import (
	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

func GetStoreTableStruct(packet []byte, config pcmodels.StoreConfigModel) (error, pcmodels.StoreSyncApplyRequestModel, *gmodels.DevicePacket) {

	reqModel := pcmodels.StoreSyncApplyRequestModel{}

	devPacket := &gmodels.DevicePacket{}
	devPacket.Payload = &reqModel

	convErr := ghelper.ConvertFromJSONBytes(packet, devPacket)
	if convErr != nil {
		return convErr, pcmodels.StoreSyncApplyRequestModel{}, nil
	}

	m := make([]map[string]interface{}, 0)
	reqModel.Data = &m

	convErr = ghelper.ConvertFromJSONBytes(packet, devPacket)
	if convErr != nil {
		return convErr, pcmodels.StoreSyncApplyRequestModel{}, nil
	}

	ss := *reqModel.Data.(*[]map[string]interface{})
	reqModel.Data = ss

	return nil, reqModel, devPacket

}
