package helper

import (
	"crypto/rand"
	"fmt"
	mrand "math/rand"

	"opensoach.com/core/logger"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateToken(length int, prepend string) string {

	b := make([]byte, length)
	_, err := rand.Read(b)

	if err != nil {
		logger.Context().WithField("Error", err).Log("", logger.Server, logger.Error, "Unable to create Task token. Switchin algo")

		b := make([]rune, length)
		for i := range b {
			b[i] = letterRunes[mrand.Intn(len(letterRunes))]
		}
		randNum := string(b)
		return prepend + randNum
	}

	uuid := fmt.Sprintf("%s%X", prepend, b[0:])

	return uuid
}

func GenerateAPIToken() string {
	return GenerateToken(16, "API")
}

func GenerateDeviceToken() string {
	return GenerateToken(8, "Dev")
}

func GenerateTaskToken() string {
	return GenerateToken(8, "TaskToken")
}

func GenerateUUID() string {
	return GenerateToken(16, "DB")
}

func GenerateDeviceUserToken() string {
	return GenerateToken(8, "DU")
}

func GenerateUserOtp() string {
	return GenerateToken(4, "")
}
