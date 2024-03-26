package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	dsn := "root:910315990055@tcp(127.0.0.1:3306)/stockweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 在無法連接到資料庫時創建名為 "stockweb" 的資料庫
		dsnWithoutDB := "root:910315990055@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
		createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS stockweb")
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dsnWithoutDB,
		}), &gorm.Config{})
		if err != nil {
			panic("無法創建資料庫: " + err.Error())
		}
		db.Exec(createDB)

		// 再次嘗試連接到 stockweb 資料庫
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("無法連接到資料庫: " + err.Error())
		}
	}
	return db
} //defer db.Close()
