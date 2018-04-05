package main

import (
	"github.com/YMhao/EasyApi/serv"
)

func main() {
	conf := serv.NewAPIServConf("1.0", "", "routeGuide", "EasyApi - routeGuide server demo")
	conf.DebugOn = true
	conf.ListenAddr = ":8088"

	setsOfAPIs := serv.APISets{
		"MessageAPIs": []serv.API{
			GetFeatureAPi,
			ListFeatureAPi,
		},
	}
	//conf.HTTPProxy = "yuminghao.top:8088"
	serv.RunAPIServ(conf, setsOfAPIs)
}
