package common

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ObjDoc object相关文档信息
type ObjDoc struct {
	obj     interface{}
	pkgPath string
	objType string
}

// NewObjDoc 从一个结构体对象new一个object文档
func NewObjDoc(obj interface{}) *ObjDoc {
	pkgPath := ""
	objType := ""
	if nil != obj {
		t := reflect.TypeOf(obj).Elem()
		pkgPath = t.PkgPath()
		objType = t.Name()
	}
	return &ObjDoc{
		obj:     obj,
		pkgPath: pkgPath,
		objType: objType,
	}
}

// FieldAttrMap 返回子元素map
func (o *ObjDoc) FieldAttrMap() (attrMap map[string]*Attr) {
	attrMap = make(map[string]*Attr)
	if nil == o.obj {
		return attrMap
	}
	objType := reflect.TypeOf(o.obj).Elem()
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		o.fieldDoc(attrMap, field)
	}
	return
}

func (o *ObjDoc) checkNameAndDesc(fieldName, name, desc string) {
	if strings.ContainsAny(name, "_0123456789") {
		fmt.Printf("Warn: not recommend name %s!%s#%s %s\n", name, o.pkgPath, o.objType, fieldName)
	}
	if "" == desc {
		fmt.Printf("Warn: miss desc for %s!%s#%s %s\n", name, o.pkgPath, o.objType, fieldName)
	}
}

func (o *ObjDoc) fieldDoc(attrMap map[string]*Attr, field reflect.StructField) {
	name := field.Tag.Get("json")
	desc := field.Tag.Get("desc")
	o.checkNameAndDesc(field.Name, name, desc)
	if "" == name {
		return
	}
	_, ok := attrMap[name]
	if ok {
		fmt.Printf("Error:duplicate attr name %s!%s#%s %s\n", name, o.pkgPath, o.objType, field.Name)
		os.Exit(1)
	}
	t := field.Type
	for {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		} else {
			break
		}
	}
	k := o.getTypeName(t.Kind())
	if k == "int" {
		attrMap[name] = &Attr{
			Type:      "int",
			Desc:      desc,
			ExtraAttr: o.getIntExtraAttr(field),
		}
	} else if k == "float" {
		attrMap[name] = &Attr{
			Type:      "float",
			Desc:      desc,
			ExtraAttr: o.getFloatExtraAttr(field),
		}
	} else if k == "string" {
		attrMap[name] = &Attr{
			Type:      "string",
			Desc:      desc,
			ExtraAttr: o.getStringExtraAttr(field),
		}
	} else if k == "array" {
		attrMap[name] = &Attr{
			Type:      "array",
			Desc:      desc,
			ExtraAttr: o.getArrayExtraAttr(t),
		}
	} else if k == "struct" {
		attrMap[name] = &Attr{
			Type:      t.Name(),
			Desc:      desc,
			IsStruct:  true,
			PkgPath:   t.PkgPath(),
			ExtraAttr: nil,
		}
	} else if k == "bool" {
		attrMap[name] = &Attr{
			Type:      "bool",
			Desc:      desc,
			ExtraAttr: nil,
		}
	} else {
		fmt.Printf("Error:unkwon type %s#%s.%s\n", o.pkgPath, o.objType, t.Name())
		os.Exit(1)
	}
}

func (o *ObjDoc) getIntExtraAttr(field reflect.StructField) *IntExtraAttr {
	ret := &IntExtraAttr{
		EnumValueList: []int64{},
	}
	minStr, ok := field.Tag.Lookup("min")
	if ok {
		ret.HasMin = true
		min, err := strconv.ParseInt(minStr, 10, 64)
		if nil != err {
			fmt.Printf("Error: not int type!%s#%s %s\n", o.pkgPath, o.objType, minStr)
			os.Exit(1)
		}
		ret.Min = min
	}
	maxStr, ok := field.Tag.Lookup("max")
	if ok {
		ret.HasMax = true
		max, err := strconv.ParseInt(maxStr, 10, 64)
		if nil != err {
			fmt.Printf("Error: not int type!%s#%s %s\n", o.pkgPath, o.objType, maxStr)
			os.Exit(1)
		}
		ret.Max = max
	}
	enumStr, ok := field.Tag.Lookup("enum")
	if ok {
		ret.IsEnumValue = true
		numStrList := strings.Split(enumStr, ",")
		for _, numStr := range numStrList {
			num, err := strconv.ParseInt(numStr, 10, 64)
			if nil != err {
				fmt.Printf("Error: not int list!type %s#%s %s\n", o.pkgPath, o.objType, enumStr)
				os.Exit(1)
			}
			ret.EnumValueList = append(ret.EnumValueList, num)
		}
	}
	return ret
}

