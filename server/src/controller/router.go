package controller

import (
	//HTTPクライアントとサーバーの実装
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ssizawa/taskapp/server/src/repository"
)

var name string

func Router(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"judge": true,
		})
	})

	router.POST("/login", func(c *gin.Context) {
		name = c.PostForm("name")
		password := c.PostForm("password")

		if repository.Verify(name, password) {
			c.Redirect(http.StatusMovedPermanently, "/taskapp")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"judge": false,
			})
		}
	})

	router.GET("/taskapp", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": name,
		})
	})

	router.GET("taskapp/settings", func(c *gin.Context) {
		c.HTML(http.StatusOK, "settings.html", gin.H{})
	})

	router.GET("/changepass", func(c *gin.Context) {
		c.HTML(http.StatusOK, "changepass.html", gin.H{
			"judge": true,
		})
	})

	router.POST("/changepass", func(c *gin.Context) {

		oldpass := c.PostForm("oldpass")
		newpass := c.PostForm("newpass")

		oldpass_judge := false
		changepass_successful := false
		if repository.Verify(name, oldpass) {

			repository.ChangePass(name, newpass)

			changepass_successful = true
		} else {
			oldpass_judge = true
		}

		c.HTML(http.StatusOK, "changepass.html", gin.H{
			"oldpass_judge":         oldpass_judge,
			"changepass_successful": changepass_successful,
		})
	})

}
