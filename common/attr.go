package common

// Attr : the attribute
type Attr struct {
	Type      string      `json:"type"`
	Desc      string      `json:"desc"`
	PkgPath   string      `json:"-"`
	IsStruct  bool        `json:"-"`
	ExtraAttr interface{} `json:"extraAttr"`
}

// IntExtraAttr : the  attribute of the int in extra
type IntExtraAttr struct {
	HasMin        bool    `json:"hasMin"`
	Min           int64   `json:"min"`
	HasMax        bool    `json:"hasMax"`
	Max           int64   `json:"max"`
	IsEnumValue   bool    `json:"isEnumValue"`
	EnumValueList []int64 `json:"enumValueList"`
}

// FloatExtraAttr : the Property of the float in extra
type FloatExtraAttr struct {
	HasMin bool    `json:"hasMin"`
	Min    float64 `json:"min"`
	HasMax bool    `json:"hasMax"`
	Max    float64 `json:"max"`
}

// StrExtraAttr : the Property of the string in extra
type StrExtraAttr struct {
	IsEnumValue   bool     `json:"isEnumValue"`
	EnumValueList []string `json:"enumValueList"`
}

// ArrayExtraAttr : the Property of the array
type ArrayExtraAttr struct {
	ItemType string `json:"itemType"`
	IsStruct bool   `json:"-"`
	PkgPath  string `json:"-"`
}
