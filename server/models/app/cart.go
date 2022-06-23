package app

// 购物车添加参数模型
type CartAddParam struct {
	GoodsId    uint   `form:"goodsId"`
	GoodsCount uint   `form:"goodsCount"`
	OpenId     string `form:"openId"`
}

// 购物车删除参数模型
type CartDeleteParam struct {
	OpenId  string `form:"openId"`
	GoodsId string `form:"goodsId"`
}

// 购物车清除参数模型
type CartClearParam struct {
	OpenId string `form:"openId"`
}

// 购物车信息查询参数模型
type CartQueryParam struct {
	OpenId string `form:"openId"`
}

// 购物车商品项传输模型
type CartItem struct {
	Id       uint64  `gorm:"primaryKey" json:"id"`
	ImageUrl string  `gorm:"image_url"  json:"imageUrl"`
	Title    string  `gorm:"title"      json:"title"`
	Price    float64 `gorm:"price"      json:"price"`
	Count    int     `gorm:"count"      json:"count"`
}

// 购物车信息传输模型
type CartInfo struct {
	CartItem   []CartItem `json:"cartItem"`
	TotalPrice float64    `json:"totalPrice"`
}
