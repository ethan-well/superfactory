package powertrain

import (
	"github.com/ItsWewin/superfactory/logger"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Options struct {
	// 服务初始化时候优先执行的方法
	// 这里进行全局的 panic 处理
	RecoverFunc func()
	InitFunc func()
	DeferFunc func()
	RegisterRouter func(mux *mux.Router)
	HttpServerOptions *http.Server
}

type HttpServerOptions struct {
	Addr string
	ReadTimeout time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout time.Duration
	IdleTimeout      time.Duration
}

func defaultOptions() *Options {
	recoverFunc := func() {
		if err := recover(); err != nil {
			logger.Errorf("some err recovered: %s", err)
		}
	}

	initFunc := func() {
		logger.Infof("server init")
	}

	return &Options{
		RecoverFunc: recoverFunc,
		InitFunc: initFunc,
		DeferFunc: func(){},
		HttpServerOptions: &http.Server{
			Addr:              ":8080",
			ReadTimeout:       30 * time.Second,
			ReadHeaderTimeout: 30 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       30 * time.Second,
		},
	}
}
