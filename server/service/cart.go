package service

import (
	"imall/global"
	"imall/models/app"
	"strconv"
	"strings"
)

type AppCartService struct {
}

// 添加购物车
func (c *AppCartService) Add(param app.CartAddParam) int64 {
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	goodsId := strconv.Itoa(int(param.GoodsId))
	return global.Rdb.HSet(ctx, key, goodsId, param.GoodsCount).Val()
}

// 删除购物车中的某项商品
func (c *AppCartService) Delete(param app.CartDeleteParam) int64 {
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	return global.Rdb.HDel(ctx, key, param.GoodsId).Val()
}

// 清空购物车中的商品
func (c *AppCartService) Clear(param app.CartClearParam) int64 {
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	pidsAndCounts := global.Rdb.HGetAll(ctx, key).Val()
	var rows int64
	for id, _ := range pidsAndCounts {
		rows += global.Rdb.HDel(ctx, key, id).Val()
	}
	return rows
}

// 获取购物车信息
func (c *AppCartService) GetInfo(param app.CartQueryParam) app.CartInfo {
	var cartInfo app.CartInfo
	// 从缓存中获取加购商品与数量
	key := strings.Join([]string{"user", param.OpenId, "cart"}, ":")
	goodsIdsAndCounts := global.Rdb.HGetAll(ctx, key).Val()

	// 遍历、关联加购商品与数量
	goodsIds := make([]uint, 0)
	idsAndCounts := make(map[uint64]int, 0)
	for gid, gcount := range goodsIdsAndCounts {
		id, _ := strconv.Atoi(gid)
		count, _ := strconv.Atoi(gcount)
		goodsIds = append(goodsIds, uint(id))
		idsAndCounts[uint64(id)] = count
	}

	// 获取商品信息，并计算全部商品的价格
	if len(goodsIds) > 0 {
		global.Db.Table("goods").Find(&cartInfo.CartItem, goodsIds)
		for i, item := range cartInfo.CartItem {
			cartInfo.CartItem[i].Count = idsAndCounts[item.Id]
			cartInfo.TotalPrice = cartInfo.TotalPrice + item.Price*float64(idsAndCounts[item.Id])
		}
		return cartInfo
	}
	return cartInfo
}
