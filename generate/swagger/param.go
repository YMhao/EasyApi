package swagger

import "github.com/go-openapi/spec"

// 例如 ?v=xx
func NewSwaggerQueryParamter(description, name string, required bool) *spec.Parameter {
	return &spec.Parameter{
		ParamProps:   spec.ParamProps{Description: description, Name: name, In: "query", Required: required},
		SimpleSchema: spec.SimpleSchema{Type: "string"},
	}
}

// 表单里的参数
func NewSwaggerFormDataParamter(description, name string, required bool) *spec.Parameter {
	return &spec.Parameter{
		ParamProps:   spec.ParamProps{Description: description, Name: name, In: "formData", Required: required},
		SimpleSchema: spec.SimpleSchema{Type: "string"},
	}
}

// 文件上传
func NewSwaggerFileParamter(description, name string, required bool) *spec.Parameter {
	return &spec.Parameter{
		SimpleSchema: spec.SimpleSchema{Type: "file"},
		ParamProps: spec.ParamProps{
			Description: description,
			Name:        name,
			In:          "formData",
			Required:    required,
		},
	}
}

// schema
func NewSwaggerSchemaParamter(obj interface{}, required bool) *spec.Parameter {
	return &spec.Parameter{
		ParamProps: spec.ParamProps{Description: "请求参数", Name: "body", In: "body", Required: required, Schema: &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Ref: spec.MustCreateRef("#/definitions/demoRsp"),
			},
		}},
	}
}

// schema
func NewSwaggerSchemaRefParamter(name string, required bool) *spec.Parameter {
	return &spec.Parameter{
		ParamProps: spec.ParamProps{Description: "请求参数", Name: "body", In: "body", Required: required, Schema: &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Ref: spec.MustCreateRef("#/definitions/" + name),
			},
		}},
	}
}
