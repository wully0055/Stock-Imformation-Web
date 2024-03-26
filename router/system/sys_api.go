package system

import (
	"Stock-Imformation-Web/api"
	"github.com/gin-gonic/gin"
)

type DetailRouter struct{}

func (e *DetailRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("/api")

	apiRouterApi := api.ApiGroupApp.SystemApiGroup.StockDetail
	{
		apiRouter.GET("stockimformation", apiRouterApi.StockImformation)
		apiRouter.POST("stockdata", apiRouterApi.StockData) //上市個股日本益比、殖利率及股價淨值比
		//apiRouter.POST("stockdata2", apiRouterApi.StockData2)           //上櫃個股日本益比、殖利率及股價淨值比
		apiRouter.POST("check_favorited", apiRouterApi.Check_Favorited) //(檢查)收藏股票
		apiRouter.GET("myfavorited", apiRouterApi.MyFavorited)          //用戶收藏的股票
	}
}
