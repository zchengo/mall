package initialize

import (
	"imall/common"
	"imall/global"
	"imall/models/app"
	"imall/models/web"
	"imall/service"
	"math/rand"
	"time"
)

// 定时任务
func Cron() {

	if !global.Config.Cron.Enable {
		return
	}

	ticker := time.NewTicker(3600 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var c service.AppCartService
		var o service.AppOrderService
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		openId := "oUT385ZLmRr6R_a9xKSfSW9SekYI"

		goodsId := make([]uint64, 3)
		global.Db.Select("id").Model(&web.Goods{}).Limit(3).Find(&goodsId)

		t1 := r.Intn(10) + 1
		t2 := r.Intn(3) + 1

		for i := 0; i < t1; i++ {
			// 模拟添加购物车
			for i := 0; i < t2; i++ {
				c.Add(app.CartAddParam{
					GoodsId:    uint(goodsId[r.Intn(3)]),
					GoodsCount: uint(r.Intn(5) + 1),
					OpenId:     openId,
				})
			}
			// 模拟提交订单
			o.Submit(app.OrderSubmitParam{
				OpenId: openId,
				Sid:    100001,
			})
		}

		// 模拟更新订单
		var oc int64

		orderId := make([]int, 0)
		global.Db.Select("id").Model(&web.Order{}).Find(&orderId).Count(&oc)
		status := r.Intn(4) + 2

		order := web.Order{
			Status:  status,
			Updated: common.NowTime(),
		}
		global.Db.Model(&order).Where("id IN ?", orderId[oc-3:oc]).Updates(order)
	}
}
