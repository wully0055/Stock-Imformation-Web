package main

import (
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gland_test/initialize"
	"net/http"
)

func main() {

	// 創建一個Gin引擎
	r := gin.Default()

	// Use CORS middleware
	r.Use(cors.Default())

	// 定義一個處理程序，該處理程序調用外部API並將結果傳遞到前端
	r.GET("/getExternalAPI", func(c *gin.Context) {
		//dataset := c.Param("dataset")
		// 使用http.Client發起對外部API的GET請求
		apiURL := "https://api.finmindtrade.com/api/v4/data?dataset=TaiwanStockFinancialStatements&data_id=5225&start_date=2023-02-01&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRlIjoiMjAyNC0wMy0wOSAyMTo0MTowNyIsInVzZXJfaWQiOiJ3aWxseTAwNTUiLCJpcCI6IjEyMi4xMTcuMTgwLjE1OSJ9.HkU9BX9aUbyQ66zPPzgrkp0wz98g9HwymeUOYPUYPi4"
		resp, err := http.Get(apiURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch external API"})
			return
		}
		defer resp.Body.Close()

		// 解析API的JSON響應
		var stocks interface{}
		if err := json.NewDecoder(resp.Body).Decode(&stocks); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse external API response"})
			return
		}

		// 將API的結果傳遞到前端
		c.JSON(http.StatusOK, stocks)
	})

	// 啟動Gin服務器
	r.Run(":8080")

	initialize.Routers()

}
