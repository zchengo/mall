package api

import (
	"fmt"
	"imall/constant"
	"imall/models/app"
	"imall/response"
	"imall/service"
	"github.com/gin-gonic/gin"
)

type AppCart struct {
	service.AppCartService
}

func GetAppCart() *AppCart {
	return &AppCart{}
}

func (c *AppCart) AddCart(context *gin.Context) {
	var param app.CartAddParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		fmt.Println(err)
		return
	}
	added := c.Add(param)
	if added > 0 {
		response.Success(constant.Created, nil, context)
		return
	}
	response.Failed(constant.NotCreated, context)
}

func (c *AppCart) DeleteCart(context *gin.Context) {
	var param app.CartDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := c.Delete(param); count > 0 {
		response.Success(constant.Deleted, count, context)
		return
	}
	response.Failed(constant.NotDeleted, context)
}

// AppDeleteCart 微信小程序，清除购物车中的商品
func (c *AppCart) ClearCart(context *gin.Context) {
	var param app.CartClearParam
	if err := context.ShouldBind(&param); err != nil {
		fmt.Println(param)
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := c.Clear(param); count > 0 {
		response.Success("清除成功", count, context)
		return
	}
	response.Failed("清除失败", context)
}

func (c *AppCart) GetCartInfo(context *gin.Context) {
	var param app.CartQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	info := c.GetInfo(param)
	if len(info.CartItem) == 0 {
		response.Success("购物车竟然是空的", info, context)
	}
	response.Success("查询成功", info, context)
}
