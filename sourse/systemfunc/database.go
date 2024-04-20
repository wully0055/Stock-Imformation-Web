package systemfunc

import (
	"Stock-Imformation-Web/global"
	"Stock-Imformation-Web/model/common/request"
	"Stock-Imformation-Web/model/system"
	"fmt"
	"log"
	"time"
)

// StockTableData 新增全部股票代號和名稱到資料庫
func StockTableData(merge_data []request.StockImformation) {
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
				Type:      response.Type,
			}
			result := db.Create(&data)
			if result.Error != nil {
				log.Println("Failed to insert data:", result.Error)
			} else {
				fmt.Println("Data inserted successfully:", data)
			}
		}
	} else {
		//message := "股票資料已新增至資料庫"
		message := "股票資料首頁已更新"
		currentTime := time.Now()
		currentTimeString := currentTime.Format("2006-01-02 15:04:05")
		messageWithTime := fmt.Sprintf("%s %s", message, currentTimeString)
		fmt.Println(messageWithTime)
	}
}
