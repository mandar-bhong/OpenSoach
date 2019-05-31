package webcontent

import (
	"github.com/gin-gonic/gin"
)

func registerRouters(router *gin.RouterGroup) {

	router.Static("/web", "./web/hkt/")
	router.Static("/assets/", "./web/hkt/assets")
	router.Static("/shared/", "./web/hkt/shared")
	router.StaticFile("/", "./web/hkt/index.html")
}
