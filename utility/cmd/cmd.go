package cmd

import (
	"fmt"
	"github.com/WangaduoApi/ad-api-gin/router"
	"github.com/WangaduoApi/ad-api-gin/utility/Z"
	"github.com/WangaduoApi/ad-api-gin/utility/Z/captcha"
	"github.com/WangaduoApi/ad-api-gin/utility/db"
	"github.com/spf13/cobra"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{}
)

func initConfig() {
	MustInit(os.Stdout, cfgFile) // 配置初始化
}

// init
// @Description 初始化命令参数
// @Author aDuo 2024-08-15 17:47:25
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/dev.yaml", "config file (default is $HOME/.cobra.yaml)")

	rootCmd.PersistentFlags().StringP("author", "a", "Wang A Duo", "版权归属的作者姓名")

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "是否开启debug")

	err := viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	if err != nil {
		return
	}
	viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))

	fmt.Println(viper.GetBool("gin.mode"))
	fmt.Println(viper.GetString("author"))
	initConfig()

}

// Execute
// @Description 初始化各种配置
// @Author aDuo 2024-08-15 17:47:08
// @Return error
func Execute() error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {

		// MYSQL配置
		err := db.MysqlInit()
		if err != nil {
			return err
		}
		// 使用日志
		defer db.DB.Close()

		// redis 配置
		_, err = db.RedisInit(
			viper.GetString("redis.addr"),
			viper.GetString("redis.password"),
			viper.GetInt("redis.db"),
		)
		if err != nil {
			return err
		}

		// 在程序结束时 关闭redis连接
		defer db.Rdb.Close()

		//初始化验证码配置
		captcha.CaptchaInit()

		port := viper.GetString("server.port")

		r := router.SetupRouter()
		Z.LoggerObj.Info("port = *** =", port)

		err = r.Run(port)
		if err != nil {
			return err
		}
		return http.ListenAndServe(port, nil) // listen and serve
	}

	return rootCmd.Execute()

}
