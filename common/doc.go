package common

type SwaggerAPIType string

const (
	SwaggerAPITypeJson     SwaggerAPIType = "json"
	SwaggerAPITypeDownload SwaggerAPIType = "download"
	SwaggerAPITypeUpload   SwaggerAPIType = "upload"
)

type Parminfo struct {
	Name    string
	Desc    string
	IsArray bool
}

type ObjInfo struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Fields      map[string]*Attr `json:"fields"`
	DepObjList  []DepObjDoc      `json:"depObjList"`
	PkgPath     string           `json:"-"` // for avoiding repetition
}

type ApiDoc struct {
	ApiId   string `json:"apiId"`
	ApiDesc string `json:"apiDesc"`
	Tag     string `json:"-"`
	Path    string `json:"-"`

	Request  ObjInfo `json:"request"`
	Response ObjInfo `json:"response"`

	SwaggerAPIType   SwaggerAPIType `json:"-"`
	Mime             []string       `json:"-"`
	QueryParmList    []Parminfo     `json:"query"`
	FormDataParmList []Parminfo     `json:"formdata"`
}
