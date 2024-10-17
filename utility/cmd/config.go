package cmd

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 配置文件类型
const (
	configType = "yaml"
)

// Init
// @Description  初始化引入配置文件
// @Author aDuo 2024-08-14 22:27:53
// @Param output
// @Param configFile
// @Return error
func Init(output io.Writer, configFile string) error {
	if output == nil {
		output = ioutil.Discard
	}

	viper.SetConfigFile(configFile)
	viper.SetConfigType(configType) // or viper.SetConfigType("YAML")
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		_, _ = fmt.Fprintf(output, "Config file changed %s \n", e.Name)
	})
	return nil
}

// MustInit
// @Description  配置文件引入初始化的方法
// @Author aDuo 2024-08-14 22:28:58
// @Param output
// @Param conf  配置文件的名称
// @Return {
func MustInit(output io.Writer, conf string) { // MustInit if fail panic
	if err := Init(output, conf); err != nil {
		// 以下是引入配置文件错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
