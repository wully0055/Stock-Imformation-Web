package system

import (
	"Stock-Imformation-Web/api"
	"github.com/gin-gonic/gin"
)

type DetailRouter struct{}

func (e *DetailRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("/api")

	apiStockDetailApi := api.ApiGroupApp.SystemApiGroup.StockDetail
	{
		apiRouter.GET("stockimformation", apiStockDetailApi.StockImformation) //首頁資訊
		apiRouter.POST("stockdata", apiStockDetailApi.StockData)              //上市個股日本益比、殖利率及股價淨值比
	}
	apiStockFavoriteApi := api.ApiGroupApp.SystemApiGroup.StockFavorite
	{
		apiRouter.POST("check_favorited", apiStockFavoriteApi.Check_Favorited) //(檢查)收藏股票
		apiRouter.GET("myfavorited", apiStockFavoriteApi.MyFavorited)          //用戶收藏的股票
	}
}
