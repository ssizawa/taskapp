package controller

import (
	//HTTPクライアントとサーバーの実装
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ssizawa/taskapp/server/src/repository"
)

func Router(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"judge": true,
		})
	})

	router.POST("/login", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")

		repository.Opendb()

		if repository.VerifybyName(name, password) {
			c.Redirect(http.StatusMovedPermanently, "/taskapp")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"judge": false,
			})
		}
	})

	router.GET("/taskapp", func(c *gin.Context) {

		var user_list []string
		user_list = repository.GetUserList()

		c.HTML(http.StatusOK, "index.html", gin.H{
			"user_list": user_list,
		})
	})
}
