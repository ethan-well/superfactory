package config

import (
	"github.com/ItsWewin/superfactory/xerror"
)

type ConfBasicObj interface {
	Unmarshal(configObject interface{}, fileName string) *xerror.Error
}