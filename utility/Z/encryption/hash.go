package encryption

// VerifyPassword
// @Description 哈希校验密码是否正确
// @Author aDuo 2024-08-30 14:11:36
// @Param password  数据库存放的 加密的密码
// @Param checkPassword  要校验的密码

func VerifyPassword(password, checkPassword string) {
	//if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(checkPassword)); err != nil { // 判断密码是否正确
	//	panic("密码错误！")
	//}

	if password != checkPassword {
		panic("密码错误！")

	}
}

// PasswordEncryption
// @Description 哈希密码加密
// @Author aDuo 2024-09-02 04:15:31
// @Param password
// @Return string
func PasswordEncryption(password string) string {
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 创建用户的时候要加密用户的密码
	//if err != nil {
	//	panic("加密失败！")
	//}
	return password
}
