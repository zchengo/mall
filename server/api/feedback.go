package api

import (
	"github.com/gin-gonic/gin"
	"imall/constant"
	"imall/models/web"
	"imall/response"
	"imall/service"
)

type WebFeedback struct {
	service.WebFeedBackService
}

func GetWebFeedback() *WebFeedback {
	return &WebFeedback{}
}

func (f *WebFeedback) SendFeedback(context *gin.Context) {
	var param web.FeedbackParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if err := f.Send(param); err != nil {
		response.Failed(constant.SendFailed, context)
		return
	}
	response.Success(constant.SendSuccess, nil, context)
}
