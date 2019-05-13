package processor

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
)

func GetOnlineDeviceTokens(msg string) (string, error) {

	tokens := GetTokens()

	isJsonSucc, jsonData := ghelper.ConvertToJSON(&tokens)

	if isJsonSucc == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert packet to json data", nil)
		return "", fmt.Errorf("Unable to convert to json packet")
	}

	return jsonData, nil

}
