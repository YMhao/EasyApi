package main

import (
	"github.com/YMhao/EasyApi/serv"
)

type HelloRequest struct {
	Name string `desc:"The request message containing the user's name."`
}

type HelloRespone struct {
	Message string `desc:"The response message containing the greetings"`
}

var HelloAPI = serv.NewAPI(
	"helloWord",
	`api for helloword`,
	&HelloRequest{},
	&HelloRespone{},
	HelloAPICall,
)

func HelloAPICall(data []byte) (interface{}, *serv.APIError) {
	req := &HelloRequest{}
	err := serv.UnmarshalAndCheckValue(data, req)
	if err != nil {
		return nil, serv.NewError(err)
	}

	return &HelloRespone{
		Message: "hello " + req.Name + "!",
	}, nil
}

func main() {
	conf := serv.NewAPIServConf("1.0", "", "helloWorld", "EasyApi demo - hello World")
	conf.DebugOn = true
	conf.ListenAddr = ":8089"

	setsOfAPIs := serv.APISets{
		"MessageAPIs": []serv.API{
			HelloAPI,
		},
	}

	//conf.HTTPProxy = "http://yuminghao.top:8089"
	serv.RunAPIServ(conf, setsOfAPIs)
}
