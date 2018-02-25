package swagger

import (
	"fmt"
	"sync"
)

type _AvoidReMap struct {
	Map     map[string]string
	NameMap map[string]int
	Lock    sync.Mutex
}

// AvoidRepeatMap 避免重复，名字相同但包路径不同的struct将重命名
var AvoidRepeatMap = _AvoidReMap{
	Map:     make(map[string]string),
	NameMap: make(map[string]int),
}

func (a *_AvoidReMap) GetTypeName(pkgPath string, typeName string) string {
	a.Lock.Lock()
	defer a.Lock.Unlock()

	key := pkgPath + typeName

	name, ok := a.Map[key]
	if ok {
		return name
	}

	i, ok := a.NameMap[typeName]
	if !ok {
		a.NameMap[typeName] = 0
		a.Map[key] = typeName
		return typeName
	}

	i++
	a.NameMap[typeName] = i
	typeName = typeName + fmt.Sprintf("R%d", i)
	a.Map[key] = typeName
	return typeName
}
