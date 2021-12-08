package config_test

import (
	"github.com/ItsWewin/superfactory/basicmatter/config"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

type Config struct {
	Server *Server           `yaml:"server"`
	Mysql  Mysqls            `yaml:"mysql"`
	Redis  map[string]*Redis `yaml:"redis"`
}

type Server struct {
	Name       string      `yaml:"name"`
	HttpServer *HttpServer `yaml:"http_server"`
	GrpcServer *GrpcServer `yaml:"grpc_server"`
}

type HttpServer struct {
	Addr string `yaml:"addr"`
}

type GrpcServer struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
}

type Mysqls struct {
	Plutodb *Mysql `yaml:"plutodb"`
	Callito *Mysql `yaml:"callito"`
}

type Mysql struct {
	DB           string `yaml:"db"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Addr         string `yaml:"addr"`
	Protocol     string `yaml:"protocol"`
	MaxLifeTime  int    `yaml:"max_life_time"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

type Redis struct {
	Addr string `yaml:"addr"`
	DBNo string `yaml:"db_no"`
}

func TestBasicYamlConf(t *testing.T) {
	var conf Config
	err := config.NewBasicYamlConf().Unmarshal(&conf, "testing_conf.yaml")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("conf: %s", logger.ToJson(conf))
	t.Log(conf.Mysql.Plutodb.Addr)

	t.Log("succeed")
}
