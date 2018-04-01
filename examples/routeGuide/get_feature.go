package main

import (
	"encoding/json"

	"github.com/YMhao/EasyApi/serv"
)

type Point struct {
	Latitude  int `json:"latitude" desc:"latitude value" min:"0" max:"90"`
	Longitude int `json:"longitude" desc:"longitude value" min:"0" max:"180"`
}

type GetFeatureAPi struct {
}

type Feature struct {
	Name     string `json:"Name" desc:"feature name" enum:"GuangZhou,Beijing,Shenzhen,Shanghai"`
	Location Point  `json:"location" desc:"location"`
}

func (g GetFeatureAPi) Doc() *serv.APIDoc {
	return &serv.APIDoc{
		ID:               "getFeature",
		Descript:         "Obtains the feature at a given position",
		RequestDescript:  "该请求内容为Point",
		Request:          &Point{},
		ResponseDescript: "该响应内容为Feature",
		Response:         &Feature{},
	}
}

func (g GetFeatureAPi) Call(reqData []byte) (interface{}, *serv.APIError) {
	req := &Point{}
	err := json.Unmarshal([]byte(reqData), req)
	if err != nil {
		return nil, &serv.APIError{
			Code:     "json.unmarshal",
			Descript: err.Error(),
		}
	}
	return &Feature{
		Name:     "Guangdong",
		Location: *req,
	}, nil
}
