package app

// 营销数据映射模型
type Market struct {
	Id          uint64 `gorm:"primaryKey"`   // 编号
	Name        string `gorm:"name"`         // 名称
	Type        int    `gorm:"type"`         // 类型
	BannerImage string `json:"banner_image"` // 活动图片
	BeginTime   string `gorm:"begin_time"`   // 开始时间
	OverTime    string `gorm:"over_time"`    // 结束时间
	GoodsIds    string `gorm:"goods_ids"`    // 关联商品
	Status      int    `gorm:"status"`       // 状态，1-开启，2-关闭
	Created     string `gorm:"created"`      // 创建时间
	Updated     string `gorm:"updated"`      // 更新时间
}

// 钻展列表传输模型
type BannerList struct {
	BannerImage string `json:"bannerImage"`
	GoodsIds    string `json:"goodsIds"`
}