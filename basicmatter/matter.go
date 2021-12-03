package basicmatter

// 星球构成基本物质基础: 服务的基本配置项
type Master interface {
	MaterType() string
}

func New(m Master) * Master {
	switch m.MaterType() {
	case "config": {}
	}
}