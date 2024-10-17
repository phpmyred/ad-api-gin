package main

import (
	"github.com/WangaduoApi/ad-api-gin/utility/cmd"
	"os"
)

// main
// @Description  主入口
// @Author aDuo 2024-08-14 22:14:25
func main() {

	if err := cmd.Execute(); err != nil {
		println("start fail: ", err.Error())
		os.Exit(-1)
	}
}
