package webcontent

import (
	"github.com/gin-gonic/gin"
)

func registerRouters(router *gin.RouterGroup) {

	router.Static("/web", "./web/vst/")
	router.Static("/assets/", "./web/vst/assets")
	router.Static("/shared/", "./web/vst/shared")
	router.StaticFile("/", "./web/vst/index.html")
}
