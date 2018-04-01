package main

import "github.com/YMhao/EasyApi/serv"

type APIColl struct {
}

func (a APIColl) AllAPI() map[serv.CateName][]serv.API {
	return map[serv.CateName][]serv.API{
		"helloServ": []serv.API{
			&GetFeatureAPi{},
			&ListFeatureAPi{},
		},
	}
}
