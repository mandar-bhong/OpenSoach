package processor

import (
	"strconv"

	repo "opensoach.com/hkt/endpoint/repository"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
	pcepproc "opensoach.com/prodcore/endpoint/processor"
	pchelper "opensoach.com/prodcore/helper"
)

var chnIDAuthData map[int]string

func init() {
	chnIDAuthData = make(map[int]string)
}

func Init() {

}

func GetHeaderKey(packetBytes []byte) (error, string) {

	err, deviceHeader := pcepproc.DecodeHeader(packetBytes)

	if err != nil {
		return err, ""
	}

	key := pchelper.GetDeviceCmdKeyFromHeader(deviceHeader)
	return nil, key
}

func GetAuthKey() string {
	return strconv.Itoa(pcconst.DEVICE_CMD_CAT_DEVICE_VALIDATION) + "_" +
		strconv.Itoa(pcconst.DEVICE_CMD_DEVICE_AUTH) + "_" +
		strconv.Itoa(pcconst.DEVICE_CMD_KEY)
}

func PreProcessExecutor(epModel *gmodels.PacketProcessingTaskModel) *gmodels.PacketProcessingTaskResult {

	packetProcessingResult := &gmodels.PacketProcessingTaskResult{}

	errHeaderKey, cmdKey := GetHeaderKey(epModel.Message)

	if errHeaderKey != nil {
		//Unable to decode header packet is invalid
		return packetProcessingResult
	}

	if cmdKey == GetAuthKey() { //Authorization
		return packetProcessingResult
	}

	token := chnIDAuthData[epModel.ChannelID]

	isKeyGetSuccess, _ := repo.Instance().Context.Master.Cache.Get(token)

	if isKeyGetSuccess == false {
		//Key is expired or not available need to stop communication
	}

	//epModel.AuthData = authData
	epModel.Token = token

	return packetProcessingResult
}
