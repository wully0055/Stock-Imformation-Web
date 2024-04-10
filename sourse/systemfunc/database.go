package systemfunc

import (
	"Stock-Imformation-Web/global"
	"Stock-Imformation-Web/model/common/request"
	"Stock-Imformation-Web/model/system"
	"fmt"
	"log"
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
		fmt.Println("股票資料已新增至資料庫") //之後可註解
	}
}
