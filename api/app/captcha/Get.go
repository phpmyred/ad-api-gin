package captcha

import (
	constant "github.com/TestsLing/aj-captcha-go/const"
	"github.com/WangaduoApi/ad-api-gin/api/contextData/request"
	"github.com/WangaduoApi/ad-api-gin/api/contextData/response"
	"github.com/WangaduoApi/ad-api-gin/utility/Z/captcha"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {

	var t = request.GetReqParamByJson(ctx).Get("captchaType").String()
	if t == "" {
		t = constant.BlockPuzzleCaptcha
	}
	data, err := captcha.Factory.GetService(t).Get()

	if err != nil {
		panic("获取验证码错误！")
	}

	response.Success(ctx, "", data)

}
