package httpserver

import (
	"github.com/gorilla/mux"
	"net/http"
)

type HttpServer struct {
	s http.Server
}

func (s *HttpServer) Init(opt *http.Server) *HttpServer {
	planet := &HttpServer{}
	planet.s = http.Server{
		Addr:              opt.Addr,
		TLSConfig:         opt.TLSConfig,
		ReadTimeout:       opt.ReadTimeout,
		ReadHeaderTimeout: opt.ReadHeaderTimeout,
		WriteTimeout:      opt.WriteTimeout,
		IdleTimeout:       opt.IdleTimeout,
		MaxHeaderBytes:    opt.MaxHeaderBytes,
		TLSNextProto:      opt.TLSNextProto,
		ConnState:         opt.ConnState,
		ErrorLog:          opt.ErrorLog,
		BaseContext:       opt.BaseContext,
		ConnContext:       opt.ConnContext,
	}

	return planet
}

func (s *HttpServer) Register(registerRouter func(mux *mux.Router)) *HttpServer {
	r := mux.NewRouter()
	registerRouter(r)
	s.s.Handler = r

	return s
}

func (s *HttpServer) HttpStart() {
	s.s.ListenAndServe()
}
