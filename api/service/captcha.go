// 验证码 服务层

package service

import (
	"admin-go-api/common/util"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 使用 redis 作为 store
var store = util.RedisStore{}

// CaptMake 生成验证码
func CaptMake() (id, b64s string) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	// 配置验证信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          6,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64s, _ := captcha.Generate()

	return lid, lb64s
}

// CaptVerify 验证 captcha 是否正确
func CaptVerify(id string, capt string) bool {
	if store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}
