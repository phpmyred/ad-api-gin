package encryption

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	harsher := md5.New()
	// 写入需要哈希的数据
	harsher.Write([]byte(str))
	// 计算哈希值
	hash := harsher.Sum(nil)
	// 返回十六进制表示的哈希值
	return hex.EncodeToString(hash)
}
