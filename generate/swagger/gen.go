package swagger

import (
	"strings"

	"github.com/YMhao/EasyApi/common"
	"github.com/go-openapi/spec"
)

func GenCode(conf *common.APIServConf, apiDocList []*common.ApiDoc) *Swagger {
	swagger := &Swagger{}
	swagger.Init()

	swagger.SetHost(GetUrl(conf.ListenAddr))
	swagger.SetBasePath("/")

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

func GenSchema(apiDoc *common.ApiDoc, schemaMap map[string]spec.Schema) {
	setReqSchema := func() {
		schema := &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: make(map[string]spec.Schema),
			},
		}

		for k, v := range apiDoc.RequestObj {
			schema.SchemaProps.Properties[k] = *itemToSchame(v)
		}
		schemaMap[apiDoc.ApiId+"Req"] = *schema

		for _, doc := range apiDoc.RequestDepObjList {
			schema = &spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: make(map[string]spec.Schema),
				},
			}
			for k, v := range doc.ObjectBody {
				schema.SchemaProps.Properties[k] = *itemToSchame(v)
			}
			objName := AvoidRepeatMap.GetTypeName(doc.PkgPath, doc.ObjectName)
			schemaMap[objName] = *schema
		}
	}

	setRspSchema := func() {
		schema := &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"HasError": spec.Schema{
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"boolean",
							},
						},
					},
					"ErrorDesc": spec.Schema{
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"string",
							},
						},
					},
					"Data": spec.Schema{
						SchemaProps: spec.SchemaProps{
							Type: spec.StringOrArray{
								"object",
							},
							Properties: make(map[string]spec.Schema),
						},
					},
				},
			},
		}

		for k, v := range apiDoc.ResponseObj {
			schema.SchemaProps.Properties["Data"].SchemaProps.Properties[k] = *itemToSchame(v)
		}
		schemaMap[apiDoc.ResponseName()] = *schema

		for _, doc := range apiDoc.ResponseDepObjList {
			schema = &spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: make(map[string]spec.Schema),
				},
			}
			for k, v := range doc.ObjectBody {
				schema.SchemaProps.Properties[k] = *itemToSchame(v)
			}
			objName := AvoidRepeatMap.GetTypeName(doc.PkgPath, doc.ObjectName)
			schemaMap[objName] = *schema
		}
	}

	switch apiDoc.SwaggerApiType {
	case common.SwaggerApiTypeDownload:
	case common.SwaggerApiTypeJson:
		setReqSchema()
		setRspSchema()
	case common.SwaggerApiTypeUpload:
		setRspSchema()
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

	for k, v := range apiDoc.ResponseObj {
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
	if apiDoc.SwaggerApiType == common.SwaggerApiTypeJson {
		return &spec.PathItem{
			PathItemProps: spec.PathItemProps{
				Parameters: func() []spec.Parameter {
					Parameters := []spec.Parameter{}
					Parameters = append(Parameters, *NewSwaggerQueryParamter("版本号", "v", true))
					Parameters = append(Parameters, *NewSwaggerSchemaRefParamter(GetApiId(apiDoc.ApiId)+"Req", true))
					return Parameters
				}(),
				Post: NewPostJsonOperation(apiDoc.ResponseDesc, apiDoc.Tag, statusCodeResponses),
			},
		}
	} else if apiDoc.SwaggerApiType == common.SwaggerApiTypeDownload {
		return &spec.PathItem{
			PathItemProps: spec.PathItemProps{
				Parameters: func() []spec.Parameter {
					Parameters := []spec.Parameter{}
					Parameters = append(Parameters, *NewSwaggerQueryParamter("版本号", "v", true))
					Parameters = append(Parameters, *NewSwaggerQueryParamter("文件id", "fileId", true))
					return Parameters
				}(),
				Get: NewGetFileOperation(apiDoc.ResponseDesc, apiDoc.Tag, statusCodeResponses, apiDoc.Mime),
			},
		}
	} else if apiDoc.SwaggerApiType == common.SwaggerApiTypeUpload {
		return &spec.PathItem{
			PathItemProps: spec.PathItemProps{
				Parameters: func() []spec.Parameter {
					Parameters := []spec.Parameter{}
					Parameters = append(Parameters, *NewSwaggerQueryParamter("版本号", "v", true))
					Parameters = append(Parameters, *NewSwaggerQueryParamter("会话id", "sessionId", false))
					return Parameters
				}(),
				Post: NewPostFileOperation(apiDoc.ResponseDesc, apiDoc.Tag, statusCodeResponses),
			},
		}
	}

	return &spec.PathItem{
		PathItemProps: spec.PathItemProps{
			Parameters: func() []spec.Parameter {
				Parameters := []spec.Parameter{}
				Parameters = append(Parameters, *NewSwaggerQueryParamter("版本号", "v", true))
				Parameters = append(Parameters, *NewSwaggerSchemaRefParamter(GetApiId(apiDoc.ApiId)+"Req", true))
				return Parameters
			}(),
			Post: NewPostJsonOperation(apiDoc.ResponseDesc, apiDoc.Tag, statusCodeResponses),
		},
	}

	panic("invalid api.SwaggerApiType: " + apiDoc.SwaggerApiType)
}

func GetApiId(id string) string {
	return strings.Replace(id, ".", "", -1)
}

func GetStatusCodeResponses(apiDoc *common.ApiDoc) map[int]spec.Response {
	if apiDoc.ResponseDesc == "" {
		apiDoc.ResponseDesc = "没有写描述"
	}

	switch apiDoc.SwaggerApiType {
	case common.SwaggerApiTypeJson:
		return map[int]spec.Response{
			200: spec.Response{ResponseProps: spec.ResponseProps{
				Description: apiDoc.ResponseDesc,
				Schema: &spec.Schema{
					SchemaProps: spec.SchemaProps{
						Ref: spec.MustCreateRef("#/definitions/" + GetApiId(apiDoc.ApiId) + "Rsp"),
					},
				},
			}},
		}
	case common.SwaggerApiTypeDownload:
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
	case common.SwaggerApiTypeUpload:
		return map[int]spec.Response{
			200: spec.Response{ResponseProps: spec.ResponseProps{
				Description: apiDoc.ResponseDesc,
				Schema: &spec.Schema{
					SchemaProps: spec.SchemaProps{
						Ref: spec.MustCreateRef("#/definitions/" + GetApiId(apiDoc.ApiId) + "Rsp"),
					},
				},
			}},
		}
	default:
		panic("invalid type")
	}
}
