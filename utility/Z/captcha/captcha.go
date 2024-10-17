package captcha

import (
	"github.com/mojocn/base64Captcha"
)

// StringCaptcha
// @Description: 验证码工具类
type StringCaptcha struct {
	captcha *base64Captcha.Captcha
}

// NewCaptcha
// @Description 初始化验证码
// @Author aDuo 2024-09-02 04:13:47
// @Return *StringCaptcha
func NewCaptcha() *StringCaptcha {
	// store
	store := base64Captcha.DefaultMemStore

	// 包含数字和字母的字符集
	source := "1234567890"
	//source := "1234"
	// driver
	driver := base64Captcha.NewDriverString(
		42,     // height int
		110,    // width int
		6,      // noiseCount int
		1,      // showLineOptions int
		6,      // length int
		source, // source string
		nil,    // bgColor *color.RGBA
		nil,    // fontsStorage FontsStorage
		nil,    // fonts []string
	)
	captcha := base64Captcha.NewCaptcha(driver, store)
	return &StringCaptcha{
		captcha: captcha,
	}
}

// Generate
// @Description  生成验证码
// @Author aDuo 2024-08-19 16:20:31
// @Return string
// @Return string
// @Return string
func (stringCaptcha *StringCaptcha) Generate() (string, string, string) {
	id, b64s, answer, _ := stringCaptcha.captcha.Generate()
	return id, b64s, answer
}

// Verify
// @Description 校验验证码
// @Author aDuo 2024-09-02 04:13:29
// @Param id
// @Param answer
// @Return bool
func (stringCaptcha *StringCaptcha) Verify(id string, answer string) bool {
	return stringCaptcha.captcha.Verify(id, answer, true)
}
