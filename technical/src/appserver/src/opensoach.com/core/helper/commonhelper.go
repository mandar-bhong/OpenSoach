package helper

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	mrand "math/rand"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"opensoach.com/core/logger"
)

var BaseDir string
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GetExeFolder() string {

	if BaseDir != "" {
		return BaseDir
	}

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func ReadFileContent(filePath ...string) (error, []byte) {

	var fileFullPath string
	fileFullPath = ""

	if len(filePath) == 0 {
		emptyByte := []byte{}
		return errors.New("No parameter available"), emptyByte
	}

	if len(filePath) == 1 {
		fileFullPath = filePath[0]
	} else {
		for i := 0; i < len(filePath); i++ {
			fileFullPath = filepath.Join(fileFullPath, filePath[i])
		}
	}

	byteData, readError := ioutil.ReadFile(fileFullPath)

	if readError != nil {
		return readError, []byte{}
	}

	return nil, byteData
}

func ConvertToJSON(dataStruct interface{}) (bool, string) {
	dataBytes, err := json.Marshal(dataStruct)

	if err != nil {
		return false, ""
	}
	jsonData := string(dataBytes)

	return true, jsonData
}

func ConvertFromJSONString(jsonData string, pConvertType interface{}) bool {
	err := json.Unmarshal([]byte(jsonData), pConvertType)

	if err != nil {
		return false
	}

	return true
}

func ConvertFromJSONBytes(jsonData []byte, pConvertType interface{}) error {
	err := json.Unmarshal(jsonData, pConvertType)

	if err != nil {
		return err
	}

	return nil
}

func CreateToken() (bool, string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		//logger.Log(MODULENAME, logger.ERROR, "createSessionToken:Unable to create session token. Error: %s", err.Error())
		return false, ""
	}

	uuid := fmt.Sprintf("%X%X%X%X%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return true, uuid
}

func GenerateTaskToken() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)

	if err != nil {
		logger.Context().WithField("Error", err).Log("", logger.Server, logger.Error, "Unable to create Task token. Switchin algo")

		b := make([]rune, 8)
		for i := range b {
			b[i] = letterRunes[mrand.Intn(len(letterRunes))]
		}
		randNum := string(b)
		return "TaskToken" + randNum
	}

	uuid := fmt.Sprintf("TaskToken%X", b[0:])

	return uuid
}

func GetModelFields(model interface{}) []reflect.StructField {
	fields := make([]reflect.StructField, 0)
	modelval := reflect.ValueOf(model)
	modeltype := reflect.TypeOf(model)

	for i := 0; i < modeltype.NumField(); i++ {
		v := modelval.Field(i)
		t := modeltype.Field(i)
		switch v.Kind() {
		case reflect.Struct:
			if t.Type == reflect.TypeOf(time.Time{}) {
				fields = append(fields, t)
			} else {
				fields = append(fields, GetModelFields(v.Interface())...)
			}
		default:
			fields = append(fields, t)
		}
	}
	return fields
}

func GetCurrentTime() time.Time {
	currentTime := time.Now()
	return currentTime
}
