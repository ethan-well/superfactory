package config

import (
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/basicmatter"
)

type ConfBasicObj interface {
	Unmarshal(configObject interface{}, fileName ...string) aerror.Error
}

func UnmarshalConf(conf basicmatter.Master, fileName ...string) aerror.Error {
	var confObj ConfBasicObj

	switch conf.MaterType() {
	case basicmatter.MasterConfigBasicSection:
		confObj = NewBasicSectionConf()
	case basicmatter.MasterConfigBasicYaml:
		confObj = NewBasicYamlConf()
	default:
		return aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "MaterType is not supported")
	}

	return confObj.Unmarshal(conf, fileName...)
}
