package main

import (
	"Stock-Imformation-Web/global"
	"Stock-Imformation-Web/initialize"
)

func main() {

	global.SKW_DB = initialize.Gorm()

	if global.SKW_DB != nil {
		initialize.RegisterTables()

		// 程式结束前關閉資料庫連接
		db, _ := global.SKW_DB.DB()
		defer db.Close()
	}
	initialize.Routers()

}
