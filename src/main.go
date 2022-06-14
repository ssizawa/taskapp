package main

import (
	//HTTPクライアントとサーバーの実装
	"fmt"
	"net/http"

	//Ginの実装
	"github.com/gin-gonic/gin"

	"github.com/ssizawa/taskapp/src/repository"
)

func main() {
	repository.Create_table()

	//Ginの変数
	router := gin.Default()

	//html/css/jsファイルのディレクトリをロード
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	//GET
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//POSTのデータ
	router.POST("/TaskApp", func(c *gin.Context) {

		name := c.PostForm("name")
		password := c.PostForm("password")

		fmt.Print(name, password)
	})

	//ポートを指定して実行
	router.Run(":3000")
}
