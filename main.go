package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//Ginの変数
	router := gin.Default()

	//htmlファイルのディレクトリをロード
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	//GET
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//ポートを指定して実行
	router.Run(":3000")

}