func (o *ObjDoc) getTypeName(k reflect.Kind) string {
	if k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 || k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64 {
		return "int"
	} else if k == reflect.Float32 || k == reflect.Float64 {
		return "float"
	} else if k == reflect.String {
		return "string"
	} else if k == reflect.Slice {
		return "array"
	} else if k == reflect.Struct {
		return "struct"
	} else if k == reflect.Bool {
		return "bool"
	} else {
		return "unkwon"
	}
}

func (o *ObjDoc) getFloatExtraAttr(field reflect.StructField) *FloatExtraAttr {
	ret := &FloatExtraAttr{}
	minStr, ok := field.Tag.Lookup("min")
	if ok {
		ret.HasMin = true
		min, err := strconv.ParseFloat(minStr, 64)
		if nil != err {
			fmt.Printf("Error: not float type!%s#%s %s\n", o.pkgPath, o.objType, minStr)
			os.Exit(1)
		}
		ret.Min = min

	}
	maxStr, ok := field.Tag.Lookup("max")
	if ok {
		ret.HasMax = true
		max, err := strconv.ParseFloat(maxStr, 64)
		if nil != err {
			fmt.Printf("Error: not float type!%s#%s %s\n", o.pkgPath, o.objType, maxStr)
			os.Exit(1)
		}
		ret.Max = max
	}
	return ret
}

func (o *ObjDoc) getStringExtraAttr(field reflect.StructField) *StrExtraAttr {
	ret := &StrExtraAttr{
		EnumValueList: []string{},
	}
	enumStr, ok := field.Tag.Lookup("enum")
	if ok {
		ret.IsEnumValue = true
		ret.EnumValueList = strings.Split(enumStr, ",")
	}
	return ret
}

func (o *ObjDoc) getArrayExtraAttr(t reflect.Type) *ArrayExtraAttr {
	valueType := t.Elem()
	for {
		if valueType.Kind() == reflect.Ptr {
			valueType = valueType.Elem()
		} else {
			break
		}
	}
	typeName := o.getTypeName(valueType.Kind())
	isStruct := false
	pkgPath := ""
	if typeName == "array" || typeName == "map" {
		fmt.Printf("Error:contains array or map type in array is not allow!type %s#%s\n", o.pkgPath, o.objType)
		os.Exit(1)
	}
	if typeName == "struct" {
		typeName = valueType.Name()
		pkgPath = valueType.PkgPath()
		isStruct = true
	}
	return &ArrayExtraAttr{
		ItemType: typeName,
		IsStruct: isStruct,
		PkgPath:  pkgPath,
	}
}

// DepObjDoc 依赖的object文档
type DepObjDoc struct {
	Name    string           `json:"name"`
	Fields  map[string]*Attr `json:"fields"`
	PkgPath string           `json:"-"`
}

// ListDepObjDoc 列出依赖的object文档
func (o *ObjDoc) ListDepObjDoc() []DepObjDoc {
	retList := []DepObjDoc{}
	if nil == o.obj {
		return retList
	}
	typeMap := map[reflect.Type]int{}
	objType := reflect.TypeOf(o.obj).Elem()
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		o.walkType(typeMap, field.Type, 0)
	}
	maxLv := 0
	for _, lv := range typeMap {
		if lv > maxLv {
			maxLv = lv
		}
	}
	for i := maxLv; i >= 0; i-- {
		for t, lv := range typeMap {
			if lv == i {
				attrMap := map[string]*Attr{}
				for j := 0; j < t.NumField(); j++ {
					o.fieldDoc(attrMap, t.Field(j))
				}
				retList = append(retList, DepObjDoc{
					Name:    t.Name(),
					Fields:  attrMap,
					PkgPath: t.PkgPath(),
				})
			}
		}
	}
	return retList
}

func (o *ObjDoc) walkType(typeMap map[reflect.Type]int, t reflect.Type, lv int) {
	for {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		} else {
			break
		}
	}
	k := t.Kind()
	if k == reflect.Slice {
		o.walkType(typeMap, t.Elem(), lv+1)
	} else if k == reflect.Map {
		o.walkType(typeMap, t.Elem(), lv+1)
	} else if k == reflect.Struct {
		curLv, ok := typeMap[t]
		if false == ok {
			typeMap[t] = lv
		} else if ok && lv > curLv {
			typeMap[t] = lv
		}
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			o.walkType(typeMap, field.Type, lv+1)
		}
	}
}
