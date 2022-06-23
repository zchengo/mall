package api

import (
	"github.com/gin-gonic/gin"
	"imall/constant"
	"imall/response"
	"imall/service"
)

type WebStatistics struct {
	service.WebStatisticsService
}

func GetWebStatistics() *WebStatistics {
	return &WebStatistics{}
}

func (s *WebStatistics) GetOrderData(context *gin.Context) {
	orderData := s.OrderData()
	response.Success(constant.Selected, orderData, context)
}

func (s *WebStatistics) GetWeekData(context *gin.Context) {
	weekData := s.WeekData()
	response.Success(constant.Selected, weekData, context)
}
