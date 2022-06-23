package service

import (
	"imall/common"
	"imall/global"
	"imall/models/app"
	"imall/models/web"
	"strconv"
	"strings"
)

type WebMarketService struct {
}

type AppMarketService struct {
}

// 创建营销活动
func (m *WebMarketService) Create(param web.MarketCreateParam) int64 {
	market := web.Market{
		Name:        param.Name,
		Type:        param.Type,
		BannerImage: param.BannerImage,
		BeginTime:   param.BeginTime,
		OverTime:    param.OverTime,
		GoodsIds:    param.GoodsIds,
		Created:     common.NowTime(),
	}
	return global.Db.Create(&market).RowsAffected
}

// 删除营销活动
func (m *WebMarketService) Delete(param web.MarketDeleteParam) int64 {
	return global.Db.Delete(&web.Market{}, param.Id).RowsAffected
}

// 更新营销活动
func (m *WebMarketService) Update(param web.MarketUpdateParam) int64 {
	market := web.Market{
		Id:          param.Id,
		Name:        param.Name,
		Type:        param.Type,
		BannerImage: param.BannerImage,
		BeginTime:   param.BeginTime,
		OverTime:    param.OverTime,
		GoodsIds:    param.GoodsIds,
		Status:      param.Status,
		Updated:     common.NowTime(),
	}
	return global.Db.Model(&market).Updates(market).RowsAffected
}

// 开启或关闭营销活动
func (m *WebMarketService) UpdateStatus(param web.MarketStatusUpdateParam) int64 {
	market := web.Market{
		Id:      param.Id,
		Status:  param.Status,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&market).Updates(market).RowsAffected
}

// 获取营销活动列表
func (m *WebMarketService) GetList(param web.MarketQueryParam) ([]web.MarketList, int64) {
	query := &web.Market{
		Id:     param.Id,
		Type:   param.Type,
		Status: param.Status,
	}
	marketList := make([]web.MarketList, 0)
	rows := common.RestPage(param.Page, "market", query, &marketList, &[]web.Market{})
	return marketList, rows
}

// 获取可关联活动的商品列表
func (m *WebMarketService) GetGoods(param web.MarketGoodsQueryParma) []web.MarketGoods {
	marketGoods := make([]web.MarketGoods, 0)
	goodsIds := strings.Split(param.GoodsIds, ",")
	if param.GoodsIds != "" {
		ids := make([]uint64, 0)
		for _, id := range goodsIds {
			i, _ := strconv.Atoi(id)
			ids = append(ids, uint64(i))
		}
		global.Db.Table("goods").Where("status = 1").Find(&marketGoods, ids)
	} else {
		global.Db.Table("goods").Where("status = 1").Find(&marketGoods)
	}
	return marketGoods
}

// 获取钻石展位列表
func (m *AppMarketService) GetBannerList() []app.BannerList {
	bannerList := make([]app.BannerList, 0)
	global.Db.Table("market").Where("status = 1").Find(&bannerList)
	return bannerList
}
