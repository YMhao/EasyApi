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
