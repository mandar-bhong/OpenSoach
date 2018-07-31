package webcontent

import (
	"github.com/gin-gonic/gin"
)

func registerRouters(router *gin.RouterGroup) {

	router.Static("/web", "./web/hpft/")
	router.Static("/assets/", "./web/hpft/assets")
	router.Static("/shared/", "./web/hpft/shared")
	router.StaticFile("/", "./web/hpft/index.html")
}
