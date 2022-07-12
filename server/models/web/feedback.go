package web

// 问题反馈参数模型
type FeedbackParam struct {
	Content string `json:"content" binding:"required"` // 反馈的内容
}
