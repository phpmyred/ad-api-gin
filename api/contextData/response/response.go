package response

import (
	"github.com/goccy/go-json"
	"github.com/spf13/viper"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Response
// @Description 接口返回的统一方法
// @Author aDuo 2024-08-14 22:18:42
// @Param ctx ctx上下文
// @Param httpStatus http状态码
// @Param code 返回的状态码
// @Param data 返回的数据
// @Param msg  返回的信息
func Response(ctx *gin.Context, httpStatus int, code int, data any, msg string) {

	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})

	ctx.Abort()
	return
}

// Success
// @Description 成功返回数据
// @Author aDuo 2024-08-14 22:20:10
// @Param ctx ctx上下文
// @Param data 返回的数据
// @Param msg  返回的信息
func Success(ctx *gin.Context, msg string, data interface{}) {
	// 返回信息加密
	//returnData, _ := aesCBCEncryptData(ctx, data)
	if msg == "" {
		msg = "success"
	}
	Response(ctx, http.StatusOK, viper.GetInt("reqCode.successCode"), data, msg)
}

// Fail 警告返回
// @Description 警告返回
// @Author aDuo 2024-08-14 22:25:55
// @Param ctx ctx上下文
// @Param data 返回的数据
// @Param msg  返回的信息
func Fail(ctx *gin.Context, msg string, data gin.H) {

	Response(ctx, http.StatusOK, viper.GetInt("reqCode.failCode"), data, msg)
}

// AuthFail
// @Description  鉴权无权限返回
// @Author aDuo 2024-08-14 22:24:07
// @Param ctx ctx上下文
// @Param msg  返回的信息
func AuthFail(ctx *gin.Context, msg string) {

	Response(ctx, http.StatusOK, viper.GetInt("reqCode.authCode"), nil, msg)
}

// Error
// @Description 系统错误返回
// @Author aDuo 2024-08-14 22:24:52
// @Param ctx
func Error(ctx *gin.Context) {
	Response(ctx, http.StatusOK, viper.GetInt("reqCode.errorCode"), nil, "系统错误！!!")
}

func BytesToGinH(b []byte) (gin.H, error) {
	var result gin.H
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}
