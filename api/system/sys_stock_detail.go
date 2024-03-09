package system

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StockDetail struct{}

type StockImformation struct {
	Code          string `json:"Code"`
	Name          string `json:"Name"`
	PEratio       string `json:"PEratio"`
	DividendYield string `json:"DividendYield"`
	PBratio       string `json:"PBratio"`
}

type StockCode struct {
	Code string `json:"code"`
	Name string `json:"name"`
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

func (s *StockDetail) StockImformation(c *gin.Context) {
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

	// 將API的結果傳遞到前端
	c.JSON(http.StatusOK, stocks)
}

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
	}

	c.JSON(http.StatusOK, ValuesMap)

	// 返回響應
	//c.JSON(http.StatusOK, gin.H{"message": "POST request received", "data": requestData})

	//fmt.Printf("%+v\n", tv)

}
