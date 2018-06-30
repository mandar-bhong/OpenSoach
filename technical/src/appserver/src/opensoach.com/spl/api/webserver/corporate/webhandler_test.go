package corporate

import (
	"fmt"
	"os"
	"testing"

	"github.com/oliveagle/jsonpath"
	"opensoach.com/spl/api/constants"
	lhelper "opensoach.com/spl/api/helper"
	login "opensoach.com/spl/api/webserver/login"
)

const TEST_USER_LOGIN_VALID = `{"username":"admin@servicepoint.live","password":"admin","prodcode":"SPL_HKT"}`
const TEST_CORPORATE_UPDATE_VALID = `{"corpid":5,"corpname":"Corporate 5","corpmobileno":"435435","corpemailid":"corp5@gmail.com",	"corplandlineno":"43243242" }`
const TEST_CORPORATE_ADD_VALID = `{"corpname":"Corporate 5","corpmobileno":"435435","corpemailid":"corp5@email.com","corplandlineno":"43243242" }`
const TEST_CORPORATE_LIST_OSU_VALID = `{"page":1,"limit":10,"orderby":"corpid","orderdirection":"asc"}`

func TestMain(m *testing.M) {

	lhelper.PrepareTestSetup()

	os.Exit(m.Run())
}

func Login(t *testing.T) (isSuccess bool, token string) {

	requestType := "POST"
	API := constants.API_USER_LOGIN
	jsonReqData := TEST_USER_LOGIN_VALID

	webResponse, jsonPathStruct := lhelper.ExecuteTestRequest(t, requestType, API, login.LoginRequestHandler, jsonReqData, token)

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

func Test_CorporateInsert(t *testing.T) {

	isSuccess, token := Login(t)

	if isSuccess == false {
		t.Errorf("Unable to login")
		return
	}

	requestType := "POST"
	API := constants.API_CORPORATE_OSU_ADD
	jsonReqData := TEST_CORPORATE_ADD_VALID

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

func Test_CorporateUpdate(t *testing.T) {

	isSuccess, token := Login(t)

	if isSuccess == false {
		t.Errorf("Unable to login")
		return
	}

	requestType := "POST"
	API := constants.API_CORPORATE_OSU_UPDATE
	jsonReqData := TEST_CORPORATE_UPDATE_VALID

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

func Test_CorporateGetShortList(t *testing.T) {

	isSuccess, token := Login(t)

	if isSuccess == false {
		t.Errorf("Unable to login")
		return
	}

	requestType := "GET"
	API := constants.API_CORPORATE_OSU_LIST_SHORT
	jsonReqData := ""

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

func Test_CorporateGetList(t *testing.T) {

	isSuccess, token := Login(t)

	if isSuccess == false {
		t.Errorf("Unable to login")
		return
	}

	requestType := "GET"
	API := constants.API_CORPORATE_OSU_LIST
	jsonReqData := TEST_CORPORATE_LIST_OSU_VALID

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
