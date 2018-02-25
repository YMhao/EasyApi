package swagger

import (
	"fmt"

	"github.com/YMhao/EasyApi/common"
	"github.com/go-openapi/spec"
)

var typeMap = map[string]string{
	"int":     "integer",
	"float":   "number",
	"string":  "string",
	"array":   "array",
	"struct":  "object",
	"bool":    "boolean",
	"unknown": "unknown",
}

func getSwaggerType(tp string) (t string, ref bool) {
	t, ok := typeMap[tp]
	if !ok {
		return tp, true
	}
	return t, false
}

func _getSchame(t, desc string, ref bool) *spec.Schema {
	if t == "array" {
		panic("In swagger, type " + t + " is not supported")
	}
	if ref {
		return &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Ref: spec.MustCreateRef("#/definitions/" + t),
			},
		}
	}
	return &spec.Schema{
		SchemaProps: spec.SchemaProps{
			Description: desc,
			Type: spec.StringOrArray{
				t,
			},
		},
	}
}

func itemToSchame(item *common.Attr) *spec.Schema {
	t, ref := getSwaggerType(item.Type)
	if ref {
		if item.IsStruct {
			t = AvoidRepeatMap.GetTypeName(item.PkgPath, t)
		}
		return &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: item.Desc,
				Ref:         spec.MustCreateRef("#/definitions/" + t),
			},
		}
	}
	switch t {
	case "array":
		attr := item.ExtraAttr.(*common.ArrayExtraAttr)
		itemType, ref := getSwaggerType(attr.ItemType)
		if ref {
			if attr.IsStruct {
				itemType = AvoidRepeatMap.GetTypeName(attr.PkgPath, itemType)
			}
		}
		s := &spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: item.Desc,
				Type: spec.StringOrArray{
					t,
				},
				Items: &spec.SchemaOrArray{
					Schema: _getSchame(itemType, item.Desc, ref),
				},
			},
		}
		return s
	case "string":
		attr := item.ExtraAttr.(*common.StrExtraAttr)
		strPro := spec.StringProperty()
		strPro.Description = item.Desc
		if attr.IsEnumValue {
			for _, enum := range attr.EnumValueList {
				strPro.Enum = append(strPro.Enum, enum)
			}
		}
		return strPro
	case "integer":
		attr := item.ExtraAttr.(*common.IntExtraAttr)
		intPro := spec.Int64Property()
		intPro.Description = item.Desc
		if attr.IsEnumValue {
			for _, enum := range attr.EnumValueList {
				intPro.Enum = append(intPro.Enum, enum)
			}
		}
		if attr.HasMax {
			max := float64(attr.Max)
			fmt.Println("max:", max)
			intPro.Maximum = &max
		}
		if attr.HasMin {
			min := float64(attr.Min)
			fmt.Println("min:", min)
			intPro.Minimum = &min
		}
		return intPro
	case "unknown":
		fmt.Println("Warn: In swagger, type unknown  is not supported")
		return &spec.Schema{}
	case "boolean":
		return _getSchame(t, item.Desc, ref)
	case "number":
		return _getSchame(t, item.Desc, ref)
	default:
		fmt.Println("item: ", *item)
		if item.IsStruct {
			t = AvoidRepeatMap.GetTypeName(item.PkgPath, t)
		}
		return _getSchame(t, item.Desc, ref)
	}
}
