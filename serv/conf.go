package serv

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

// NewAPIServConf create a new configuration of the server
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
