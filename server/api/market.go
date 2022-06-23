package api

import (
	"github.com/gin-gonic/gin"
	"imall/constant"
	"imall/models/web"
	"imall/response"
	"imall/service"
)

type WebMarket struct {
	service.WebMarketService
}

type AppMarket struct {
	service.AppMarketService
}

func GetWebMarket() *WebMarket {
	return &WebMarket{}
}

func GetAppMarket() *AppMarket {
	return &AppMarket{}
}

func (m *WebMarket) CreateMarket(context *gin.Context) {
	var param web.MarketCreateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := m.Create(param); count > 0 {
		response.Success(constant.Created, count, context)
		return
	}
	response.Failed(constant.NotCreated, context)
}

func (m *WebMarket) DeleteMarket(context *gin.Context) {
	var param web.MarketDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := m.Delete(param); count > 0 {
		response.Success(constant.Deleted, count, context)
		return
	}
	response.Failed(constant.NotDeleted, context)
}

func (m *WebMarket) UpdateMarket(context *gin.Context) {
	var param web.MarketUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := m.Update(param); count > 0 {
		response.Success(constant.Updated, count, context)
		return
	}
	response.Failed(constant.NotUpdated, context)
}

func (m *WebMarket) UpdateMarketStatus(context *gin.Context) {
	var param web.MarketStatusUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := m.UpdateStatus(param); count > 0 {
		response.Success(constant.Updated, count, context)
		return
	}
	response.Failed(constant.NotUpdated, context)
}

func (m *WebMarket) GetMarketList(context *gin.Context) {
	var param web.MarketQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	marketList, rows := m.GetList(param)
	response.SuccessPage(constant.Selected, marketList, rows, context)
}

func (m *WebMarket) GetMarketGoods(context *gin.Context) {
	var param web.MarketGoodsQueryParma
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	marketGoods := m.GetGoods(param)
	response.Success(constant.Selected, marketGoods, context)
}

func (m *AppMarket) GetBanners(context *gin.Context) {
	bannerList := m.GetBannerList()
	response.Success(constant.Selected, bannerList, context)
}
