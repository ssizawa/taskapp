package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//Ginの変数
	engine := gin.Default()

	//htmlファイルのディレクトリをロード
	engine.LoadHTMLGlob("templates/*")

	//GET
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	//ポートを指定して実行
	engine.Run(":3000")

}
