package main

import (
	"gland_test/initialize"
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

	initialize.Routers()
	initialize.Gorm()

}
