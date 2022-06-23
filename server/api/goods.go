package api

import (
	"github.com/gin-gonic/gin"
	"imall/constant"
	"imall/models/app"
	"imall/models/web"
	"imall/response"
	"imall/service"
)

type WebGoods struct {
	service.WebGoodsService
}

type AppGoods struct {
	service.AppGoodsService
}

func GetWebGoods() *WebGoods {
	return &WebGoods{}
}

func GetAppGoods() *AppGoods {
	return &AppGoods{}
}

func (g *WebGoods) CreateGoods(c *gin.Context) {
	var param web.GoodsCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if count := g.Create(param); count > 0 {
		response.Success(constant.Created, count, c)
		return
	}
	response.Failed(constant.NotCreated, c)
}

func (g *WebGoods) DeleteGoods(c *gin.Context) {
	var param web.GoodsDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if count := g.Delete(param); count > 0 {
		response.Success(constant.Deleted, count, c)
		return
	}
	response.Failed(constant.NotDeleted, c)
}

func (g *WebGoods) UpdateGoods(c *gin.Context) {
	var param web.GoodsUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if count := g.Update(param); count > 0 {
		response.Success(constant.Updated, count, c)
		return
	}
	response.Failed(constant.NotUpdated, c)
}

func (g *WebGoods) UpdateGoodsStatus(c *gin.Context) {
	var param web.GoodsStatusUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if count := g.UpdateStatus(param); count > 0 {
		response.Success(constant.Updated, count, c)
		return
	}
	response.Failed(constant.NotUpdated, c)
}

func (g *WebGoods) GetGoodsList(c *gin.Context) {
	var param web.GoodsListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	goodsList, rows := g.GetList(param)
	response.SuccessPage(constant.Selected, goodsList, rows, c)
}

func (g *AppGoods) GetGoodsList(c *gin.Context) {
	var param app.GoodsQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	goodsList := g.GetList(param)
	response.Success(constant.Selected, goodsList, c)
}

func (g *AppGoods) GetGoodsInfo(c *gin.Context) {
	var param app.GoodsInfoParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	goodsInfo := g.GetInfo(param)
	response.Success(constant.Selected, goodsInfo, c)
}

func (g *AppGoods) SearchGoods(c *gin.Context) {
	var param app.GoodsSearchParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	goodsList := g.Search(param)
	response.Success(constant.Selected, goodsList, c)
}
