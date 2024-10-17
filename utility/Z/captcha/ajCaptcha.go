package captcha

import (
	config2 "github.com/TestsLing/aj-captcha-go/config"
	constant "github.com/TestsLing/aj-captcha-go/const"
	"github.com/TestsLing/aj-captcha-go/service"
	"github.com/spf13/viper"
	"image/color"
)

// **********************默认配置***************************************************
// 默认配置，可以根据项目自行配置，将其他类型配置序列化上去
//var captchaConfig = config2.NewConfig()

// *********************自定义配置**************************************************
// 水印配置（参数可从业务系统自定义）
var watermarkConfig = &config2.WatermarkConfig{
	FontSize: 12,
	Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
	Text:     viper.GetString("author"),
}

// 点击文字配置（参数可从业务系统自定义）
var clickWordConfig = &config2.ClickWordConfig{
	FontSize: 25,
	FontNum:  4,
}

// 滑动模块配置（参数可从业务系统自定义）
var blockPuzzleConfig = &config2.BlockPuzzleConfig{Offset: 10}

// 行为校验配置模块（具体参数可从业务系统配置文件自定义）

var isChaVal = viper.GetString("")
var captchaConfig = config2.BuildConfig(constant.RedisCacheKey, "./config", watermarkConfig,
	clickWordConfig, blockPuzzleConfig, 2*60)

var captchaConfig = config2.BuildConfig(constant.MemCacheKey, "./config", watermarkConfig,
	clickWordConfig, blockPuzzleConfig, 2*60)

// 服务工厂，主要用户注册 获取 缓存和验证服务
var Factory = service.NewCaptchaServiceFactory(captchaConfig)

//
// CaptchaInit
// @Description
// @Author aDuo 2024-10-04 21:00:15

func CaptchaInit() {
	// 这里默认是注册了 内存缓存，但是不足以应对生产环境，希望自行注册缓存驱动 实现缓存接口即可替换（CacheType就是注册进去的 key）
	//Factory.RegisterCache(constant.MemCacheKey, service.NewMemCacheService(20)) // 这里20指的是缓存阈值

	//注册使用默认redis数据库
	//Factory.RegisterCache(constant.RedisCacheKey, service.NewDftRedisCacheService(20))
	//注册自定义配置redis数据库
	Factory.RegisterCache(constant.RedisCacheKey, service.NewConfigRedisCacheService([]string{viper.GetString("redis.addr")}, "", "", false, viper.GetInt("redis.db")))

	// 注册了两种验证码服务 可以自行实现更多的验证
	Factory.RegisterService(constant.ClickWordCaptcha, service.NewClickWordCaptchaService(Factory))
	Factory.RegisterService(constant.BlockPuzzleCaptcha, service.NewBlockPuzzleCaptchaService(Factory))
}
