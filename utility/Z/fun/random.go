package fun

import (
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

// RandomString
// @Description  生成随机字符串
// @Author aDuo 2024-08-23 03:46:13
// @Param n
// @Return string
func RandomString(n int) string {
	// 定义字符集
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())
	// 使用strings.Builder提高字符串构建的性能
	var b strings.Builder
	for i := 0; i < n; i++ {
		// 从字符集中随机选取字符，拼接到字符串中
		b.WriteRune(letterRunes[rand.Intn(len(letterRunes))])
	}
	// 返回生成的随机字符串
	return b.String()
}

// RandomUUIDString
// @Description UUID获取随机UUID
// @Author aDuo 2024-08-29 13:23:55
// @Return string
func RandomUUIDString(i int) string {
	u1 := uuid.NewV4()
	original := u1.String()
	withoutDashes := strings.ReplaceAll(original, "-", "")

	return withoutDashes[0:i]
}

func UUIDString() string {
	u1 := uuid.NewV4()
	return u1.String()
}
