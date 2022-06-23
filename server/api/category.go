package api

import (
	"github.com/gin-gonic/gin"
	"imall/constant"
	"imall/models/web"
	"imall/response"
	"imall/service"
)

type WebCategory struct {
	service.WebCategoryService
}

type AppCategory struct {
	service.AppCategoryService
}

func GetWebCategory() *WebCategory {
	return &WebCategory{}
}

func GetAppCategory() *AppCategory {
	return &AppCategory{}
}

func (c *WebCategory) CreateCategory(context *gin.Context) {
	var param web.CategoryCreateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := c.Create(param); count > 0 {
		response.Success(constant.Created, count, context)
		return
	}
	response.Failed(constant.NotCreated, context)
}

func (c *WebCategory) DeleteCategory(context *gin.Context) {
	var param web.CategoryDeleteParam
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

func (c *WebCategory) UpdateCategory(context *gin.Context) {
	var param web.CategoryUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if count := c.Update(param); count > 0 {
		response.Success(constant.Updated, count, context)
		return
	}
	response.Failed(constant.NotUpdated, context)
}

func (c *WebCategory) GetCategoryList(context *gin.Context) {
	var param web.CategoryQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	categoryList := c.GetList(param)
	response.Success(constant.Selected, categoryList, context)
}

func (c *WebCategory) GetCategoryOption(context *gin.Context) {
	option := c.GetOption()
	response.Success(constant.Selected, option, context)
}

func (c *AppCategory) GetCategoryOption(context *gin.Context) {
	option := c.GetOption()
	response.Success(constant.Selected, option, context)
}
