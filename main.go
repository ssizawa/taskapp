package main

import (
	//HTTPクライアントとサーバーの実装
	"net/http"

	//Ginの実装
	"github.com/gin-gonic/gin"
)

func main() {

	//Ginの変数
	router := gin.Default()

	//html/css/jsファイルのディレクトリをロード
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	//GET
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//ポートを指定して実行
	router.Run(":3000")
}
