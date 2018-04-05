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
func getObjPkgPath(obj interface{}) string {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.PkgPath()
}

func genAPIDocList(conf *APIServConf, setsOfAPIs APISets) []*common.ApiDoc {
	m := []*common.ApiDoc{}
	for nameOfset, APIs := range setsOfAPIs {
		for _, API := range APIs {
			doc := API.Doc()
			apiDoc := &common.ApiDoc{
				ApiId:   doc.ID,
				ApiDesc: doc.Descript,
				Tag:     string(nameOfset),
				Path:    "/" + doc.ID,
				Request: common.ObjInfo{
					Name: getObjName(doc.Request),
					Description: func(desc string) string {
						if desc == "" {
							return "request context"
						}
						return desc
					}(doc.RequestDescript),
					Fields:     common.NewObjDoc(doc.Request).FieldAttrMap(),
					DepObjList: common.NewObjDoc(doc.Request).ListDepObjDoc(),
					PkgPath:    getObjPkgPath(doc.Request),
				},
				Response: common.ObjInfo{
					Name: getObjName(doc.Response),
					Description: func(desc string) string {
						if desc == "" {
							return "response context"
						}
						return desc
					}(doc.ResponseDescript),
					Fields:     common.NewObjDoc(doc.Response).FieldAttrMap(),
					DepObjList: common.NewObjDoc(doc.Response).ListDepObjDoc(),
					PkgPath:    getObjPkgPath(doc.Response),
				},
				SwaggerAPIType: common.SwaggerAPITypeJson,
			}
			m = append(m, apiDoc)
		}
	}
	return m
}
