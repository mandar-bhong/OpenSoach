package helper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

var BaseDir string

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
