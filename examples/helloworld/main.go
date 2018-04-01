package main

import (
	"github.com/YMhao/EasyApi/serv"
)

func main() {
	conf := serv.NewAPIServConf("1.0", "", "helloWorld", "EasyApi 的 hello World")
	conf.DebugOn = true
	conf.ListenAddr = ":8089"
	//conf.HTTPProxy = "http://yuminghao.top:8089"
	serv.RunAPIServ(conf, &APIColl{})
}
