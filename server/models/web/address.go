package web

// 收货地址映射模型
type Address struct {
	Id              uint64 `gorm:"primaryKey"`       // 地址编号
	OpenId          string `gorm:"open_id"`          // 微信用户ID
	Name            string `gorm:"name"`             // 收货人姓名
	Mobile          string `gorm:"mobile"`           // 手机号
	Province        string `gorm:"province"`         // 省份
	City            string `gorm:"city"`             // 城市
	District        string `gorm:"district"`         // 区（县）
	DetailedAddress string `gorm:"detailed_address"` // 详细地址
	IsDefault       int    `gorm:"is_default"`       // 是否默认，1-默认，2-非默认
	Created         string `gorm:"created"`          // 创建时间
	Updated         string `gorm:"updated"`          // 更新时间
}
