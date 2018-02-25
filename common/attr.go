package common

// Attr 是元素的属性值
type Attr struct {
	Type      string      `json:"type"`
	Desc      string      `json:"desc"`
	PkgPath   string      `json:"-"`
	IsStruct  bool        `json:"-"`
	ExtraAttr interface{} `json:"extraAttr"`
}

// IntExtraAttr 整型扩展属性值
type IntExtraAttr struct {
	HasMin        bool    `json:"hasMin"`
	Min           int64   `json:"min"`
	HasMax        bool    `json:"hasMax"`
	Max           int64   `json:"max"`
	IsEnumValue   bool    `json:"isEnumValue"`
	EnumValueList []int64 `json:"enumValueList"`
}

// FloatExtraAttr 浮点扩展属性值
type FloatExtraAttr struct {
	HasMin bool    `json:"hasMin"`
	Min    float64 `json:"min"`
	HasMax bool    `json:"hasMax"`
	Max    float64 `json:"max"`
}

// StrExtraAttr 字符串扩展属性值
type StrExtraAttr struct {
	IsEnumValue   bool     `json:"isEnumValue"`
	EnumValueList []string `json:"enumValueList"`
}

// ArrayExtraAttr 数组扩展属性值
type ArrayExtraAttr struct {
	ItemType string `json:"itemType"`
	IsStruct bool   `json:"-"`
	PkgPath  string `json:"-"`
}
