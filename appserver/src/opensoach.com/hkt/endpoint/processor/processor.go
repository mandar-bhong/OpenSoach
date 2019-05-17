package processor

import (
	"strconv"

	"opensoach.com/core/logger"
	repo "opensoach.com/hkt/endpoint/repository"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
	pcepproc "opensoach.com/prodcore/endpoint/processor"
	pchelper "opensoach.com/prodcore/helper"
)

var SUB_MODULE_NAME = "HKT.Endpoint.Manager.Processor"

var chnIDvsToken map[int]string
var tokenvsChnID map[string]int

func init() {
	chnIDvsToken = make(map[int]string)
	tokenvsChnID = make(map[string]int)
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
		strconv.Itoa(pcconst.DEVICE_CMD_DEVICE_AUTH)
}

func PreProcessExecutor(epModel *gmodels.PacketProcessingTaskModel) *gmodels.PacketProcessingTaskResult {

	packetProcessingResult := &gmodels.PacketProcessingTaskResult{}

	errHeaderKey, cmdKey := GetHeaderKey(epModel.Message)

	if errHeaderKey != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occurred while decoding header", errHeaderKey)
		return packetProcessingResult
	}

	if cmdKey == GetAuthKey() { //This is auth command hence skipping validation.
		packetProcessingResult.IsSuccess = true
		return packetProcessingResult
	}

	token := chnIDvsToken[epModel.ChannelID]

	isKeyGetSuccess, _ := repo.Instance().Context.Master.Cache.Get(token)

	if isKeyGetSuccess == false {
		logger.Context().WithField("Token", token).Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to get token from cache. Skipping further processing.")

		packetProcessingResult.StatusCode = gmodels.DEVICE_PROCESSING_AUTH_TOKEN_NOT_FOUND
		return packetProcessingResult
	}

	epModel.Token = token

	packetProcessingResult.IsSuccess = true

	return packetProcessingResult
}

func GetTokens() []string {

	tokenList := []string{}

	for token := range tokenvsChnID {
		tokenList = append(tokenList, token)
	}

	return tokenList
}
