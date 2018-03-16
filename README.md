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
package main

import (
	"encoding/json"

	"github.com/YMhao/EasyApi/serv"
)

// HelloRequest 是 hello api 的请求参数
type HelloRequest struct {
	Name string `json:"name" desc:"The request message containing the user's name."`
}

// HelloResp 是 hello aoi 的响应参数
type HelloResp struct {
	Message string `json:"message" desc:"The response message containing the greetings"`
}

// HelloAPI is a hello api
// type API interface {
// 	Doc() *APIDoc
// 	Call(reqData []byte) (interface{}, *APIError)
// }
type HelloAPI struct {
}

// Doc api的文档
func (h HelloAPI) Doc() *serv.APIDoc {
	return &serv.APIDoc{
		ID:               "Hello",
		Descript:         "helloworld service",
		RequestDescript:  "该请求包含用户名信息",
		Request:          &HelloRequest{},
		ResponseDescript: "该响应包含问候信息",
		Response:         &HelloResp{},
	}
}

// Call 回调
func (h HelloAPI) Call(reqData []byte) (interface{}, *serv.APIError) {
	req := &HelloRequest{}
	err := json.Unmarshal([]byte(reqData), req)
	if err != nil {
		return nil, &serv.APIError{
			Code:     "json.unmarshal",
			Descript: err.Error(),
		}
	}

	return &HelloResp{
		Message: "hello " + req.Name + "!",
	}, nil
}

```

api_coll.go
```
// APIColl api集合
import "github.com/YMhao/EasyApi/serv"

type APIColl struct {
}

// AllAPI 列出所有api
func (a APIColl) AllAPI() map[serv.CateName][]serv.API {
	return map[serv.CateName][]serv.API{
		"helloServ": []serv.API{
			&HelloAPI{},
		},
	}
}

```


main.go
```

package main

import (
	"github.com/YMhao/EasyApi/serv"
)

func main() {
	conf := serv.NewAPIServConf("1.0", "", "helloWorld", "EasyApi 的 hello World")
	serv.RunAPIServ(conf, &APIColl{})
}


```
