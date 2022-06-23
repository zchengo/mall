package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"imall/constant"
	"imall/models/app"
	"imall/models/web"
	"imall/response"
	"imall/service"
)

type WebOrder struct {
	service.WebOrderService
}

type AppOrder struct {
	service.AppOrderService
}

func GetWebOrder() *WebOrder {
	return &WebOrder{}
}

func GetAppOrder() *AppOrder {
	return &AppOrder{}
}

func (o *WebOrder) DeleteOrder(context *gin.Context) {
	var param web.OrderDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := o.Delete(param); count > 0 {
		response.Success(constant.Deleted, count, context)
		return
	}
	response.Failed(constant.NotDeleted, context)
}

func (o *WebOrder) UpdateOrder(context *gin.Context) {
	var param web.OrderUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := o.Update(param); count > 0 {
		response.Success("更新成功", count, context)
		return
	}
	response.Failed("更新失败", context)
}

func (o *WebOrder) GetOrderList(context *gin.Context) {
	var param web.OrderListParam
	if err := context.ShouldBind(&param); err != nil {
		fmt.Println(err)
		response.Failed(constant.ParamInvalid, context)
		return
	}
	productList, rows := o.GetList(param)
	response.SuccessPage("查询成功", productList, rows, context)
}

func (o *WebOrder) GetOrderDetail(context *gin.Context) {
	var param web.OrderDetailParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	productDetail := o.GetDetail(param)
	response.Success("查询成功", productDetail, context)
}

func (o *AppOrder) UpdateOrder(context *gin.Context) {
	var param app.OrderUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := o.Update(param); count > 0 {
		response.Success("更新成功", count, context)
		return
	}
	fmt.Println(param)
	response.Failed("更新失败", context)
}

func (o *AppOrder) SubmitOrder(context *gin.Context) {
	var param app.OrderSubmitParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := o.Submit(param); count > 0 {
		response.Success("提交成功", count, context)
		return
	}
	response.Failed("提交失败", context)
}

func (o *AppOrder) GetOrderList(context *gin.Context) {
	var param app.OrderQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	orderList := o.GetList(param)
	response.Success("查询成功", orderList, context)
}
