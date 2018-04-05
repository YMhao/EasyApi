package main

import (
	"github.com/YMhao/EasyApi/serv"
)

type Point struct {
	Latitude  int `json:"latitude" desc:"latitude value" min:"0" max:"90"`
	Longitude int `json:"longitude" desc:"longitude value" min:"0" max:"180"`
}

type Feature struct {
	Name     string `json:"Name" desc:"feature name" enum:"GuangZhou,Beijing,Shenzhen,Shanghai"`
	Location Point  `json:"location" desc:"location"`
}

var GetFeatureAPi = serv.NewAPI(
	"getFeature",
	`Obtains the feature at a given position`,
	&Point{},
	&Feature{},
	GetFeatureCall,
)

func GetFeatureCall(data []byte) (interface{}, *serv.APIError) {
	req := &Point{}
	err := serv.UnmarshalAndCheckValue(data, req)
	if err != nil {
		return nil, serv.NewError(err)
	}
	return &Feature{
		Name:     "Guangdong",
		Location: *req,
	}, nil
}
