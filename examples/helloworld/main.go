package main

import (
	"github.com/YMhao/EasyApi/serv"
)

func main() {
	conf := serv.NewAPIServConf("1.0", "", "helloWorld", "EasyApi 的 hello World")
	serv.RunAPIServ(conf, &APIColl{})
}
