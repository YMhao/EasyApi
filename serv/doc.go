package serv

import (
	"github.com/YMhao/EasyApi/common"
)

func genAPIDocList(conf *APIServConf, apiColl APICollect) []*common.ApiDoc {
	m := []*common.ApiDoc{}
	allAPI := apiColl.AllAPI()
	for cateName, apiList := range allAPI {
		for _, api := range apiList {
			doc := api.Doc()
			apiDoc := &common.ApiDoc{
				ApiId:              doc.ID,
				ApiDesc:            doc.Descript,
				Tag:                string(cateName),
				Path:               "/" + conf.ServiceName + "/" + conf.Version + "/" + doc.ID,
				RequestDesc:        doc.RequestDescript,
				RequestObj:         common.NewObjDoc(doc.Request).FieldAttrMap(),
				RequestDepObjList:  common.NewObjDoc(doc.Request).ListDepObjDoc(),
				ResponseDesc:       doc.ResponseDescript,
				ResponseObj:        common.NewObjDoc(doc.Response).FieldAttrMap(),
				ResponseDepObjList: common.NewObjDoc(doc.Response).ListDepObjDoc(),
			}
			m = append(m, apiDoc)
		}
	}
	return m
}
