package router

import (
	"addnewer.com/godemo/internel/controller"
	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/info", controller.UserInfo)
			user.GET("/list", controller.UserList)
		}
	}
}
