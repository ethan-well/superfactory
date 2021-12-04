package config

import (
	"flag"
	"github.com/ItsWewin/superfactory/xerror"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
)

type BasicYamlConf struct {
}

func NewBasicYamlConf() *BasicYamlConf {
	return &BasicYamlConf{}
}

func (c *BasicYamlConf) Unmarshal(configObject interface{}, fileName ...string) *xerror.Error {
	var configFile string
	flag.StringVar(&configFile, "cf", "test.conf", "-cf is expected")
	flag.Parse()

	if len(fileName) != 0 {
		configFile = fileName[0]
	}

	bt, err := os.ReadFile(configFile)
	if err != nil {
		return xerror.NewError(err, xerror.Code.OtherSystemError, "read file failed")
	}

	if reflect.TypeOf(configObject).Kind() != reflect.Ptr {
		return xerror.NewError(nil, xerror.Code.CParamsError, "config obj is must be a ptr")
	}

	err = yaml.Unmarshal(bt, configObject)
	if err != nil {
		return xerror.NewError(nil, xerror.Code.OtherSystemError, "unmarshal file failed")
	}

	return nil
}
