package main

import (
	//Ginの実装
	"github.com/gin-gonic/gin"
	//"github.com/ssizawa/taskapp/server/src/repository"
)

func main() {
	//repository.Opendb()

	//Ginの変数
	router := gin.Default()

	//html/css/jsファイルのディレクトリをロード
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	//ポートを指定して実行
	router.Run(":3000")
}
