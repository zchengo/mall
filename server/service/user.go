package service

import (
	"encoding/json"
	"fmt"
	"imall/common"
	"imall/global"
	"imall/models/app"
	"imall/models/web"
	"net/http"
)

type UserService struct {
}

type AppUserService struct {
}

// 商家用户登录
func (u *UserService) Login(param web.LoginParam) uint64 {
	var user web.User
	global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&user)
	return user.Id
}

// 买家用户登录
func (u *AppUserService) Login(code string) *app.UserInfo {
	var acsJson app.Code2SessionResult
	acs := app.Code2Session{
		Code:      code,
		AppId:     global.Config.Code2Session.AppId,
		AppSecret: global.Config.Code2Session.AppSecret,
	}
	api := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	res, err := http.DefaultClient.Get(fmt.Sprintf(api, acs.AppId, acs.AppSecret, acs.Code))
	if err != nil {
		fmt.Println("微信登录凭证校验接口请求错误")
		return nil
	}
	if err := json.NewDecoder(res.Body).Decode(&acsJson); err != nil {
		fmt.Println("decoder error...")
		return nil
	}

	// 查看用户是否已经存在
	rows := global.Db.Where("open_id = ?", acsJson.OpenId).First(&app.User{}).RowsAffected
	if rows == 0 {
		// 不存在，添加用户
		fmt.Println(acsJson.OpenId)
		user := app.User{
			OpenId: acsJson.OpenId,
			Status: 1,
			Created: common.NowTime(),
		}
		row := global.Db.Create(&user).RowsAffected
		if row == 0 {
			fmt.Println("add app user error...")
			return nil
		}
	}
	return &app.UserInfo{OpenId: acsJson.OpenId}
}
