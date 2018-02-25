package common

// APIServConf 是服务器配置
type APIServConf struct {
	Version     string // 版本号
	BuildTime   string // 编译时间
	ServiceName string // 服务名
	Description string // 服务的描述
	ListenAddr  string // 监听端口
	DebugPage   bool   // 是否启用web调试页面
}
