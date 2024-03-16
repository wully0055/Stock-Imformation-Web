package main

import (
	"gland_test/global"
	"gland_test/initialize"
)

func main() {

	// 創建一個Gin引擎
	//r := gin.Default()

	// Use CORS middleware
	//r.Use(cors.Default())

	//r.GET("/getExternalAPI", func(c *gin.Context) {
	//
	//	// 调用外部API获取数据
	//	apiData, err := fetchDataFromExternalAPI()
	//	if err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//		return
	//	}
	//	epsData := filterDataByType(apiData, "EPS")
	//
	//	c.JSON(http.StatusOK, epsData)
	//
	//	// 將API的結果傳遞到前端
	//})

	// 啟動Gin服務器
	//r.Run(":8080")

	//
	global.SKW_DB = initialize.Gorm()

	if global.SKW_DB != nil {
		initialize.RegisterTables()
		// 程式结束前關閉資料庫連接
		db, _ := global.SKW_DB.DB()
		defer db.Close()
	}
	initialize.Routers()
	//dsn := "root:@tcp(127.0.0.1:3306)/stockweb?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal(err)
	//	//return nil
	//}
	//// 創建名為 stockweb 的資料庫
	//err = db.Exec("CREATE DATABASE IF NOT EXISTS stockweb").Error
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 重新連接到 stockweb 資料庫
	//dsn = "root:@tcp(127.0.0.1:3306)/stockweb?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Connected to stockweb database successfully!")

}
