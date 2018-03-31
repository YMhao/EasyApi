package serv

// APIServConf 是服务器配置
type APIServConf struct {
	Version     string // 版本号
	BuildTime   string // 编译时间
	ServiceName string // 服务名
	Description string // 服务的描述
	ListenAddr  string // 监听端口
	DebugOn     bool   // 是否启用debug调试页面
	HTTPProxy   string // 代理地址
}

// NewAPIServConf 创建一个服务配置
func NewAPIServConf(version, buildTime, serviceName, description string) *APIServConf {
	return &APIServConf{
		Version:     version,
		BuildTime:   buildTime,
		ServiceName: serviceName,
		Description: description,
		ListenAddr:  ":8089",
		DebugOn:     false,
		HTTPProxy:   "",
	}
}
