package serv

import (
	"reflect"

	"github.com/YMhao/EasyApi/common"
)

func getObjName(obj interface{}) string {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

func genAPIDocList(conf *APIServConf, apiColl APICollect) []*common.ApiDoc {
	m := []*common.ApiDoc{}
	allAPI := apiColl.AllAPI()
	for cateName, apiList := range allAPI {
		for _, api := range apiList {
			doc := api.Doc()
			apiDoc := &common.ApiDoc{
				ApiId:   doc.ID,
				ApiDesc: doc.Descript,
				Tag:     string(cateName),
				Path:    "/" + conf.ServiceName + "/" + conf.Version + "/" + doc.ID,
				Request: common.ObjInfo{
					Name: getObjName(doc.Request),
					Description: func(desc string) string {
						if desc == "" {
							return "请求参数"
						}
						return desc
					}(doc.RequestDescript),
					Fields:     common.NewObjDoc(doc.Request).FieldAttrMap(),
					DepObjList: common.NewObjDoc(doc.Request).ListDepObjDoc(),
				},
				Response: common.ObjInfo{
					Name: getObjName(doc.Response),
					Description: func(desc string) string {
						if desc == "" {
							return "响应参数"
						}
						return desc
					}(doc.ResponseDescript),
					Fields:     common.NewObjDoc(doc.Response).FieldAttrMap(),
					DepObjList: common.NewObjDoc(doc.Response).ListDepObjDoc(),
				},
				SwaggerAPIType: common.SwaggerAPITypeJson,
			}
			m = append(m, apiDoc)
		}
	}
	return m
}
