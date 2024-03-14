package initialize

import "gorm.io/gorm"

func GormMysql() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/stockweb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
