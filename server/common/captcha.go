package common

import "github.com/mojocn/base64Captcha"

var store = base64Captcha.DefaultMemStore
var driver = base64Captcha.NewDriverDigit(40, 120, 4, 0.7, 80)

// 生成验证码
func GenerateCaptcha() (id, b64s string, err error) {
	return base64Captcha.NewCaptcha(driver, store).Generate()
}

// 验证验证码
func VerifyCaptcha(id, value string) bool {
	return store.Verify(id, value, true)
}
