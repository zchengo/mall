package api

import (
	"github.com/gin-gonic/gin"
	"imall/constant"
	"imall/models/app"
	"imall/response"
	"imall/service"
)

type AppAddress struct {
	service.AppAddressService
}

func GetAppAddress() *AppAddress {
	return &AppAddress{}
}

func (a *AppAddress) SaveAddress(c *gin.Context) {
	var param app.AddressSaveParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if count := a.Save(param); count > 0 {
		response.Success(constant.SaveSuccess, count, c)
		return
	}
	response.Failed(constant.SaveFailed, c)
}

func (a *AppAddress) DeleteAddress(c *gin.Context) {
	var param app.AddressDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	if count := a.Delete(param); count > 0 {
		response.Success(constant.Deleted, count, c)
		return
	}
	response.Failed(constant.NotDeleted, c)
}

func (a *AppAddress) GetAddressList(c *gin.Context) {
	var param app.AddressListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	categoryList := a.GetList(param)
	response.Success(constant.Selected, categoryList, c)
}

func (a *AppAddress) GetAddressInfo(c *gin.Context) {
	var param app.AddressInfoParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}
	categoryList := a.GetInfo(param)
	response.Success(constant.Selected, categoryList, c)
}
