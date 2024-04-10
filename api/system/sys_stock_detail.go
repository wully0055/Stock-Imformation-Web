package system

import (
	"Stock-Imformation-Web/model/common/request"
	"Stock-Imformation-Web/model/common/response"
	"Stock-Imformation-Web/sourse/systemfunc"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type StockDetail struct{}

func (s *StockDetail) StockImformation(c *gin.Context) {
	//上市資料
	apiURL := "https://openapi.twse.com.tw/v1/exchangeReport/BWIBBU_ALL"
	resp, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch external API"})
		return
	}
	defer resp.Body.Close()

	// 解析API的JSON響應
	var stocks []request.StockImformation
	if err := json.NewDecoder(resp.Body).Decode(&stocks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse external API response"})
		return
	}
	// 在每個 StockImformation 中添加 type 字段並設置為 "上市"
	for i := range stocks {
		stocks[i].Type = "上市"
	}
	//上櫃資料
	apiURL2 := "https://www.tpex.org.tw/openapi/v1/tpex_mainboard_peratio_analysis"
	resp2, err := http.Get(apiURL2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch external API"})
		return
	}
	defer resp2.Body.Close()

	// 解析API的JSON響應
	var stocks2 []request.StockImformation2
	if err := json.NewDecoder(resp2.Body).Decode(&stocks2); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse external API response"})
		return
	}
	// 在每個 StockImformation2 中添加 type 字段並設置為 "上櫃"
	for i := range stocks2 {
		stocks2[i].Type = "上櫃"
	}

	merge_data := systemfunc.MergeStructs(stocks, stocks2)
	systemfunc.StockTableData(merge_data)
	// 將API的結果傳遞到前端
	c.JSON(http.StatusOK, merge_data)
}

// StockData 上市櫃股票詳細資料
func (s *StockDetail) StockData(c *gin.Context) {

	var requestData request.StockCode
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var apiURL string

	if requestData.Type == "上市" {
		// 處理 requestData
		apiURL = "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_" + requestData.Code + ".tw"
	} else {
		// 處理 requestData
		apiURL = "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=otc_" + requestData.Code + ".tw"
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch external API"})
		return
	}
	defer resp.Body.Close()

	// 解析API的JSON響應
	var apiResponse response.StockMsg

	// 僅解碼 JSON 數據中的 msgArray 字段
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse external API response"})
		return
	}

	eps := systemfunc.StockEPS(requestData.Code)

	floateps, err := strconv.ParseFloat(eps, 64)
	if err != nil {
		fmt.Println("eps轉換float失敗:", err)
		return
	}
	//若沒有本益比資料，本益比設為 0
	var stockPEratio string
	stockPEratio = requestData.PEratio
	if stockPEratio == "" || stockPEratio == "N/A" {
		stockPEratio = "0"
	}

	floatPEratio, err := strconv.ParseFloat(stockPEratio, 64)
	if err != nil {
		fmt.Println("本益比轉換float失敗:", err)
		return
	}
	price := floateps * floatPEratio
	eps_str := fmt.Sprintf("%.3f", price)

	ValuesMap := make(map[string]string)
	for _, item := range apiResponse.MsgArray {
		ValuesMap["股票代號"] = item.C
		ValuesMap["公司簡稱"] = item.N
		ValuesMap["成交價"] = item.Z
		ValuesMap["成交量"] = item.Tv
		ValuesMap["累積成交量"] = item.V
		ValuesMap["開盤價"] = item.O
		ValuesMap["最高價"] = item.H
		ValuesMap["最低價"] = item.L
		ValuesMap["昨收價"] = item.Y
		ValuesMap["本益比"] = requestData.PEratio
		ValuesMap["EPS"] = eps
		ValuesMap["合理價"] = eps_str
		ValuesMap["類型"] = "上市"
	}

	c.JSON(http.StatusOK, ValuesMap)

	// 返回響應
}
