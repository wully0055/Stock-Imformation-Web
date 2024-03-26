package main

import (
	"Stock-Imformation-Web/global"
	"Stock-Imformation-Web/initialize"
)

func main() {

	// 創建一個Gin引擎
	//r := gin.Default()

	// Use CORS middleware
	//r.Use(cors.Default())

	//r.GET("/getExternalAPI", func(c *gin.Context) {
	//
	//	// 调用外部API获取数据
	//	apiData, err := fetchDataFromExternalAPI()
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//		return
	//	}
	//	epsData := filterDataByType(apiData, "EPS")
	//
	//	c.JSON(http.StatusOK, epsData)
	//
	//	// 將API的結果傳遞到前端
	//})

	// 啟動Gin服務器
	//r.Run(":8080")

	//
	global.SKW_DB = initialize.Gorm()

	if global.SKW_DB != nil {
		initialize.RegisterTables()

		// 程式结束前關閉資料庫連接
		db, _ := global.SKW_DB.DB()
		defer db.Close()
	}
	initialize.Routers()

}
