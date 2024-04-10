package systemfunc

import (
	"Stock-Imformation-Web/model/common/request"
	"Stock-Imformation-Web/model/common/response"
	"Stock-Imformation-Web/model/system"
	"encoding/json"
	"fmt"
	"net/http"
)

// 合併上市櫃股票資料
func MergeStructs(data1 []request.StockImformation, data2 []request.StockImformation2) []request.StockImformation {
	var mergedData []request.StockImformation
	mergedData = append(mergedData, data1...)

	for _, entry := range data2 {

		newdata1 := request.StockImformation{
			Code:          entry.Code,
			Name:          entry.Name,
			PEratio:       entry.PEratio,
			DividendYield: entry.DividendYield,
			PBratio:       entry.PBratio,
			Type:          entry.Type,
		}
		mergedData = append(mergedData, newdata1)

	}
	return mergedData
}

// 取得個股詳細資料
func FetchStockData(requestData system.SysStockTable) (map[string]string, error) {
	var apiURL string

	if requestData.Type == "上市" {
		apiURL = "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_" + requestData.StockID + ".tw"
	} else {
		apiURL = "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=otc_" + requestData.StockID + ".tw"
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch external API: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse response.StockMsg
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse external API response: %v", err)
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
		ValuesMap["類型"] = "上市"
	}

	return ValuesMap, nil
}
