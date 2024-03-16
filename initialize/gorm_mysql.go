package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {

	dsn := "root:@tcp(127.0.0.1:3306)/stockweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		return nil
	} else {

		return db
	}

	//defer db.Close()
}
