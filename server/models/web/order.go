package web

import "imall/models"

// 数据库，订单数据映射模型
type Order struct {
	Id            uint64  `gorm:"primaryKey"`
	OpenId        string  `gorm:"open_id"`
	GoodsIdsCount string  `gorm:"goods_ids_count"`
	GoodsCount    uint    `gorm:"goods_count"`
	TotalPrice    float64 `gorm:"total_price"`
	AddressId     uint64  `gorm:"address_id"`
	Status        int     `gorm:"status"` // 订单状态，1-待付款，2-已取消，3-已付款，4-配送中，5-已完成
	Created       string  `gorm:"created"`
	Updated       string  `gorm:"updated"`
}

// 订单删除参数模型
type OrderDeleteParam struct {
	Id uint64 `form:"id"`
}

// 订单更新参数模型
type OrderUpdateParam struct {
	Id     uint64 `json:"id"`
	Status int    `json:"status"`
}

// 订单列表查询参数模型
type OrderListParam struct {
	Page   models.Page
	Id     uint64 `form:"id"`
	Status int    `form:"status"`
}

// 订单列表查询参数模型
type OrderDetailParam struct {
	Id uint64 `form:"id"`
}

// 订单列表传输模型
type OrderList struct {
	Id         uint64  `json:"id"`
	Avatar     string  `json:"avatar"`
	NickName   string  `json:"nickName"`
	TotalPrice float64 `json:"totalPrice"`
	GoodsCount uint    `json:"goodsCount"`
	Status     int     `json:"status"`
	Created    string  `json:"created"`
}

// 订单详情传输模型
type OrderDetail struct {
	Id              uint64      `json:"id"`
	Status          int         `json:"status"`
	TotalPrice      float64     `json:"totalPrice"`
	GoodsItem       []GoodsItem `json:"goodsItem"`
	Name            string      `json:"name"`
	Mobile          string      `json:"mobile"`
	Province        string      `json:"province"`
	City            string      `json:"city"`
	District        string      `json:"district"`
	DetailedAddress string      `json:"detailedAddress"`
	Created         string      `json:"created"`
}

// 订单商品项传输模型
type GoodsItem struct {
	Id       uint64  `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	ImageUrl string  `json:"imageUrl"`
	Count    int     `json:"count"`
}
