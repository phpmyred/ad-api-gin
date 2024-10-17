package captcha

import (
	"github.com/WangaduoApi/ad-api-gin/api/contextData/request"
	"github.com/WangaduoApi/ad-api-gin/api/contextData/response"
	"github.com/WangaduoApi/ad-api-gin/utility/Z/captcha"
	"github.com/gin-gonic/gin"
)

type ClientParams struct {
	Token       string `json:"token"`
	PointJson   string `json:"pointJson"`
	CaptchaType string `json:"captchaType"`
}

func Check(ctx *gin.Context) {

	params := request.GetReqParam(ctx, ClientParams{})

	if params.Token == "" || params.PointJson == "" || params.CaptchaType == "" {
		panic("请求参数错误！")
	}

	ser := captcha.Factory.GetService(params.CaptchaType)

	err := ser.Check(params.Token, params.PointJson)
	if err != nil {
		panic(err)

	}
	response.Success(ctx, "", nil)

}
