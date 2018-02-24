package serv

// CateName 是 分类名
type CateName string

// APICollect : collect all api
type APICollect interface {
	AllAPI() map[CateName][]API // 所有api，key是api类别
}
