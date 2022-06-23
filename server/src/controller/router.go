package controller

import (
	//HTTPクライアントとサーバーの実装

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ssizawa/taskapp/server/src/repository"
)

func Router(router *gin.Engine) {

	var user_name string

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "login.html", gin.H{
			"judge": true,
		})
	})

	router.POST("/login", func(c *gin.Context) {

		user_name = c.PostForm("name")
		password := c.PostForm("password")

		if repository.VerifybyName(user_name, password) {
			c.Redirect(302, "/taskapp")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"judge": false,
			})
		}
	})

	router.GET("/taskapp", func(c *gin.Context) {

		var user_list []string
		user_list = repository.GetUserList()

		todo_task_list, doing_task_list, done_task_list := repository.GetTasks(user_name)

		c.HTML(200, "index.html", gin.H{
			"user_name":       user_name,
			"user_list":       user_list,
			"todo_task_list":  todo_task_list,
			"doing_task_list": doing_task_list,
			"done_task_list":  done_task_list,
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

		if repository.Verify(user_name, oldpass) {
			repository.ChangePass(user_name, newpass)
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

		c.Redirect(302, "/taskapp")
	})

	router.POST("/update_todo", func(c *gin.Context) {

		taskName := c.PostForm("taskName")
		taskDescription := c.PostForm("taskDescription")

		repository.UpdateTaskStatus(user_name, taskName, taskDescription, "todo")

		fmt.Println("\n[" + user_name + "]")
		fmt.Println("Status of \"" + taskName + "-" + taskDescription + "\" is changed into \"todo\"\n")
	})

	router.POST("/update_doing", func(c *gin.Context) {

		taskName := c.PostForm("taskName")
		taskDescription := c.PostForm("taskDescription")

		repository.UpdateTaskStatus(user_name, taskName, taskDescription, "doing")

		fmt.Println("\n[" + user_name + "]")
		fmt.Println("Status of \"" + taskName + "-" + taskDescription + "\" is changed into \"doing\"\n")
	})

	router.POST("/update_done", func(c *gin.Context) {

		taskName := c.PostForm("taskName")
		taskDescription := c.PostForm("taskDescription")

		repository.UpdateTaskStatus(user_name, taskName, taskDescription, "done")

		fmt.Println("\n[" + user_name + "]")
		fmt.Println("Status of \"" + taskName + "-" + taskDescription + "\" is changed into \"done\"\n")
	})

	router.POST("/delete_task", func(c *gin.Context) {

		taskName := c.PostForm("taskName")
		taskDescription := c.PostForm("taskDescription")

		fmt.Println(taskDescription)

		repository.DeleteTask(user_name, taskName, taskDescription)
	})
}
