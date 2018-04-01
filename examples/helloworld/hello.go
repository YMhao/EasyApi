package main

import (
	"encoding/json"

	"github.com/YMhao/EasyApi/serv"
)

type HelloRequest struct {
	Name string `json:"name" desc:"The request message containing the user's name."`
}

type HelloResp struct {
	Message string `json:"message" desc:"The response message containing the greetings"`
}

type HelloAPI struct {
}

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
