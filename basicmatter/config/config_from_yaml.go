package config

import (
	"flag"
	"github.com/ItsWewin/superfactory/aerror"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
)

type BasicYamlConf struct {
}

func NewBasicYamlConf() *BasicYamlConf {
	return &BasicYamlConf{}
}

func (c *BasicYamlConf) Unmarshal(configObject interface{}, fileName ...string) aerror.Error {
	var configFile string
	flag.StringVar(&configFile, "cf", "test.conf", "-cf is expected")
	flag.Parse()

	if len(fileName) != 0 {
		configFile = fileName[0]
	}

	bt, err := os.ReadFile(configFile)
	if err != nil {
		return aerror.NewError(err, aerror.Code.OtherSystemError, "read file failed")
	}

	if reflect.TypeOf(configObject).Kind() != reflect.Ptr {
		return aerror.NewError(nil, aerror.Code.CParamsError, "config obj is must be a ptr")
	}

	err = yaml.Unmarshal(bt, configObject)
	if err != nil {
		return aerror.NewError(nil, aerror.Code.OtherSystemError, "unmarshal file failed")
	}

	return nil
}
