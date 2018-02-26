package main

import (
	"encoding/json"

	"github.com/YMhao/EasyApi/serv"
)

type Rectangle struct {
	Lo Point `json:"lo" desc:"One corner of the rectangle."`
	Hi Point `json:"hi" desc:"The other corner of the rectangle."`
}

type FeatureList struct {
	FeatureList []*Feature `json:"featureList" desc:"FeatureList"`
}

type ListFeatureAPi struct {
}

// Doc api的文档
func (l ListFeatureAPi) Doc() *serv.APIDoc {
	return &serv.APIDoc{
		ID:               "listFeature",
		Descript:         "Get Feature",
		RequestDescript:  "该请求对象为Rectangle",
		Request:          &Rectangle{},
		ResponseDescript: "该响应对象为FeatureList",
		Response:         &FeatureList{},
	}
}

// Call 回调
func (l ListFeatureAPi) Call(reqData []byte) (interface{}, *serv.APIError) {
	req := &Rectangle{}
	err := json.Unmarshal([]byte(reqData), req)
	if err != nil {
		return nil, &serv.APIError{
			Code:     "json.unmarshal",
			Descript: err.Error(),
		}
	}
	return &FeatureList{
		FeatureList: []*Feature{
			&Feature{
				Name: "Guangdong",
				Location: Point{
					Latitude:  1,
					Longitude: 2,
				},
			},
			&Feature{
				Name: "Beijing",
				Location: Point{
					Latitude:  1,
					Longitude: 2,
				},
			},
		},
	}, nil
}
