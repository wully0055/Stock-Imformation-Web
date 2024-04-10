package system

import (
	"Stock-Imformation-Web/global"
	"Stock-Imformation-Web/model/common/request"
	"Stock-Imformation-Web/model/system"
	"Stock-Imformation-Web/sourse/systemfunc"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type StockFavorite struct{}

func (s *StockFavorite) Check_Favorited(c *gin.Context) {
	var requestData request.FavoritedStock
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := global.SKW_DB
	var stockid system.SysMyFavourite
	db.Where("stock_id = ?", requestData.Code).First(&stockid)
	if stockid.ID != 0 {
		if requestData.Type == 0 {
			c.JSON(http.StatusOK, 1)
		} else {
			db.Unscoped().Delete(&stockid) //整個從table砍掉
			//db.Delete(&stockid)
			c.JSON(http.StatusOK, 0)
		}
		// 找到了，執行刪除操作

	} else {
		if requestData.Type == 0 {
			c.JSON(http.StatusOK, 0)
		} else {
			// 未找到，新增至資料表
			data := system.SysMyFavourite{
				StockID:   requestData.Code,
				StockName: requestData.Name,
				PEratio:   requestData.PEratio,
			}
			result := db.Create(&data)
			if result.Error != nil {
				log.Println("Failed to insert data:", result.Error)
			} else {
				fmt.Println("Data inserted successfully:", data)
			}
			c.JSON(http.StatusOK, 1)
		}

	}
}

// MyFavorited 用戶收藏的股票
func (s *StockFavorite) MyFavorited(c *gin.Context) {
	db := global.SKW_DB
	var data []system.SysMyFavourite
	db.Find(&data)
	var result []map[string]string

	for _, favorite := range data {
		var stock system.SysStockTable
		if err := db.Where("stock_id = ?", favorite.StockID).First(&stock).Error; err != nil {
			// 處理錯誤
			continue
		}
		stockData, err := systemfunc.FetchStockData(stock)
		if err != nil {
			// 處理錯誤
			continue
		}
		// 將兩個結構體的欄位合併成一個 map
		item := map[string]string{
			"code":           stock.StockID,
			"name":           stock.StockName,
			"peratio":        favorite.PEratio,
			"type":           stock.Type,
			"yesterdayprice": stockData["昨收價"],
			"todayprice":     stockData["成交價"],
		}
		result = append(result, item)
	}

	c.JSON(http.StatusOK, result)
}
