package middleware

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		switch c.Request.Method {
		case "POST":
			buf, _ := ioutil.ReadAll(c.Request.Body)
			rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
			rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

			fmt.Println("---------------------------------------------------------")
			fmt.Printf("Request Body: %s \n", readBody(rdr1)) // Print request body
			fmt.Println("---------------------------------------------------------")
			c.Request.Body = rdr2
			c.Next()
			break
		case "GET":

		}
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
