package app

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

// 商品类目选项传输模型
type CategoryOption struct {
	Id   uint64 `json:"id"`
	Text string `json:"text"`
}
