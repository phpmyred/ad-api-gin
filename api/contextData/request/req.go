package request

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/tidwall/gjson"
	"io"
)

// GetReqParam
// @Description 解析HTTP 请求参数 返回结构体
// @Author aDuo 2024-08-30 13:43:09
// @Param ctx
// @Param params
func GetReqParam[DTO any](ctx *gin.Context, params DTO) DTO {

	reqData, err := ctx.GetRawData()
	if err != nil {
		panic("解析请求参数错误！")
	}
	if err = json.Unmarshal(reqData, &params); err != nil {
		// 处理绑定错误
		panic("请求参数错误!")
	}
	return params
}

func GetReqParamByBytes(context *gin.Context) []byte {
	bytes, err := io.ReadAll(context.Request.Body)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func GetReqParamByJson(context *gin.Context) gjson.Result {
	bytes, _ := io.ReadAll(context.Request.Body)

	return gjson.ParseBytes(bytes)
}
