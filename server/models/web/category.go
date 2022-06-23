package web

// 商品类目映射模型
type Category struct {
	Id       uint64 `gorm:"primaryKey"` // 类目编号
	Name     string `gorm:"name"`       // 类目名称
	ParentId uint64 `gorm:"parent_id"`  // 父级编号
	Level    uint   `gorm:"level"`      // 类目级别
	Sort     uint   `gorm:"sort"`       // 类目排序
	Created  string `gorm:"created"`    // 创建时间
	Updated  string `gorm:"updated"`    // 更新时间
}

// 商品类目创建参数模型
type CategoryCreateParam struct {
	Name     string `json:"name"`
	ParentId uint64 `json:"parentId"`
	Level    uint   `json:"level"`
	Sort     uint   `json:"sort"`
}

// 商品类目删除参数模型
type CategoryDeleteParam struct {
	Id uint64 `form:"id"`
}

// 商品类目更新参数模型
type CategoryUpdateParam struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Sort uint   `json:"sort"`
}

// 商品类目查询参数模型
type CategoryQueryParam struct {
	Name     string `form:"name"`
	ParentId uint64 `form:"parentId"`
}

// 商品类目列表传输模型
type CategoryList struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	ParentId uint64 `json:"parentId"`
	Level    uint   `json:"level"`
	Sort     uint   `json:"sort"`
}

// 商品类目选项传输模型
type CategoryOption struct {
	Value    uint64           `json:"value"`
	Label    string           `json:"label"`
	Children []CategoryOption `json:"children"`
}
