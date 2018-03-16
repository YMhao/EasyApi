package main

import (
	"github.com/YMhao/EasyApi/serv"
)

func main() {
	conf := serv.NewAPIServConf("1.0", "", "routeGuide", "EasyApi 的 routeGuide server demo")
	conf.ListenAddr = ":8088"
	serv.RunAPIServ(conf, &APIColl{})
}
