package helper

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetExeFolder() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	return dir
}

func ReadFileContent(filePath ...string) (bool, []byte, string) {

	var fileFullPath string
	fileFullPath = ""

	if len(filePath) == 0 {
		emptyByte := []byte{}
		return false, emptyByte, "No parameter available"
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
		errorString := fmt.Sprintf("Read Error: %s, FilePath: %s ", readError.Error(), fileFullPath)
		return false, []byte{}, errorString
	}

	return true, byteData, ""
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

func ConvertFromJSONBytes(jsonData []byte, pConvertType interface{}) bool {
	err := json.Unmarshal(jsonData, pConvertType)

	if err != nil {
		return false
	}

	return true
}

func CreateToken() (bool, string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		//logger.Log(MODULENAME, logger.ERROR, "createSessionToken:Unable to create session token. Error: %s", err.Error())
		return false, ""
	}

	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return true, uuid
}
