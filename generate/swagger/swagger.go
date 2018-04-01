package swagger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/go-openapi/spec"
)

type Swagger struct {
	Swagger spec.Swagger
}

func (s *Swagger) Init() {
	s.Swagger.SwaggerProps.Swagger = "2.0"
	s.Swagger.SwaggerProps.Schemes = []string{"http"}
}

func (s *Swagger) SetPaths(paths *spec.Paths) {
	s.Swagger.SwaggerProps.Paths = paths
}

func (s *Swagger) SetBasePath(BasePath string) {
	s.Swagger.SwaggerProps.BasePath = BasePath
}

func (s *Swagger) SetHost(host string) {
	if strings.HasPrefix(host, "http://") {
		host = host[7:]
	}
	if strings.HasPrefix(host, "https://") {
		s.Swagger.SwaggerProps.Schemes = []string{"https"}
		host = host[8:]
	}
	s.Swagger.SwaggerProps.Host = host
}

func (s *Swagger) SetInfo(info *spec.Info) {
	if info.Title == "" {
		info.Title = "no_defined"
		fmt.Println("Warn:ApiServConf.Title is not defined")
	}
	if info.Description == "" {
		fmt.Println("Warn:ApiServConf.Description is not defined")
	}
	if info.Version == "" {
		fmt.Println("Warn:ApiServConf.Version is not defined")
	}

	s.Swagger.SwaggerProps.Info = info
}

func (s *Swagger) SetDefinitions(definitions map[string]spec.Schema) {
	s.Swagger.SwaggerProps.Definitions = definitions
}

func (s *Swagger) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(s.Swagger, "", "    ")
}

func (s *Swagger) MarshalYAML() ([]byte, error) {
	data, err := s.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return yaml.JSONToYAML(data)
}

type StdError struct {
	HasError  bool
	ErrorDesc string
}

func (s *Swagger) PostSwaggerUi(url string) error {
	if url == "" {
		url = "http://192.168.1.113:8000/call?id=test.AddProject&v=xx"
	}
	data, err := s.MarshalJSON()
	if err != nil {
		return err
	}

	rsp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	if rsp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("StatusCode: %d", rsp.StatusCode))
	}

	data, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	e := &StdError{}
	err = json.Unmarshal(data, e)
	if err != nil {
		return err
	}
	if e.HasError {
		return errors.New(e.ErrorDesc)
	}
	return nil
}
