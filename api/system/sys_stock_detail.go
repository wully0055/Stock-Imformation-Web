package system

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gland_test/global"
	"gland_test/model/system"
	"log"
	"net/http"
	"strconv"
	"time"
)

type StockDetail struct{}

type StockImformation struct {
	Code          string `json:"Code"`
	Name          string `json:"Name"`
	PEratio       string `json:"PEratio"`
	DividendYield string `json:"DividendYield"`
	PBratio       string `json:"PBratio"`
}

type StockImformation2 struct {
	//Date                  string `json:"Date"`
	Code          string `json:"SecuritiesCompanyCode"`
	Name          string `json:"CompanyName"`
	PEratio       string `json:"PriceEarningRatio"`
	DividendYield string `json:"YieldRatio"`
	PBratio       string `json:"PriceBookRatio"`
	//DividendPerShare      string `json:"DividendPerShare"`
}

type StockCode struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	PEratio string `json:"peratio"`
}

type FavoritedStock struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

type MsgArrayItem struct {
	Tv    string `json:"tv"`
	Ps    string `json:"ps"`
	Pz    string `json:"pz"`
	Bp    string `json:"bp"`
	Fv    string `json:"fv"`
	Oa    string `json:"oa"`
	Ob    string `json:"ob"`
	A     string `json:"a"`
	B     string `json:"b"`
	C     string `json:"c"`
	D     string `json:"d"`
	Ch    string `json:"ch"`
	Ot    string `json:"ot"`
	Tlong string `json:"tlong"`
	F     string `json:"f"`
	Ip    string `json:"ip"`
	G     string `json:"g"`
	Mt    string `json:"mt"`
	Ov    string `json:"ov"`
	H     string `json:"h"`
	I     string `json:"i"`
	It    string `json:"it"`
	Oz    string `json:"oz"`
	L     string `json:"l"`
	N     string `json:"n"`
	O     string `json:"o"`
	P     string `json:"p"`
	Ex    string `json:"ex"`
	S     string `json:"s"`
	T     string `json:"t"`
	U     string `json:"u"`
	V     string `json:"v"`
	W     string `json:"w"`
	Nf    string `json:"nf"`
	Y     string `json:"y"`
	Z     string `json:"z"`
	Ts    string `json:"ts"`
}

type APIResponse struct {
	MsgArray []MsgArrayItem `json:"msgArray"`
}

type Epsdata struct {
	Data []FinancialEntry `json:"data"`
}

type FinancialEntry struct {
	Date       string  `json:"date"`
	StockID    string  `json:"stock_id"`
	Type       string  `json:"type"`
	Value      float64 `json:"value"`
	OriginName string  `json:"origin_name"`
}

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
	var stocks []StockImformation
	if err := json.NewDecoder(resp.Body).Decode(&stocks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse external API response"})
		return
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
	var stocks2 []StockImformation2
	if err := json.NewDecoder(resp2.Body).Decode(&stocks2); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse external API response"})
		return
	}

	merge_data := mergeStructs(stocks, stocks2)
	StockTableData(merge_data)
	// 將API的結果傳遞到前端
	c.JSON(http.StatusOK, merge_data)
}

// 合併上市櫃股票資料
func mergeStructs(data1 []StockImformation, data2 []StockImformation2) []StockImformation {
	var mergedData []StockImformation
	mergedData = append(mergedData, data1...)

	for _, entry := range data2 {

		newdata1 := StockImformation{
			Code:          entry.Code,
			Name:          entry.Name,
			PEratio:       entry.PEratio,
			DividendYield: entry.DividendYield,
			PBratio:       entry.PBratio,
		}
		mergedData = append(mergedData, newdata1)

	}
	return mergedData
}

// StockData 上市股票詳細資料
func (s *StockDetail) StockData(c *gin.Context) {

	var requestData StockCode
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 處理 requestData
	apiURL := "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_" + requestData.Code + ".tw"

	resp, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch external API"})
		return
	}
	defer resp.Body.Close()

	// 解析API的JSON響應
	var apiResponse APIResponse

	// 僅解碼 JSON 數據中的 msgArray 字段
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse external API response"})
		return
	}

	eps := StockEPS(requestData.Code)

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
		ValuesMap["EPS"] = eps
		ValuesMap["合理價"] = eps_str
	}

	c.JSON(http.StatusOK, ValuesMap)

	// 返回響應
}

// StockData2 上櫃股票詳細資料
func (s *StockDetail) StockData2(c *gin.Context) {

	var requestData StockCode
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 處理 requestData
	apiURL2 := "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=otc_" + requestData.Code + ".tw"

	resp, err := http.Get(apiURL2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch external API"})
		return
	}
	defer resp.Body.Close()

	// 解析API的JSON響應
	var apiResponse APIResponse

	// 僅解碼 JSON 數據中的 msgArray 字段
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse external API response"})
		return
	}

	eps := StockEPS(requestData.Code)

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
		ValuesMap["EPS"] = eps
		ValuesMap["合理價"] = eps_str
	}

	c.JSON(http.StatusOK, ValuesMap)

	// 返回響應
}

// StockEPS 取得近四次EPS總和
func StockEPS(value string) string {
	now := time.Now()
	oneYearAgo := now.AddDate(-1, -6, 0)

	dataset := "TaiwanStockFinancialStatements"
	data_id := value
	start_date := oneYearAgo.Format("2006-01-02")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRlIjoiMjAyNC0wMy0wOSAyMTo0MTowNyIsInVzZXJfaWQiOiJ3aWxseTAwNTUiLCJpcCI6IjEyMi4xMTcuMTgwLjE1OSJ9.HkU9BX9aUbyQ66zPPzgrkp0wz98g9HwymeUOYPUYPi4"

	apiURL := "https://api.finmindtrade.com/api/v4/data?dataset=" + dataset + "&data_id=" + data_id + "&start_date=" + start_date + "&token=" + token + ""
	response, err := http.Get(apiURL)
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	var epsResponse Epsdata
	err = json.NewDecoder(response.Body).Decode(&epsResponse)
	if err != nil {
		return ""
	}

	//return epsResponse.Data, nil
	var eps float64
	count := 0
	//for _, entry := range epsResponse.Data {
	//	if entry.Type == "EPS" {
	//		eps += entry.Value
	//		//filteredData = append(filteredData, entry)
	//	}
	//}
	//改為反向迴圈, 取最近 ４ 筆 EPS
	for i := len(epsResponse.Data) - 1; i >= 0; i-- {
		entry := epsResponse.Data[i]
		if entry.Type == "EPS" && count < 4 {
			eps += entry.Value
			count++
			//filteredData = append(filteredData, entry)
		}
	}
	eps_str := fmt.Sprintf("%.3f", eps)

	return eps_str
}

// StockTableData 新增股票代號和名稱到資料庫
func StockTableData(merge_data []StockImformation) {
	db := global.SKW_DB
	var count int64
	//檢查是否為空table
	db.Model(&system.SysStockTable{}).Count(&count)
	if count == 0 {

		// 將 API 回傳的資料新增到資料庫表格中
		for _, response := range merge_data {
			data := system.SysStockTable{
				StockID:   response.Code,
				StockName: response.Name,
			}
			result := db.Create(&data)
			if result.Error != nil {
				log.Println("Failed to insert data:", result.Error)
			} else {
				fmt.Println("Data inserted successfully:", data)
			}
		}
	} else {
		fmt.Println("股票資料已新增至資料庫") //之後可註解
	}
}

func (s *StockDetail) Check_Favorited(c *gin.Context) {
	var requestData FavoritedStock
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
