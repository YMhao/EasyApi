package swagger

import (
	"strings"

	"github.com/YMhao/EasyApi/common"
	"github.com/go-openapi/spec"
)

func GetHostFromConf(httpProxy, listenAddr string) string {
	if httpProxy != "" {
		return httpProxy
	}
	return AutoGetUrl(listenAddr)
}

func GenCode(conf *common.APIServConf, apiDocList []*common.ApiDoc) *Swagger {
	swagger := &Swagger{}
	swagger.Init()

	swagger.SetHost(GetHostFromConf(conf.HTTPProxy, conf.ListenAddr))
	swagger.SetBasePath("/" + conf.ServiceName + "/" + conf.Version)

	info := GenInfo(conf.Version, conf.ServiceName, conf.Description)
	swagger.SetInfo(info)

	paths := GenPaths(apiDocList)
	swagger.SetPaths(paths)

	definitions := GenDefinitions(apiDocList)
	swagger.SetDefinitions(definitions)
	return swagger
}

func GenInfo(version, title, desc string) *spec.Info {
	return &spec.Info{
		InfoProps: spec.InfoProps{
			Version:     version,
			Title:       title,
			Description: desc,
		},
	}
}

func GenDefinitions(apiDocList []*common.ApiDoc) map[string]spec.Schema {
	defs := make(map[string]spec.Schema)
	for _, apidoc := range apiDocList {
		GenSchema(apidoc, defs)
	}
	return defs
}

func setReqSchema(apiDoc *common.ApiDoc, schemaMap map[string]spec.Schema) {
	schema := &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Properties: make(map[string]spec.Schema),
		},
	}

	for k, v := range apiDoc.Request.Fields {
		schema.SchemaProps.Properties[k] = *itemToSchame(v)
	}
	objName := AvoidRepeatMap.GetTypeName(apiDoc.Request.PkgPath, apiDoc.Request.Name)
	schemaMap[objName] = *schema

	for _, doc := range apiDoc.Request.DepObjList {
		schema = &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: make(map[string]spec.Schema),
			},
		}
		for k, v := range doc.Fields {
			schema.SchemaProps.Properties[k] = *itemToSchame(v)
		}
		objName := AvoidRepeatMap.GetTypeName(doc.PkgPath, doc.Name)
		schemaMap[objName] = *schema
	}
}

func setRspSchema(apiDoc *common.ApiDoc, schemaMap map[string]spec.Schema) {
	schema := &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Properties: make(map[string]spec.Schema),
		},
	}
	for k, v := range apiDoc.Response.Fields {
		schema.SchemaProps.Properties[k] = *itemToSchame(v)
	}
	objName := AvoidRepeatMap.GetTypeName(apiDoc.Response.PkgPath, apiDoc.Response.Name)
	schemaMap[objName] = *schema

	for _, doc := range apiDoc.Response.DepObjList {
		schema = &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: make(map[string]spec.Schema),
			},
		}
		for k, v := range doc.Fields {
			schema.SchemaProps.Properties[k] = *itemToSchame(v)
		}
		objName := AvoidRepeatMap.GetTypeName(doc.PkgPath, doc.Name)
		schemaMap[objName] = *schema
	}
}

func GenSchema(apiDoc *common.ApiDoc, schemaMap map[string]spec.Schema) {
	switch apiDoc.SwaggerAPIType {
	case common.SwaggerAPITypeDownload:
	case common.SwaggerAPITypeJson:
		setReqSchema(apiDoc, schemaMap)
		setRspSchema(apiDoc, schemaMap)
	case common.SwaggerAPITypeUpload:
		setRspSchema(apiDoc, schemaMap)
	default:
		return
	}
}

func GenRspSchema(apiDoc *common.ApiDoc) *spec.Schema {
	schema := &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Properties: make(map[string]spec.Schema),
		},
	}

	for k, v := range apiDoc.Response.Fields {
		schema.SchemaProps.Properties[k] = *itemToSchame(v)
	}
	return schema
}

func GenPaths(apiDocList []*common.ApiDoc) *spec.Paths {
	paths := &spec.Paths{
		Paths: make(map[string]spec.PathItem),
	}
	for _, apiDoc := range apiDocList {
		paths.Paths[apiDoc.Path] = *GenPathItem(apiDoc)
	}
	return paths
}

