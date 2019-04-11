package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"opensoach.com/core"
	gmodels "opensoach.com/models"
	repo "opensoach.com/spl/api/repository"
)

const TEST_USER_LOGIN_VALID = `{"username":"admin@servicepoint.live","password":"admin","prodcode":"SPL_HKT"}`

func PrepareTestSetup() {

	setting := gmodels.ConfigSettings{}
	setting.DBConfig = &gmodels.ConfigDB{}
	setting.DBConfig.ConnectionString = "root:welcome@tcp(localhost:3306)/spl_master?parseTime=true"

	ctx := core.Context{}

	repo.Init(&setting, &ctx)
	repo.Instance().Context.Master.DBConn = setting.DBConfig.ConnectionString
	repo.Instance().Context.Master.Cache.CacheAddress = `{"address":"localhost", "port":6379,"password":"","db":0}`

}

func ExecuteTestRequest(t *testing.T,
	reqMethod string,
	API string,
	handler RequestHandler,
	reqJSONData string,
	token string) (webResponse string, jsonPathStruct interface{}) {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	var req *http.Request
	var reqErr error

	if reqMethod == "POST" {
		router.POST(API, func(c *gin.Context) { CommonWebRequestHandler(c, handler) })
		req, reqErr = http.NewRequest(reqMethod, API, bytes.NewBuffer([]byte(reqJSONData)))
	} else {
		router.GET(API, func(c *gin.Context) { CommonWebRequestHandler(c, handler) })
		req, reqErr = http.NewRequest(reqMethod, API+"?params="+reqJSONData, bytes.NewBuffer([]byte(reqJSONData)))
	}

	if reqErr != nil {
		t.Fatalf("Unable to create request. Error : %s ", reqErr.Error())
		return "", nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	webResponse = resp.Body.String()
	json.Unmarshal([]byte(webResponse), &jsonPathStruct)

	return webResponse, jsonPathStruct
}
