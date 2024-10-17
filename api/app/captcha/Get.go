package captcha

import (
	constant "github.com/TestsLing/aj-captcha-go/const"
	"github.com/WangaduoApi/ad-api-gin/api/contextData/response"
	"github.com/WangaduoApi/ad-api-gin/utility/Z/captcha"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {

	data, err := captcha.Factory.GetService(constant.BlockPuzzleCaptcha).Get()

	if err != nil {
		panic("获取验证码错误！")
	}

	response.Success(ctx, "", data)

}
