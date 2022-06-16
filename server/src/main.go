package main

import (
	//Ginの実装
	"github.com/gin-gonic/gin"

	"github.com/ssizawa/taskapp/server/src/controller"
)

func main() {

	//Ginの変数
	router := gin.Default()

	//html/css/jsファイルのディレクトリをロード
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	//routerを渡す
	controller.Router(router)

	//ポートを指定して実行
	router.Run(":80")
}
