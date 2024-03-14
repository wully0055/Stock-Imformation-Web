package initialize

import (
	"gland_test/global"
	"gland_test/model/system"
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
	)
	if err != nil {
		os.Exit(0)
	}
}
