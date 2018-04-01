package serv

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/YMhao/EasyApi/common"
	"github.com/YMhao/EasyApi/generate/swagger"
	"github.com/YMhao/EasyApi/web"
	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
)

// RunAPIServ 启动服务
func RunAPIServ(conf *APIServConf, apiColl APICollect) {
	router := gin.Default()
	web.SetHTMLTemplate(router)
	allAPI := apiColl.AllAPI()
	for _, apiList := range allAPI {
		for _, api := range apiList {
			apiDoc := api.Doc()
			path := getPath(conf, apiDoc.ID)
			router.POST(path, func(c *gin.Context) {
				runAPICall(api, c)
			})
			router.OPTIONS(path, func(c *gin.Context) {
				runOpthionCall(c)
			})
		}
	}
	if conf.DebugOn {
		initHTML(conf, apiColl, router)
	}
	router.Run(conf.ListenAddr)
}

func getSwaggerProtocolURL(conf *APIServConf, path string) string {
	url := swagger.GetHostFromConf(conf.HTTPProxy, conf.ListenAddr) + path
	if strings.HasPrefix(conf.HTTPProxy, "https://") {
		return url
	}
	if strings.HasPrefix(conf.HTTPProxy, "http://") {
		return url
	}
	return "http://" + url
}

func jsonToYaml(jsonStr string) (string, error) {
	yamlBytes, err := yaml.JSONToYAML([]byte(jsonStr))
	return string(yamlBytes), err
}

func initHTML(conf *APIServConf, apiColl APICollect, router *gin.Engine) {
	swaggerJSONStr, err := genSwagger(conf, apiColl)
	if err != nil {
		fmt.Println("Warn: ", err)
	}
	swaggerYAMLStr, err := jsonToYaml(swaggerJSONStr)
	if err != nil {
		fmt.Println("Warn: ", err)
	}
	rawSwaggerProtocolJSONPath := "/swaggerJSON"
	rawSwaggerProtocolYAMLPath := "/swaggerYAML"
	index := web.IndexInfo{
		Name:        conf.ServiceName,
		Description: conf.Description,
		URL:         getSwaggerProtocolURL(conf, rawSwaggerProtocolYAMLPath),
		SwaggerJSON: swaggerJSONStr,
		SwaggerYAML: swaggerYAMLStr,
	}
	router.GET("/", func(c *gin.Context) {
		cors(c, "*")
		c.HTML(200, "Index", index.HTMLIndexInfo())
	})
	router.OPTIONS("/", func(c *gin.Context) {
		runOpthionCall(c)
	})
	router.GET(rawSwaggerProtocolJSONPath, func(c *gin.Context) {
		cors(c, "*")
		c.String(200, swaggerJSONStr)
	})
	router.OPTIONS(rawSwaggerProtocolJSONPath, func(c *gin.Context) {
		runOpthionCall(c)
	})
	router.GET(rawSwaggerProtocolYAMLPath, func(c *gin.Context) {
		cors(c, "*")
		c.String(200, swaggerYAMLStr)
	})
	router.OPTIONS(rawSwaggerProtocolYAMLPath, func(c *gin.Context) {
		runOpthionCall(c)
	})
}

func genSwagger(conf *APIServConf, apiColl APICollect) (swagerJSON string, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
			return
		}
	}()

	docList := genAPIDocList(conf, apiColl)
	swagger := swagger.GenCode(&common.APIServConf{
		Version:     conf.Version,
		BuildTime:   conf.BuildTime,
		ServiceName: conf.ServiceName,
		Description: conf.Description,
		ListenAddr:  conf.ListenAddr,
		DebugOn:     conf.DebugOn,
		HTTPProxy:   conf.HTTPProxy,
	}, docList)
	data, err := swagger.MarshalJSON()
	if err != nil {
		return "", nil
	}
	return string(data), nil
}

func getPath(conf *APIServConf, apiID string) string {
	return "/" + conf.ServiceName + "/" + conf.Version + "/" + apiID
}

func runAPICall(api API, c *gin.Context) {
	cors(c, "*")
	contentType := c.Request.Header.Get("Content-Type")
	contentType = strings.ToLower(contentType)
	if strings.Contains(contentType, "application/json") {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			handleError(c, &APIError{
				Code:     "unknown",
				Descript: err.Error(),
			})
			return
		}
		response, apiErr := api.Call(body)
		if apiErr != nil {
			handleError(c, apiErr)
			return
		}
		handleResponse(c, response)
	} else {
		handleError(c, &APIError{
			Code:     "unknown",
			Descript: "invalid content-type " + contentType,
		})
	}
}

func handleError(c *gin.Context, apiErr *APIError) {
	c.Writer.Header().Set("content-type", "application/json")
	c.JSON(200, &CallResponse{
		HasError: true,
		Error:    apiErr,
	})
}

func handleResponse(c *gin.Context, response interface{}) {
	c.Writer.Header().Set("content-type", "application/json")
	c.JSON(200, &CallResponse{
		Data: response,
	})
}

func runOpthionCall(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Access-Control-Allow-Method,Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("content-type", "application/json")

	c.JSON(200, gin.H{})
}

func cors(c *gin.Context, origin string) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Access-Control-Allow-Method,Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
}
