package swagger

import (
	"encoding/json"
	"fmt"
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
