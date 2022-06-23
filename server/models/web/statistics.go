package web

// 统计当日订单数据
type OrderDate struct {
	PendPay    int64 `json:"pendPay"`    // 待付款
	Canceled   int64 `json:"canceled"`   // 已取消
	PendPayed  int64 `json:"pendPayed"`  // 已付款
	InDelivery int64 `json:"inDelivery"` // 配送中
	Finished   int64 `json:"finished"`   // 已完成
}

// 统计本周数据
type WeekData struct {
	Orders [7]int64   `json:"orders"`
	Amount [7]float64 `json:"amount"`
}
