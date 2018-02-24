package serv

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
	Code     string `desc:"错误码"`
	Descript string `desc:"错误的描述"`
}

// API 是个 接口
type API interface {
	Doc() *APIDoc
	Call(reqData []byte) (interface{}, *APIError)
}
