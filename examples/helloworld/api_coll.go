package main

// APIColl api集合
import "github.com/YMhao/EasyApi/serv"

type APIColl struct {
}

// AllAPI 列出所有api
func (a APIColl) AllAPI() map[serv.CateName][]serv.API {
	return map[serv.CateName][]serv.API{
		"helloServ": []serv.API{
			&HelloAPI{},
		},
	}
}
