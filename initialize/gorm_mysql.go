package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GormMysql() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 在無法連接到資料庫時創建名為 "stockweb" 的資料庫
		dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port)
		createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbname)
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
