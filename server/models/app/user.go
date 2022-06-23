package app

// 用户数据映射模型
type User struct {
	Id       uint64 `gorm:"primaryKey"` // 用户编号
	OpenId   string `gorm:"open_id"`    // 微信用户唯一标识
	Username string `gorm:"username"`   // 用户名称
	Password string `gorm:"password"`   // 用户密码
	Status   uint   `gorm:"status"`     // 用户状态
	Created  string `gorm:"created"`    // 创建时间
	Updated  string `gorm:"updated"`    // 更新时间
}

// 用户登录凭证校验模型
type Code2Session struct {
	Code      string
	AppId     string
	AppSecret string
}

// 凭证校验后返回的JSON数据包模型
type Code2SessionResult struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    uint   `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// 用户信息,OpenID用户唯一标识
type UserInfo struct {
	OpenId string `json:"openId"`
}
