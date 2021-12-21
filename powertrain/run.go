package powertrain

import (
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/basicmatter"
	"github.com/ItsWewin/superfactory/basicmatter/config"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/planet"
)

type Option func(o *Options)

func Run(configObj basicmatter.Master, options ...Option) {
	err := config.UnmarshalConf(configObj)
	if err != nil {
		panic(err)
	}

	opts := defaultOptions()
	for _, o := range options {
		o(opts)
	}
	defer opts.DeferFunc()
	opts.RecoverFunc()
	opts.InitFunc()

	err = startHttpServer(configObj, opts)
	if err != nil {
		panic(err)
	}
}

func startHttpServer(configObj interface{}, opts *Options) aerror.Error {
	if !opts.InitHttpServer {
		logger.Infof("no need init http server")
		return nil
	}

	logger.Infof("init http server")

	conf, ok := configObj.(basicmatter.ConfigMater)
	if !ok {
		logger.Infof("无法获取 http addr")
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "conf error")
	}

	// http 初始化，路由注册，监听端口
	opts.HttpServerOptions.Addr = conf.HttpAddr()
	planet.Init().HttpServerInit(opts.HttpServerOptions, opts.RegisterRouter)
	return nil
}
