package webcontent

import (
	"github.com/gin-gonic/gin"
)

func registerRouters(router *gin.Engine) {
	router.Static("/web/", "./web/")
	router.StaticFile("/", "./web/index.html")
}
