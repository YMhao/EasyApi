package main

import "github.com/YMhao/EasyApi/serv"

// APIColl api集合
type APIColl struct {
}

// AllAPI 列出所有api
func (a APIColl) AllAPI() map[serv.CateName][]serv.API {
	return map[serv.CateName][]serv.API{
		"helloServ": []serv.API{
			&GetFeatureAPi{},
			&ListFeatureAPi{},
		},
	}
}
