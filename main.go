package main

import (
	"encoding/json"
	"gland_test/initialize"
	"net/http"
)

type APIResponse struct {
	Data []FinancialEntry `json:"data"`
}

type FinancialEntry struct {
	Date       string  `json:"date"`
	StockID    string  `json:"stock_id"`
	Type       string  `json:"type"`
	Value      float64 `json:"value"`
	OriginName string  `json:"origin_name"`
}

func fetchDataFromExternalAPI() ([]FinancialEntry, error) {
	// 在这里调用外部API，这里使用一个示例URL
	apiURL := "https://api.finmindtrade.com/api/v4/data?dataset=TaiwanStockFinancialStatements&data_id=5225&start_date=2023-02-01&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRlIjoiMjAyNC0wMy0wOSAyMTo0MTowNyIsInVzZXJfaWQiOiJ3aWxseTAwNTUiLCJpcCI6IjEyMi4xMTcuMTgwLjE1OSJ9.HkU9BX9aUbyQ66zPPzgrkp0wz98g9HwymeUOYPUYPi4"

	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var apiResponse APIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	return apiResponse.Data, nil
}

func filterDataByType(data []FinancialEntry, targetType string) float64 {
	//var filteredData []FinancialEntry
	var eps float64

	for _, entry := range data {
		if entry.Type == targetType {
			eps += entry.Value
			//filteredData = append(filteredData, entry)
		}
	}

	return eps
}

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

}
