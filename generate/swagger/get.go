package swagger

import (
	"github.com/go-openapi/spec"
)

func NewGetFileOperation(description, tag string, statusCodeResponses map[int]spec.Response, mime []string) *spec.Operation {
	return &spec.Operation{
		OperationProps: spec.OperationProps{
			Summary:     Summary(description),
			Description: description,
			Produces:    mime,
			Tags:        []string{tag},
			Responses: &spec.Responses{
				ResponsesProps: spec.ResponsesProps{
					StatusCodeResponses: statusCodeResponses,
				},
			},
		},
	}
}
