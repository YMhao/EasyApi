package swagger

import "github.com/go-openapi/spec"

func NewPostJsonOperation(description, tag string, statusCodeResponses map[int]spec.Response) *spec.Operation {
	return &spec.Operation{
		OperationProps: spec.OperationProps{
			Summary:     Summary(description),
			Description: description,
			Produces:    []string{"application/json"},
			Consumes:    []string{"application/json"},
			Tags:        []string{tag},
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: statusCodeResponses,
				},
			},
		},
	}
}

func NewPostFileOperation(description, tag string, statusCodeResponses map[int]spec.Response) *spec.Operation {
	Parameters := []spec.Parameter{}
	Parameters = append(Parameters, *NewSwaggerFileParamter("文件1", "file1", false))
	Parameters = append(Parameters, *NewSwaggerFileParamter("文件2", "file2", false))
	Parameters = append(Parameters, *NewSwaggerFileParamter("文件3", "file3", false))

	return &spec.Operation{
		OperationProps: spec.OperationProps{
			Summary:     Summary(description),
			Description: description,
			Consumes:    []string{"multipart/form-data"},
			Produces:    []string{"application/json"},
			Tags:        []string{tag},
			Parameters:  Parameters,
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: statusCodeResponses,
				},
			},
		},
	}
}

func NewPostFileJsonOperation(description, tag, refJson string, statusCodeResponses map[int]spec.Response) *spec.Operation {
	Parameters := []spec.Parameter{}
	Parameters = append(Parameters, *NewSwaggerFormDataParamter("请求的json: 将对象["+refJson+"]转换成json格式字符串，然后写入这里", "reqJson", true))
	Parameters = append(Parameters, *NewSwaggerFileParamter("文件1", "file1", false))
	Parameters = append(Parameters, *NewSwaggerFileParamter("文件2", "file2", false))
	Parameters = append(Parameters, *NewSwaggerFileParamter("文件3", "file3", false))

	return &spec.Operation{
		OperationProps: spec.OperationProps{
			Summary:     Summary(description),
			Description: description,
			Consumes:    []string{"multipart/form-data"},
			Produces:    []string{"application/json"},
			Tags:        []string{tag},
			Parameters:  Parameters,
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: statusCodeResponses,
				},
			},
		},
	}
}