func GenPathItem(apiDoc *common.ApiDoc) *spec.PathItem {
	statusCodeResponses := GetStatusCodeResponses(apiDoc)
	if apiDoc.SwaggerAPIType == common.SwaggerAPITypeJson {
		return &spec.PathItem{
			PathItemProps: spec.PathItemProps{
				Parameters: func() []spec.Parameter {
					Parameters := []spec.Parameter{}
					Parameters = append(Parameters, *NewSwaggerSchemaRefParamter(GetApiId(apiDoc.Request.Name), apiDoc.Request.Description, true))
					return Parameters
				}(),
				Post: NewPostJsonOperation(apiDoc.ApiDesc, apiDoc.Tag, statusCodeResponses),
			},
		}
	} else if apiDoc.SwaggerAPIType == common.SwaggerAPITypeDownload {
		return &spec.PathItem{
			PathItemProps: spec.PathItemProps{
				Parameters: func() []spec.Parameter {
					Parameters := []spec.Parameter{}
					Parameters = append(Parameters, *NewSwaggerQueryParamter("文件id", "fileId", true))
					return Parameters
				}(),
				Get: NewGetFileOperation(apiDoc.Response.Description, apiDoc.Tag, statusCodeResponses, apiDoc.Mime),
			},
		}
	} else if apiDoc.SwaggerAPIType == common.SwaggerAPITypeUpload {
		return &spec.PathItem{
			PathItemProps: spec.PathItemProps{
				Parameters: func() []spec.Parameter {
					Parameters := []spec.Parameter{}
					Parameters = append(Parameters, *NewSwaggerQueryParamter("会话id", "sessionId", false))
					return Parameters
				}(),
				Post: NewPostFileOperation(apiDoc.Response.Description, apiDoc.Tag, statusCodeResponses),
			},
		}
	}

	return &spec.PathItem{
		PathItemProps: spec.PathItemProps{
			Parameters: func() []spec.Parameter {
				Parameters := []spec.Parameter{}
				Parameters = append(Parameters, *NewSwaggerSchemaRefParamter(GetApiId(apiDoc.Request.Name), apiDoc.Request.Description, true))
				return Parameters
			}(),
			Post: NewPostJsonOperation(apiDoc.ApiDesc, apiDoc.Tag, statusCodeResponses),
		},
	}

	panic("invalid api.SwaggerAPIType: " + apiDoc.SwaggerAPIType)
}

func GetApiId(id string) string {
	return strings.Replace(id, ".", "", -1)
}

func getStdRespSchema(apiDoc *common.ApiDoc) *spec.Schema {
	objName := AvoidRepeatMap.GetTypeName(apiDoc.Response.PkgPath, apiDoc.Response.Name)
	schema := &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Properties: map[string]spec.Schema{
				"hasError": spec.Schema{
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{
							"boolean",
						},
					},
				},
				"error": spec.Schema{
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{
							"object",
						},
						Properties: make(map[string]spec.Schema),
					},
				},
				"data": spec.Schema{
					SchemaProps: spec.SchemaProps{
						Ref: spec.MustCreateRef("#/definitions/" + objName),
					},
				},
			},
		},
	}
	schema.SchemaProps.Properties["error"].SchemaProps.Properties["code"] = *spec.StringProperty()
	schema.SchemaProps.Properties["error"].SchemaProps.Properties["message"] = *spec.StringProperty()
	return schema
}

func GetStatusCodeResponses(apiDoc *common.ApiDoc) map[int]spec.Response {
	switch apiDoc.SwaggerAPIType {
	case common.SwaggerAPITypeJson:
		return map[int]spec.Response{
			200: spec.Response{ResponseProps: spec.ResponseProps{
				Description: apiDoc.Response.Description,
				Schema:      getStdRespSchema(apiDoc),
			}},
		}
	case common.SwaggerAPITypeDownload:
		return map[int]spec.Response{
			200: spec.Response{ResponseProps: spec.ResponseProps{
				Description: "文件",
				Schema: &spec.Schema{
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{
							"file",
						},
					},
				},
			}},
		}
	case common.SwaggerAPITypeUpload:
		return map[int]spec.Response{
			200: spec.Response{ResponseProps: spec.ResponseProps{
				Description: apiDoc.Response.Description,
				Schema: &spec.Schema{
					SchemaProps: spec.SchemaProps{
						Ref: spec.MustCreateRef("#/definitions/" + apiDoc.ApiId + "Resp"),
					},
				},
			}},
		}
	default:
		panic("invalid type " + apiDoc.SwaggerAPIType)
	}
}
