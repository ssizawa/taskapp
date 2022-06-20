package controller

import (
	//HTTPクライアントとサーバーの実装
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ssizawa/taskapp/server/src/repository"
)

func Router(router *gin.Engine) {

	var name string

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "login.html", gin.H{
			"judge": true,
		})
	})

	router.POST("/login", func(c *gin.Context) {

		name = c.PostForm("name")
		password := c.PostForm("password")

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
			"user_name": name,
			"user_list": user_list,
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

	router.POST("/create_a_task", func(c *gin.Context) {

		task_name := c.PostForm("task_name")
		description := c.PostForm("task_description")
		pic := c.PostForm("task_pic")
		deadline := c.PostForm("task_deadline")

		repository.CreateTask(task_name, description, pic, deadline)

		c.Redirect(http.StatusMovedPermanently, "/taskapp")
	})
}
