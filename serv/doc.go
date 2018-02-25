package serv

import (
	"fmt"
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

			fmt.Println("id:", reflect.TypeOf(doc.Request).Elem().Name())
			apiDoc := &common.ApiDoc{
				ApiId:   doc.ID,
				ApiDesc: doc.Descript,
				Tag:     string(cateName),
				Path:    "/" + conf.ServiceName + "/" + conf.Version + "/" + doc.ID,
				Request: common.ObjInfo{
					Name:        getObjName(doc.Request),
					Description: doc.RequestDescript,
					Fields:      common.NewObjDoc(doc.Request).FieldAttrMap(),
					DepObjList:  common.NewObjDoc(doc.Request).ListDepObjDoc(),
				},
				Response: common.ObjInfo{
					Name:        getObjName(doc.Response),
					Description: doc.ResponseDescript,
					Fields:      common.NewObjDoc(doc.Response).FieldAttrMap(),
					DepObjList:  common.NewObjDoc(doc.Response).ListDepObjDoc(),
				},
				SwaggerAPIType: common.SwaggerAPITypeJson,
			}
			m = append(m, apiDoc)
		}
	}
	return m
}
