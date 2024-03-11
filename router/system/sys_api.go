package system

import (
	"github.com/gin-gonic/gin"
	"gland_test/api"
)

type DetailRouter struct{}

func (e *DetailRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("/api")

	apiRouterApi := api.ApiGroupApp.SystemApiGroup.StockDetail
	{
		apiRouter.GET("stockimformation", apiRouterApi.StockImformation)
		apiRouter.POST("stockdata", apiRouterApi.StockData)   //上市個股日本益比、殖利率及股價淨值比
		apiRouter.POST("stockdata2", apiRouterApi.StockData2) //上櫃個股日本益比、殖利率及股價淨值比

	}
}
