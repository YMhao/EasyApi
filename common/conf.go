package common

// APIServConf  is the configuration of the server
type APIServConf struct {
	Version     string // version of the server
	BuildTime   string // time build
	ServiceName string // server name
	Description string // server description
	ListenAddr  string // listen add , such as ":80", "192.168.1.1:8080"
	DebugOn     bool   // debug switch
	HTTPProxy   string // http proxy
}
