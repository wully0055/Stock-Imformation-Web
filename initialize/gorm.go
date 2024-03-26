package initialize

import (
	"Stock-Imformation-Web/global"
	"Stock-Imformation-Web/model/system"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {

	return GormMysql()
}

func RegisterTables() {
	db := global.SKW_DB
	err := db.AutoMigrate(
		system.SysStockTable{},
		system.SysMyFavourite{},
	)
	if err != nil {
		os.Exit(0)
	}
}
