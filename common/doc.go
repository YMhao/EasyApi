package common

import (
	"reflect"
)

type SwaggerApiType string

const (
	SwaggerApiTypeJson     SwaggerApiType = "json"
	SwaggerApiTypeDownload SwaggerApiType = "download"
	SwaggerApiTypeUpload   SwaggerApiType = "upload"
)

type Parminfo struct {
	Name    string
	Desc    string
	IsArray bool
}

type ApiDoc struct {
	ApiId   string `json:"apiId"`
	ApiDesc string `json:"apiDesc"`
	Tag     string `json:"-"`
	Path    string `json:"-"`

	RequestDesc       string           `json:"requestDesc"`
	RequestObj        map[string]*Attr `json:"requestObj"`
	RequestDepObjList []DepObjDoc      `json:"requestDepObjList"`

	ResponseDesc       string           `json:"responseDesc"`
	ResponseObj        map[string]*Attr `json:"responseObj"`
	ResponseDepObjList []DepObjDoc      `json:"responseDepObjList"`

	SwaggerApiType   SwaggerApiType `json:"-"`
	Mime             []string       `json:"-"`
	QueryParmList    []Parminfo     `json:"query"`
	FormDataParmList []Parminfo     `json:"formdata"`
}

func (a *ApiDoc) ResponseName() string {
	return ""
}

func (a *ApiDoc) RequestName() string {
	t := reflect.TypeOf(a.RequestObj)
	return t.Name()
}
