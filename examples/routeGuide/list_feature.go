package main

import (
	"github.com/YMhao/EasyApi/serv"
	"github.com/gin-gonic/gin"
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
	`Obtains the Features available within the given Rectangle.  Results are
	streamed rather than returned at once (e.g. in a response message with a
	repeated field), as the rectangle may cover a large area and contain a
	huge number of features.`,
	&Rectangle{},
	&FeatureList{},
	ListFeatureCall,
)

func ListFeatureCall(data []byte, c *gin.Context) (interface{}, *serv.APIError) {
	req := &Rectangle{}
	err := serv.UnmarshalAndCheckValue(data, req)
	if err != nil {
		return nil, serv.NewError(err)
	}

	return &FeatureList{
		FeatureList: []*Feature{
			&Feature{
				Name: "Name1",
				Location: Point{
					Latitude:  1,
					Longitude: 2,
				},
			},
			&Feature{
				Name: "Name2",
				Location: Point{
					Latitude:  3,
					Longitude: 2,
				},
			},
		},
	}, nil
}
