package api

import (
	"github.com/gin-gonic/gin"
	"imall/common"
	"imall/constant"
	"imall/models/web"
	"imall/response"
	"imall/service"
)

type WebUser struct {
	service.UserService
}

type AppUser struct {
	service.AppUserService
}

func GetWebUser() *WebUser {
	return &WebUser{}
}

func GetAppUser() *AppUser {
	return &AppUser{}
}

// 获取验证码
func (u *WebUser) GetCaptcha(c *gin.Context) {
	id, base64s, _ := common.GenerateCaptcha()
	data := map[string]interface{}{"captchaId": id, "captchaImg": base64s}
	response.Success("操作成功", data, c)
}

// 用户登录
func (u *WebUser) UserLogin(c *gin.Context) {
	var param web.LoginParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, c)
		return
	}

	// 检查验证码
	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
		response.Failed("验证码错误", c)
		return
	}
	// 生成token
	uid := u.Login(param)
	if uid > 0 {
		token, _ := common.GenerateToke(param.Username)
		userInfo := web.UserInfo{
			Uid:   uid,
			Token: token,
		}
		response.Success("登录成功", userInfo, c)
		return
	}
	response.Failed("用户名或密码错误", c)
}

func (u *AppUser) UserLogin(context *gin.Context) {
	code := context.PostForm("code")
	if code == "" {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if userInfo := u.Login(code); userInfo != nil {
		response.Success(constant.LoginSuccess, userInfo, context)
		return
	}
	response.Failed(constant.LoginFailed, context)
}
