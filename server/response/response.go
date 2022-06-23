package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// PageResult 分页结果返回
type PageResult struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

// Success 请求成功返回
func Success(message string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{200, message, data})
}

// Failed 请求失败返回
func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{400, message, 0})
}

// SuccessPage 请求成功返回分页结果
func SuccessPage(message string, data interface{}, rows int64, c *gin.Context) {
	page := &PageResult{Total: rows, List: data}
	c.JSON(http.StatusOK, Response{200, message, page})
}
