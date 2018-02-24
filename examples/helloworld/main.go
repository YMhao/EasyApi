package main

import (
	"encoding/json"

	"github.com/YMhao/EasyApi/serv"
)

// HelloRequest 是 hello api 的请求参数
type HelloRequest struct {
	Name string `json:"name" desc:"The request message containing the user's name."`
}

// HelloResponse 是 hello aoi 的响应参数
type HelloResponse struct {
	Message string `json:"message" desc:"The response message containing the greetings"`
}

// HelloAPI is a hello api
type HelloAPI struct {
}

// Doc api的文档
func (h HelloAPI) Doc() *serv.APIDoc {
	return &serv.APIDoc{
		ID:               "hello",
		Descript:         "helloworld service",
		RequestDescript:  "包含用户名的信息",
		Request:          &HelloRequest{},
		ResponseDescript: "包含问候信息",
		Response:         &HelloResponse{},
	}
}

// type API interface {
// 	Doc() *APIDoc
// 	Call(reqData []byte) (interface{}, *APIError)
// }

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
	return &HelloResponse{
		Message: "hello " + req.Name + "!",
	}, nil
}

// APIColl api集合
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

func main() {
	conf := serv.NewAPIServConf("1.0", "", "helloWorld", "EasyApi 的 hello World")
	serv.RunAPIServ(conf, &APIColl{})
}
