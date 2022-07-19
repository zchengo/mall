package web

// 统计数据查询参数模型
type DataParam struct {
	Sid uint64 `form:"sid" binding:"required,gt=0"`
}

// 统计当日数据传输模型
type TodayDate struct {
	PendPay    int64   `json:"pendPay"`    // 待付款
	Payed      int64   `json:"payed"`      // 已付款
	InDelivery int64   `json:"inDelivery"` // 配送中
	Canceled   int64   `json:"canceled"`   // 已取消
	Finished   int64   `json:"finished"`   // 已完成
	PayAmount  float64 `json:"payAmount"`  // 支付金额
}

// 统计订单数据传输模型
type OrderData struct {
	Orders [17]int64 `json:"orders"`
}

// 统计店铺数据传输模型
type ShopData struct {
	Amount [7]float64 `json:"amount"`
}
