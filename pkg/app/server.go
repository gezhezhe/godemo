package app

import (
	"net/http"
	"time"

	"addnewer.com/godemo/internel/router"
	"addnewer.com/godemo/pkg/config"
)

type AppServer struct {
	conf *Config
	srv  *http.Server
}

func NewAppServer(file string) *AppServer {
	conf, err := config.New(file)
	if err != nil {
		panic("配置文件初始化错误: " + err.Error())
	}

	appConf := NewConfig(conf)

	r := router.NewGinEngine()
	router.RegisterApiRouter(r)

	srv := &http.Server{
		Addr:         ":8081",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      r,
	}

	app := &AppServer{
		conf: appConf,
		srv:  srv,
	}

	return app
}

func (a *AppServer) Run() {
	if err := a.srv.ListenAndServe(); err != nil {
		panic("Http启动错误: " + err.Error())
	}
}

// func (a *AppServer) SetHttpServer(s *http.Server) {
// 	a.srv = s
// }
