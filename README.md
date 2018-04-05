# 概述

超级简单易用的 api 框架

# 功能：
1、完整的api文档自动生成。   
2、生成swagger，借助swagger 生成调试页面、生成多语言客户端代码，方便前端和测试的对api的调用和测试。   
3、支持字段的枚举(enum)、最小值限制(min)、最大值限制(max)。   
4、不是通过注释生成文档的，时刻保证文档跟代码是一致的。   
5、编码过程不需要安装额外工具（额外工具的缺点：1、要安装，2、要学工具的命令行参数，3、出错还得花时间去找错误的地方)。    
5、实现同样多的功能的条件下，能以更少的代码量、和更简洁的代码完成需求、还能充分利用 IDE的自动补齐、错误检查。    
6、基于gin，gin性能有多好，该框架性能就有多好。   
7、‘0’ 学习成本， 你只需要记住 enum 是枚举， min是下限, max 是上限， desc是描述， serv.NewAPI 是新建一个api的接口， serv.NewAPIServConf 是新建服务的配置，serv.RunAPIServ 是启动服务就足够了。   
8、有debug的开关。   
9、有错误的规范。    
10、可扩展。   

# 错误处理

APIError 是 框架返回的标准错误   

错误码 为什么用string 而不采用数字    
有以下几点思考：   
1、传统时用一个数字代表一个code，但是数字是看不出错误大概时什么，想要知道，得查错误码表。   
2、希望Code能用一个简短的字符串来代表一个code，直接看code就可以知道大概时一个什么错误。   
3、虽然希望时可以用一个简短字符串，但是不强求，因为数字也是一个字符串。

# 未来的实现

1、文件上下传  
2、默认不填`json`属性时，可根据规范, 自动格式化 json请求的字段格式（骆峰、下划线分割等）,通过改写json库实现   
3、grpc 的支持（不需要手写.proto文件，自动化）   
4、接口监控、统计、错误收集。   
5、还没想好  

# 字段属性

字段属性有:`json`,`desc`,`enum`,`max`,`min`, 即 json的键名、 描述、枚举、上限、下限。   
每个字段属性都是选填的， desc建议填写（为了减少误解）, 属性之间用空格隔开。   

## 字段属性-json的键名

关键词:`json`   

如果不是json的键名太奇葩，不建议设置该键，不填时，默认以struct里的字段名命名。   

例子：

```golang
type XXXRequest struct {
	MessageType string `json:"message_type"`
}
```

## 字段属性-描述

关键词:`desc`   
使用建议：字段的描述可以选填，写比较好，可以减少一些误解   

例子：

```golang
type XXXRequest struct {
	MessageType string `desc:"消息类型"`
}
```

## 字段属性-枚举

枚举只支持string类型的,如何定义枚举的例子:   

`enum:"TYPE1,TYPE2,TYPEN"`

枚举值之间用英文字符`,`或空格分开   
```golang
type XXXRequest struct {
	MessageType string `desc:"消息类型" enum:"TYPE1,TYPE2"`
}
```

## 字段属性-上下限

下限:`min`   
上限制:`max`   
上下限，只支持数字类型，上下限可以同时存在,也可以只存在其中一个，也可以都不存在   

```golang
type XXXRequest struct {
	offset int `desc:"偏移" min:"0" max:"1000"`
}
```

# 注意 

## map 不支持！

 支持 请求和响应的struct里字段有map类型的，例如下面的Struct A是不支持的   
 ```golang
 type A struct{
	M map[string]string
 }
 ```
如果想支持map，可以通过其他方式，例如下面的 Struct B  

 ```golang
type C struct {
	Key   string `desc:"关键词"`
	Value string `desc:"值"`
}

 type B struct {
	M []*C `desc:"一个map"`
 }
 ```

 如果你的map key是有限的， 可以把key作为字段名   

## struct的字段的字段名是要以大写字符开头的才能导出和导入数据

能导出导入数据：
```golang
type A struct {
	Name string
}
```

不能导出导入数据
```golang
type A struct {
	name string
}
```


# 示例代码

两个自带例子：   
代码: https://github.com/YMhao/EasyApi/tree/master/examples/helloworld   
服务: http://yuminghao.top:8089


代码: https://github.com/YMhao/EasyApi/tree/master/examples/routeGuide
服务: http://yuminghao.top:8088


下面展示如何写一个helloworld的服务端   

hello.go
```
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

	serv.RunAPIServ(conf, setsOfAPIs)
}

```
