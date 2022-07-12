package app

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

// 收货地址保存参数模型
type AddressSaveParam struct {
	Id              uint64 `form:"id"`
	OpenId          string `form:"openId" binding:"required"`
	Name            string `form:"name" binding:"required"`
	Mobile          string `form:"mobile" binding:"required"`
	Province        string `form:"province" binding:"required"`
	City            string `form:"city" binding:"required"`
	District        string `form:"district" binding:"required"`
	DetailedAddress string `form:"detailedAddress" binding:"required"`
	IsDefault       int    `form:"isDefault" binding:"required,gt=0"`
}

// 收货地址删除参数模型
type AddressDeleteParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

// 收货地址列表查询参数模型
type AddressListParam struct {
	OpenId string `form:"openId" json:"openId"`
}

// 收货地址信息查询参数模型
type AddressInfoParam struct {
	Id uint64 `form:"id"`
}

// 收货地址列表传输模型
type AddressList struct {
	Id              uint64 `json:"id"`
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DetailedAddress string `json:"detailedAddress"`
	IsDefault       int    `json:"isDefault"`
}

// 收货地址信息传输模型
type AddressInfo struct {
	Id              uint   `json:"id"`
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DetailedAddress string `json:"detailedAddress"`
	IsDefault       int    `json:"isDefault"`
}
