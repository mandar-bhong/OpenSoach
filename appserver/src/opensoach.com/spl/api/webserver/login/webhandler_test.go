package login

import (
	"fmt"
	"os"

	//"fmt"

	"testing"

	"opensoach.com/spl/api/constants"
	lhelper "opensoach.com/spl/api/helper"

	"github.com/oliveagle/jsonpath"
)

const TEST_USER_LOGIN_VALID = `{"username":"admin@servicepoint.live","password":"admin","prodcode":"SPL_HKT"}`
const TEST_USER_LOGIN_INFO = `{"username":"admin@servicepoint.live","password":"admin","prodcode":"SPL_HKT"}`

func TestMain(m *testing.M) {

	lhelper.PrepareTestSetup()

	os.Exit(m.Run())
}

func Login(t *testing.T) (isSuccess bool, token string) {

	requestType := "POST"
	API := constants.API_USER_LOGIN
	jsonReqData := TEST_USER_LOGIN_VALID

	webResponse, jsonPathStruct := lhelper.ExecuteTestRequest(t, requestType, API, LoginRequestHandler, jsonReqData, token)

	t.Logf("\nRequest URI : %s \n", API)
	t.Logf("\nRequest Data : %s \n", jsonReqData)
	t.Logf("\nResponse received %v \n", webResponse)

	if res, err := jsonpath.JsonPathLookup(jsonPathStruct, "$.issuccess"); err != nil {
		t.Errorf("Recieved issuccess %v", err.Error())
		return false, ""
	} else if fmt.Sprintf("%v", res) != "true" {
		t.Errorf("Recieved issuccess %v", res)
	}

	res, err := jsonpath.JsonPathLookup(jsonPathStruct, "$.data.token")

	if err != nil {
		t.Errorf("Recieved issuccess %v", err.Error())
		return false, res.(string)
	}

	return true, res.(string)
}

func Test_UserLogin(t *testing.T) {

	requestType := "POST"
	API := constants.API_USER_LOGIN
	jsonReqData := TEST_USER_LOGIN_VALID
	token := ""

	webResponse, jsonPathStruct := lhelper.ExecuteTestRequest(t, requestType, API, requestHandler, jsonReqData, token)

	t.Logf("\nRequest URI : %s \n", API)
	t.Logf("\nRequest Data : %s \n", jsonReqData)
	t.Logf("\nResponse received %v \n", webResponse)

	if res, err := jsonpath.JsonPathLookup(jsonPathStruct, "$.issuccess"); err != nil {
		t.Errorf("Recieved issuccess %v", err.Error())
		return
	} else if fmt.Sprintf("%v", res) != "true" {
		t.Errorf("Recieved issuccess %v", res)
	}

}

func Test_UserInfo(t *testing.T) {
	requestType := "GET"
	API := constants.API_USER_LOGIN_INFO
	jsonReqData := ""
	isSuccess, token := Login(t)

	if isSuccess == false {
		t.Errorf("Unable to login")
		return
	}

	//webResponse
	_, jsonPathStruct := lhelper.ExecuteTestRequest(t, requestType, API, requestHandler, jsonReqData, token)

	if res, err := jsonpath.JsonPathLookup(jsonPathStruct, "$.issuccess"); err != nil {
		t.Errorf("Recieved issuccess %v", err.Error())
	} else if fmt.Sprintf("%v", res) != "true" {
		t.Errorf("Recieved issuccess %v", res)
	}
}
