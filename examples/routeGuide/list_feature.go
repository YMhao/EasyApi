package main

import (
	"github.com/YMhao/EasyApi/serv"
)

type Rectangle struct {
	Lo Point `json:"lo" desc:"One corner of the rectangle."`
	Hi Point `json:"hi" desc:"The other corner of the rectangle."`
}

type FeatureList struct {
	FeatureList []*Feature `json:"featureList" desc:"FeatureList"`
}

var ListFeatureAPi = serv.NewAPI(
	"listFeature",
	`list feature`,
	&Rectangle{},
	&FeatureList{},
	ListFeatureCall,
)

func ListFeatureCall(data []byte) (interface{}, *serv.APIError) {
	req := &Rectangle{}
	err := serv.UnmarshalAndCheckValue(data, req)
	if err != nil {
		return nil, serv.NewError(err)
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
