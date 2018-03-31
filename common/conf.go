package common

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
