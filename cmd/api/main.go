package main

import (
	"fmt"
	"net/http"
	"time"

	"addnewer.com/godemo/pkg/app"
	"addnewer.com/godemo/pkg/config"
	"addnewer.com/godemo/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	file := "/home/zhe/workspace/go/src/godemo/config/config.yml"
	conf, err := config.New(file)
	if err != nil {
		panic("配置文件初始化错误: " + err.Error())
	}

	fmt.Printf("%#v", conf)

	appConf := app.NewConfig(conf)
	fmt.Printf("%#v", appConf)

	// mysql
	orm, err := database.NewOrmClient(conf)
	if err != nil {
		panic("数据库连接错误: " + err.Error())
	}

	user := make(map[string]interface{}, 10)
	orm.Table("users").Select("id", "realname", "email").Where("id = ?", 10).Take(user)

	fmt.Printf("%#v", user)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	srv := &http.Server{
		Addr:         ":8081",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic("Http启动错误: " + err.Error())
	}

}
