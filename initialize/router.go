package initialize

import (
	"Stock-Imformation-Web/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {

	Router := gin.New()
	Router.Use(cors.Default())

	systemRouter := router.RouterGroupApp.System

	PublicGroup := Router.Group("test")

	systemRouter.InitApiRouter(PublicGroup)

	Router.Run(":8080")

	return Router

}
