package serv

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// APIDoc 包含以下属性
// 1、api的描述
// 2、api的id
// 3、请求参数的描述
// 4、请求对象（一个struct）
// 4、响应参数的描述
// 5、响应对象（一个struct）
//
// 问：
// 为什么只要包含以上属性，就能获取到该api的所有文档信息呢?
//
// 答：
// Go的反射机制能够在程序运行过程中检查自身元素的结构，只要我们把文档信息写到struct里面就可以了。
// 举个struct的例子， 例如 请求对象GetFeatureRequest 和 响应对象GetFeatureResponse：
//
// type Point struct {
// 	Latitude  int `json:"latitude" desc:"纬度"`
// 	Longitude int `json:"longitude" desc:"纬度"`
// }
//
// type GetFeatureRequest struct {
// 	Location Point `json:"location" desc:"位置"`
// }
//
// type GetFeatureResponse struct {
// 	Name string `json:"name" desc:"特征的名称"`
// }
//
// 如上所示，文档的信息也就是记录在struct的tag内容里面。
// 同样，用反射也可以得到一个请求、或响应的内容模板，
// 例如:
// 请求对象GetFeatureRequest 的内容模板是：
// {
//		“location” ： {
//			"latitude" : 0,
//      	"Longitude" : 0
//		}
// }
// 响应对象GetFeatureResponse 的内容模板是：
// {
//		"name" : ""
// }
type APIDoc struct {
	ID               string      // api id
	Descript         string      // api的描述api
	RequestDescript  string      // 请求参数的描述
	Request          interface{} // 请求对象（一个struct）
	ResponseDescript string      // 响应参数的描述
	Response         interface{} // 响应对象（一个struct）
}

// APIError 是 框架返回的标准错误
// 错误码 为什么用string 而不采用数字
// 有以下几个思考：
// 1、传统时用一个数字代表一个code，但是数字是看不出错误大概时什么，想要知道，得查错误码表。
// 2、希望Code能用一个简短的字符串来代表一个code，直接看code就可以知道大概时一个什么错误。
// 3、虽然希望时可以用一个简短字符串，但是不强求，因为数字也是一个字符串。
type APIError struct {
	Code    string `json:"code" desc:"错误码"`
	Message string `json:"messge" desc:"错误信息"`
}

// API 是个 接口
type API interface {
	Doc() *APIDoc
	Call([]byte, *gin.Context) (interface{}, *APIError)
}

// APISets  the sets of APIs,  map key is the name of set
type APISets map[string][]API

type CommonAPI struct {
	id               string
	descript         string
	requestDescript  string
	request          interface{}
	responseDescript string
	response         interface{}
	call             func([]byte, *gin.Context) (interface{}, *APIError)
}

func (comm *CommonAPI) Doc() *APIDoc {
	return &APIDoc{
		ID:               comm.id,
		Descript:         comm.descript,
		RequestDescript:  comm.requestDescript,
		Request:          comm.request,
		ResponseDescript: comm.responseDescript,
		Response:         comm.response,
	}
}

func (comm *CommonAPI) Call(data []byte, c *gin.Context) (interface{}, *APIError) {
	if comm.call != nil {
		return comm.call(data, c)
	}
	return nil, NewError(errors.New("the callbacke function is not found"))
}

func (comm *CommonAPI) SetCallback(call func([]byte, *gin.Context) (interface{}, *APIError)) {
	comm.call = call
}

// NewAPI create a new api with some info
// APIName: be used to router
// APIDesc: the descript of this api
// request: a struct, http request
// respone: a struct, http response
// callback: a callback function, func([]byte, *gin.Context) (interface{}, *APIError)
func NewAPI(APIName, APIDesc string, request interface{}, response interface{},
	callback func([]byte, *gin.Context) (interface{}, *APIError)) API {
	return &CommonAPI{
		id:               APIName,
		descript:         formatDescript(APIDesc),
		requestDescript:  "request context",
		request:          request,
		responseDescript: "respone context",
		response:         response,
		call:             callback,
	}
}
