package service

import (
	"mall/common"
	"mall/global"
	"mall/models/web"
	"strconv"
	"time"
)

type WebStatisticsService struct {
}

// 统计当日数据
func (s *WebStatisticsService) TodayData(param web.DataParam) web.TodayDate {
	today := common.NowTime()
	day := today[:10] + "%"
	pps := "sid = ? and status = 1 and created like ?"
	pds := "sid = ? and status = 3 and created like ?"
	prs := "sid = ? and status = 4 and created like ?"
	ras := "sid = ? and status = 2 and created like ?"
	fds := "sid = ? and status = 5 and created like ?"
	pas := "sid = ? and created like ?"
	orderDate := web.TodayDate{}
	global.Db.Table("order").Where(pps, param.Sid, day).Find(&web.Order{}).Count(&orderDate.PendPay)
	global.Db.Table("order").Where(prs, param.Sid, day).Find(&web.Order{}).Count(&orderDate.Payed)
	global.Db.Table("order").Where(ras, param.Sid, day).Find(&web.Order{}).Count(&orderDate.InDelivery)
	global.Db.Table("order").Where(pds, param.Sid, day).Find(&web.Order{}).Count(&orderDate.Canceled)
	global.Db.Table("order").Where(fds, param.Sid, day).Find(&web.Order{}).Count(&orderDate.Finished)
	global.Db.Table("order").Where(pas, param.Sid, day).Pluck("sum(total_price) as pay_amount", &orderDate.PayAmount)
	return orderDate
}

// 统计订单数据
func (s *WebStatisticsService) OrderData(param web.DataParam) web.OrderData {
	var orders web.OrderData
	today := common.NowTime()
	for i, hour := 0, 6; i < 17; i++ {
		var hs string
		if hour < 10 {
			hs = "0" + strconv.Itoa(hour)
		} else {
			hs = strconv.Itoa(hour)
		}
		cond := today[:10] + " " + hs + "%"
		global.Db.Table("order").Where("sid = ? and created like ?", param.Sid, cond).Count(&orders.Orders[i])
		hour = hour + 1
	}
	return orders
}

// 统计店铺数据
func (s *WebStatisticsService) ShopData(param web.DataParam) web.ShopData {
	var weekInfo web.ShopData
	switch time.Now().Weekday() {
	case time.Monday:
		weekInfo = getWeekData(1, param.Sid)
	case time.Tuesday:
		weekInfo = getWeekData(2, param.Sid)
	case time.Wednesday:
		weekInfo = getWeekData(3, param.Sid)
	case time.Thursday:
		weekInfo = getWeekData(4, param.Sid)
	case time.Friday:
		weekInfo = getWeekData(5, param.Sid)
	case time.Saturday:
		weekInfo = getWeekData(6, param.Sid)
	case time.Sunday:
		weekInfo = getWeekData(7, param.Sid)
	default:
	}
	return weekInfo
}

func getWeekData(days int, sid uint64) web.ShopData {
	var wd web.ShopData
	for i, index := days-1, 0; i >= 0; i-- {
		var result []float64
		var amountSum float64
		nowTime := common.WeekTime(i) + "%"
		cLike := "sid = ? and created like ?"
		global.Db.Table("order").Where(cLike, sid, nowTime).Pluck("total_price", &result)
		for _, v := range result {
			amountSum += v
		}
		wd.Amount[index] = amountSum
		index++
	}
	return wd
}
