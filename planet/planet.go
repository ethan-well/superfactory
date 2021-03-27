package planet

import (
	"github.com/ItsWewin/superfactory/planet/httpserver"
	"github.com/gorilla/mux"
	"net/http"
)

type Planet struct {
	HttpServer *httpserver.HttpServer
}

func Init() *Planet {
	return &Planet{}
}

func (p *Planet) HttpServerInit(opt *http.Server, register func(mux *mux.Router)) {
	p.HttpServer = &httpserver.HttpServer{}

	p.HttpServer.Init(opt).Register(register).HttpStart()
}
