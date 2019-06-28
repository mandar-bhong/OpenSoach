package middleware

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"opensoach.com/core/logger"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		beforeReqestTime := time.Now()
		logger.Context().LogDebug("RequestLogger", logger.Instrumentation, fmt.Sprintf("-----------------%s-----------------", c.Request.RequestURI))

		switch c.Request.Method {
		case "POST":
			logger.Context().LogDebug("RequestLogger", logger.Instrumentation, fmt.Sprintf("Request Body: %s", readBody(rdr1))) // Print request body
			break
		case "GET":
			requestParams := c.Request.URL.Query()
			requestParamsData := fmt.Sprintf("Request Param: %+v", requestParams)
			logger.Context().LogDebug("RequestLogger", logger.Instrumentation, requestParamsData) // Print request body
		}

		logger.Context().LogDebug("RequestLogger", logger.Instrumentation, "---------------------------------------------------------")
		c.Request.Body = rdr2
		c.Next()

		afterReqestTime := time.Now()
		diff := afterReqestTime.Sub(beforeReqestTime)
		logger.Context().LogDebug("", logger.Performace, fmt.Sprintf("Time taken for execution '%s' request: '%s'", c.Request.RequestURI, diff.String()))

		logger.Context().LogDebug("RequestLogger", logger.Instrumentation, "---------------------------------------------------------")
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
