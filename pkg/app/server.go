package app

import (
	"fmt"
	"net/http"
	"time"
)

type AppServer struct {
	conf   *Config
	router *http.Handler
	srv    *http.Server
}

func NewAppServer(c *Config, r *http.Handler) *AppServer {
	app := &AppServer{
		conf:   c,
		router: r,
	}

	app.initHttpServer()

	return app
}

func (a *AppServer) Run() {
	if err := a.srv.ListenAndServe(); err != nil {
		panic("Http启动错误: " + err.Error())
	}
}

func (a *AppServer) SetHttpServer(s *http.Server) {
	a.srv = s
}

func (a *AppServer) initHttpServer() {
	a.srv = &http.Server{
		Addr:         ":8081",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      *a.router,
	}

	fmt.Println("app name: " + a.conf.Name)
}
