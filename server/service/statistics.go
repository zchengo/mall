package service

import (
	"imall/common"
	"imall/global"
	"imall/models/web"
	"time"
)

type WebStatisticsService struct {
}

// 统计当日订单数据
func (s *WebStatisticsService) OrderData() web.OrderDate {
	today := common.NowTime()
	day := today[:10] + "%"
	pps := "status = 1 and created like ?"
	pds := "status = 2 and created like ?"
	prs := "status = 3 and created like ?"
	ras := "status = 4 and created like ?"
	fds := "status = 5 and created like ?"
	orderDate := web.OrderDate{}
	global.Db.Table("order").Where(pps, day).Find(&web.Order{}).Count(&orderDate.PendPay)
	global.Db.Table("order").Where(pds, day).Find(&web.Order{}).Count(&orderDate.Canceled)
	global.Db.Table("order").Where(prs, day).Find(&web.Order{}).Count(&orderDate.PendPayed)
	global.Db.Table("order").Where(ras, day).Find(&web.Order{}).Count(&orderDate.InDelivery)
	global.Db.Table("order").Where(fds, day).Find(&web.Order{}).Count(&orderDate.Finished)
	return orderDate
}

// 统计本周数据
func (s *WebStatisticsService) WeekData() web.WeekData {
	var weekInfo web.WeekData
	switch time.Now().Weekday() {
	case time.Monday:
		weekInfo = getWeekData(1)
	case time.Tuesday:
		weekInfo = getWeekData(2)
	case time.Wednesday:
		weekInfo = getWeekData(3)
	case time.Thursday:
		weekInfo = getWeekData(4)
	case time.Friday:
		weekInfo = getWeekData(5)
	case time.Saturday:
		weekInfo = getWeekData(6)
	case time.Sunday:
		weekInfo = getWeekData(7)
	default:
	}
	return weekInfo
}

func getWeekData(days int) web.WeekData {
	var wd web.WeekData
	for i, index := days-1, 0; i >= 0; i-- {
		var result []float64
		var orderSum int64
		var amountSum float64
		nowTime := common.WeekTime(i) + "%"
		cLike := "created like ?"
		global.Db.Table("order").Where(cLike, nowTime).Find(&web.Order{}).Count(&orderSum)
		global.Db.Table("order").Where(cLike, nowTime).Pluck("total_price", &result)
		for _, v := range result {
			amountSum += v
		}
		wd.Orders[index] = orderSum
		wd.Amount[index] = amountSum
		index++
	}
	return wd
}
