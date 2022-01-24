package controller

import (
	"net/http"

	"addnewer.com/godemo/internel/controller/dto"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	var userParams dto.UserParams
	if err := c.ShouldBind(&userParams); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "",
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": userParams,
	})
}

func UserList(c *gin.Context) {
	var userParams dto.UserParams
	if err := c.ShouldBind(&userParams); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "",
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": "Hello world!",
	})
}
