package web

import "imall/models"

// 商品映射模型
type Goods struct {
	Id         uint64  `gorm:"id"`          // 商品编号
	CategoryId uint64  `gorm:"category_id"` // 类目编号
	Title      string  `gorm:"title"`       // 商品标题
	Name       string  `gorm:"name"`        // 商品名称
	Price      float64 `gorm:"price"`       // 商品价格
	Quantity   uint    `gorm:"quantity"`    // 商品数量
	ImageUrl   string  `gorm:"image_url"`   // 商品图片
	Remark     string  `gorm:"remark"`      // 商品备注
	Sales      uint    `gorm:"sales"`       // 商品销量
	Status     uint    `gorm:"status"`      // 商品状态，1-出售中，2-仓库中
	Created    string  `gorm:"created"`     // 创建时间
	Updated    string  `gorm:"updated"`     // 更新时间
}

// 商品创建参数模型
type GoodsCreateParam struct {
	CategoryId uint64  `json:"categoryId"`
	Title      string  `json:"title"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   uint    `json:"quantity"`
	ImageUrl   string  `json:"imageUrl"`
	Remark     string  `json:"remark"`
}

// 商品删除参数模型
type GoodsDeleteParam struct {
	Id uint64 `form:"id"`
}

// 商品更新参数模型
type GoodsUpdateParam struct {
	Id         uint64  `json:"id"`
	CategoryId uint64  `json:"categoryId"`
	Title      string  `json:"title"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   uint    `json:"quantity"`
	ImageUrl   string  `json:"imageUrl"`
	Remark     string  `json:"remark"`
}

// 商品状态更新参数模型
type GoodsStatusUpdateParam struct {
	Id     uint64 `json:"id"`
	Status uint   `json:"status"`
}

// 商品列表查询参数模型
type GoodsListParam struct {
	Page       models.Page
	Id         uint64 `form:"id"`
	CategoryId uint64 `form:"categoryId"`
	Title      string `form:"title"`
	Status     uint   `form:"status"`
}

// 商品列表传输模型
type GoodsList struct {
	Id         uint64  `json:"id"`
	CategoryId uint64  `json:"categoryId"`
	Title      string  `json:"title"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   uint    `json:"quantity"`
	ImageUrl   string  `json:"imageUrl"`
	Remark     string  `json:"remark"`
	Sales      uint    `json:"sales"`
	Status     uint    `json:"status"`
	Created    string  `json:"created"`
}
