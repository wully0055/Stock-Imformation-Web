package main

import "gland_test/initialize"

type Stock struct {
	Code          string `json:"Code"`
	Name          string `json:"Name"`
	PEratio       string `json:"PEratio"`
	DividendYield string `json:"DividendYield"`
	PBratio       string `json:"PBratio"`
}

func main() {

	//// 创建一个Gin引擎
	//r := gin.Default()
	//
	//// Use CORS middleware
	//r.Use(cors.Default())
	//
	//// 定义一个处理程序，该处理程序调用外部API并将结果传递到前端
	//r.GET("/getExternalAPI", func(c *gin.Context) {
	//	// 使用http.Client发起对外部API的GET请求
	//	apiURL := "https://openapi.twse.com.tw/v1/exchangeReport/BWIBBU_ALL"
	//	resp, err := http.Get(apiURL)
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch external API"})
	//		return
	//	}
	//	defer resp.Body.Close()
	//
	//	// 解析API的JSON响应
	//	var stocks []Stock
	//	if err := json.NewDecoder(resp.Body).Decode(&stocks); err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse external API response"})
	//		return
	//	}
	//
	//	// 将API的结果传递到前端
	//	c.JSON(http.StatusOK, stocks)
	//})
	//
	//// 启动Gin服务器
	//r.Run(":8080")

	initialize.Routers()

}
