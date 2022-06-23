package service

import (
	"fmt"
	"imall/common"
	"imall/global"
	"imall/models/app"
	"imall/models/web"
	"strconv"
	"strings"
)

type WebOrderService struct {
}

type AppOrderService struct {
}

// 删除订单
func (o *WebOrderService) Delete(param web.OrderDeleteParam) int64 {
	return global.Db.Delete(&web.Order{}, param.Id).RowsAffected
}

// 更新订单
func (o *WebOrderService) Update(param web.OrderUpdateParam) int64 {
	order := web.Order{
		Id:      param.Id,
		Status:  param.Status,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&order).Updates(order).RowsAffected
}

// 获取订单列表
func (o *WebOrderService) GetList(param web.OrderListParam) ([]web.OrderList, int64) {
	orders := make([]web.Order, 0)
	query := &web.Order{
		Id:     param.Id,
		Status: param.Status,
	}
	rows := common.RestPage(param.Page, "order", query, &orders, &[]web.Order{})
	orderList := make([]web.OrderList, 0)
	var user web.User
	for _, o := range orders {
		global.Db.Where("open_id = ?", o.OpenId).First(&user)
		order := web.OrderList{
			Id:         o.Id,
			Avatar:     user.Avatar,
			NickName:   user.NickName,
			TotalPrice: o.TotalPrice,
			GoodsCount: o.GoodsCount,
			Status:     o.Status,
			Created:    o.Created,
		}
		orderList = append(orderList, order)
	}
	return orderList, rows
}

// 获取订单详情
func (o *WebOrderService) GetDetail(param web.OrderDetailParam) (od web.OrderDetail) {
	var order web.Order
	var address web.Address

	// 查询订单信息与地址信息
	global.Db.First(&order, param.Id)
	global.Db.First(&address, order.AddressId)

	// 查询商品信息
	goodsIds := make([]uint, 0)
	goodsIdsCount := map[int64]int{}
	goodList := make([]web.Goods, 0)
	if order.GoodsIdsCount != "" {
		for _, gidCount := range strings.Split(order.GoodsIdsCount, ",") {
			ic := strings.Split(gidCount, ":")
			id, _ := strconv.Atoi(ic[0])
			count, _ := strconv.Atoi(ic[1])
			goodsIdsCount[int64(id)] = count
			goodsIds = append(goodsIds, uint(id))
		}
		global.Db.Table("goods").Find(&goodList, goodsIds)
	}

	// 组装商品项
	goodsItem := make([]web.GoodsItem, 0)
	for _, g := range goodList {
		gItem := web.GoodsItem{
			Id:       g.Id,
			Title:    g.Title,
			Price:    g.Price,
			ImageUrl: g.ImageUrl,
			Count:    goodsIdsCount[int64(g.Id)],
		}
		goodsItem = append(goodsItem, gItem)
	}
	orderDetail := web.OrderDetail{
		Id:              order.Id,
		Status:          order.Status,
		TotalPrice:      order.TotalPrice,
		GoodsItem:       goodsItem,
		Name:            address.Name,
		Mobile:          address.Mobile,
		Province:        address.Province,
		City:            address.City,
		District:        address.District,
		DetailedAddress: address.DetailedAddress,
		Created:         order.Created,
	}
	return orderDetail
}

// 更新订单
func (o *AppOrderService) Update(param app.OrderUpdateParam) int64 {
	order := web.Order{
		Id:      param.Id,
		Status:  param.Status,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&order).Updates(order).RowsAffected
}

// 提交订单
func (o *AppOrderService) Submit(param app.OrderSubmitParam) int64 {
	// 获取购物车信息
	var cartInfo app.CartInfo
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	goodsIdsAndCounts := global.Rdb.HGetAll(ctx, key).Val()
	goodsIds := make([]uint, 0)
	idsAndCounts := make(map[uint64]int, 0)
	for gid, gcount := range goodsIdsAndCounts {
		id, _ := strconv.Atoi(gid)
		count, _ := strconv.Atoi(gcount)
		goodsIds = append(goodsIds, uint(id))
		idsAndCounts[uint64(id)] = count
	}
	var goodsCount uint
	if len(goodsIds) > 0 {
		global.Db.Table("goods").Find(&cartInfo.CartItem, goodsIds)
		for index, item := range cartInfo.CartItem {
			cartInfo.CartItem[index].Count = idsAndCounts[item.Id]
			goodsCount = goodsCount + uint(cartInfo.CartItem[index].Count)
			cartInfo.TotalPrice = cartInfo.TotalPrice + item.Price*float64(idsAndCounts[item.Id])
		}
	}
	// 获取地址编号
	var addressId uint64
	global.Db.Table("address").Select("id").
		Where("is_default = ? and open_id = ?", 1, param.OpenId).Take(&addressId)

	goods := make([]string, 0)
	for _, item := range cartInfo.CartItem {
		idAndCount := fmt.Sprintf("%s:%d", strconv.Itoa(int(item.Id)), idsAndCounts[item.Id])
		goods = append(goods, idAndCount)
	}
	order := app.Order{
		GoodsIdsCount: strings.Join(goods, ","),
		TotalPrice:    cartInfo.TotalPrice,
		GoodsCount:    goodsCount,
		Status:        1,
		AddressId:     addressId,
		OpenId:        param.OpenId,
		Created:       common.NowTime(),
	}
	rowsAffected := global.Db.Create(&order).RowsAffected
	if rowsAffected > 0 {
		// 清空购物车信息
		goodsIdsAndCounts := global.Rdb.HGetAll(ctx, key).Val()
		for id, _ := range goodsIdsAndCounts {
			global.Rdb.HDel(ctx, key, id).Val()
		}
	}
	return rowsAffected
}

// 获取订单列表
func (o *AppOrderService) GetList(param app.OrderQueryParam) []app.OrderList {
	// 根据订单状态查询订单
	orders := make([]app.Order, 0)
	if param.Type == 1 {
		global.Db.Table("order").Where("status != ?", 5).Find(&orders)
	} else {
		global.Db.Table("order").Where("status = ?", 5).Find(&orders)
	}

	// 组装订单列表
	orderList := make([]app.OrderList, 0)
	for _, o := range orders {
		var order app.OrderList
		goods := make([]app.Goods, 0)
		order.Id = o.Id
		order.Status = o.Status
		order.TotalPrice = o.TotalPrice

		// 查询商品信息
		goodsIds := make([]uint, 0)
		goodsIdsCount := map[int64]int{}
		for _, gidCount := range strings.Split(o.GoodsIdsCount, ",") {
			ic := strings.Split(gidCount, ":")
			id, _ := strconv.Atoi(ic[0])
			count, _ := strconv.Atoi(ic[1])
			goodsIdsCount[int64(id)] = count
			goodsIds = append(goodsIds, uint(id))
		}
		global.Db.Table("goods").Find(&goods, goodsIds)

		// 组装商品项
		goodsItem := make([]app.GoodsItem, 0)
		for _, g := range goods {
			gItem := app.GoodsItem{
				Id:       g.Id,
				Title:    g.Title,
				Price:    g.Price,
				ImageUrl: g.ImageUrl,
				Count:    goodsIdsCount[int64(g.Id)],
			}
			order.GoodsCount = order.GoodsCount + uint(gItem.Count)
			goodsItem = append(goodsItem, gItem)
		}
		order.GoodsItem = goodsItem
		orderList = append(orderList, order)
	}
	return orderList
}
