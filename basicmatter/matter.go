package basicmatter

// Master 星球构成基本物质基础: 服务的基本配置项
type Master interface {
	MaterType() string
}

type ConfigMater interface {
	HttpAddr() string
}

const (
	MasterConfigBasicSection = "ConfigBasicSection"
	MasterConfigBasicYaml    = "ConfigBasicYaml"
)
