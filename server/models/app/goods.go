package app

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

// 商品详情参数模型
type GoodsInfoParam struct {
	Id uint64 `form:"id"`
}

// 商品列表参数模型
type GoodsQueryParam struct {
	CategoryId uint64 `form:"categoryId"`
}

// 商品搜索参数模型
type GoodsSearchParam struct {
	KeyWord string `form:"keyWord"`
}

// 商品列表传输模型
type GoodsList struct {
	Id       uint64  `json:"id"`
	ImageUrl string  `json:"imageUrl"`
	Title    string  `json:"title"`
	Sales    uint    `json:"sales"`
	Price    float64 `json:"price"`
}

// 商品详情传输模型
type GoodsInfo struct {
	Id       uint64  `json:"id"`
	Title    string  `json:"title"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageUrl string  `json:"imageUrl"`
	Remark   string  `json:"remark"`
}
