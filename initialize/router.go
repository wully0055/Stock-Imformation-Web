package initialize

import (
	"Stock-Imformation-Web/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os/exec"
	"runtime"
)

func Routers() *gin.Engine {

	Router := gin.New()
	Router.Use(cors.Default())

	systemRouter := router.RouterGroupApp.System

	// 設置靜態文件目錄
	Router.StaticFile("/", "index.html")

	Router.LoadHTMLGlob("templates/*")

	Router.GET("templates/stockdata.html", func(c *gin.Context) {
		c.File("templates/stockdata.html")
	})
	Router.GET("templates/myfavorited.html", func(c *gin.Context) {
		c.File("templates/myfavorited.html")
	})

	PublicGroup := Router.Group("test")

	systemRouter.InitApiRouter(PublicGroup)

	// 在此之前打開瀏覽器
	openBrowser("http://localhost:8080")

	Router.Run(":8080")

	return Router

}

// openBrowser 根據給定的 URL 使用系統預設的瀏覽器打開網頁
func openBrowser(url string) error {
	var cmd string
	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "windows":
		cmd = "start"
	default:
		cmd = "xdg-open"
	}
	return exec.Command(cmd, url).Start()
}
