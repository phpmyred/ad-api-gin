package file

import (
	"io/ioutil"
	"log"
	"os"
)

// MakerDir
// @Description 检查目录是否存在 不存在则创建
// @Author aDuo 2024-08-31 22:38:42
// @Param dirPath
func MakerDir(dirPath string) {

	// 使用Stat来检查目录是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 目录不存在，使用MkdirAll创建
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			panic(err) // 处理错误
		}
	}
}

func GetFile(fileDir string) string {
	file, err := os.Open(fileDir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 读取文件内容到[]byte
	byteValue, err := ioutil.ReadAll(file)

	return string(byteValue)
}
