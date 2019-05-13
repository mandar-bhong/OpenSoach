package webcontent

import (
	"github.com/gin-gonic/gin"
)

func registerRouters(router *gin.RouterGroup) {

	router.Static("/web", "./web/spl/")
	router.Static("/assets/", "./web/spl/assets")
	router.Static("/shared/", "./web/spl/shared")
	router.StaticFile("/", "./web/spl/index.html")
}
