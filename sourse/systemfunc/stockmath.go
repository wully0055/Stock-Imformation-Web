package systemfunc

import (
	"Stock-Imformation-Web/model/common/response"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// StockEPS 取得近四次EPS總和
func StockEPS(value string) string {
	now := time.Now()
	oneYearAgo := now.AddDate(-1, -6, 0)

	dataset := "TaiwanStockFinancialStatements"
	data_id := value
	start_date := oneYearAgo.Format("2006-01-02")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRlIjoiMjAyNC0wMy0wOSAyMTo0MTowNyIsInVzZXJfaWQiOiJ3aWxseTAwNTUiLCJpcCI6IjEyMi4xMTcuMTgwLjE1OSJ9.HkU9BX9aUbyQ66zPPzgrkp0wz98g9HwymeUOYPUYPi4"

	apiURL := "https://api.finmindtrade.com/api/v4/data?dataset=" + dataset + "&data_id=" + data_id + "&start_date=" + start_date + "&token=" + token + ""
	respond, err := http.Get(apiURL)
	if err != nil {
		return ""
	}
	defer respond.Body.Close()

	var epsResponse response.Epsdata
	err = json.NewDecoder(respond.Body).Decode(&epsResponse)
	if err != nil {
		return ""
	}

	var eps float64
	count := 0

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
