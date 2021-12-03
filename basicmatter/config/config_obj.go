package config

import (
	"github.com/ItsWewin/superfactory/basicmatter"
	"github.com/ItsWewin/superfactory/xerror"
)

type ConfBasicObj interface {
	Unmarshal(configObject interface{}, fileName ...string) *xerror.Error
}

func UnmarshalConf(conf basicmatter.Master, fileName ...string) *xerror.Error {
	var confObj ConfBasicObj

	switch conf.MaterType() {
	case basicmatter.MasterConfigBasicSection:
		confObj = NewBasicSectionConf()
	case basicmatter.MasterConfigBasicYaml:
		confObj = NewBasicYamlConf()
	default:
		return xerror.NewErrorf(nil, xerror.Code.BUnexpectedData, "MaterType is not supported")
	}

	return confObj.Unmarshal(conf, fileName...)
}
