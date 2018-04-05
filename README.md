# 概述

这是一个前端、后端、测试友好型服务器框架。    
能自动生成调试页及api文档，支持swagger扩展，利用swagger可生成各种语言客户端及服务端代码.   

目前只做了基础功能。    
还在增加功能中， 敬请期待。。。

# 示例代码

两个自带例子：
https://github.com/YMhao/EasyApi/tree/master/examples/helloworld   

https://github.com/YMhao/EasyApi/tree/master/examples/routeGuide


下面展示如何写一个helloworld的服务端

hello.go
```
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
	conf := serv.NewAPIServConf("1.0", "", "helloWorld", "EasyApi 的 hello World")
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

```
