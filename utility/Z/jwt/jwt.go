package jwt

import (
	"github.com/spf13/viper"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(viper.GetString("jwt.signingKey"))

type Claims struct {
	UserUUId string
	UserId   uint
	jwt.StandardClaims
}

// ReleaseToken 签发token的方法
// @Description 生成设置TOKEN的方法
// @Author aDuo 2024-08-14 22:59:01
// @Param user
// @Return string
// @Return error
func releaseToken(ID uint, UUID, subject string) string {
	// 过期时间设置 7天
	//expirationTime := time.Now().Add(7 * 24 * time.Hour)
	expTime := time.Duration(viper.GetInt("jwt.expiresTime")) * time.Hour
	expirationTime := time.Now().Add(expTime)

	claims := &Claims{
		UserId: ID,
		// 将用户ID 写入TOKEN 之后可以根据TOKEN解析获取到
		UserUUId: UUID,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			// 签发发人
			Issuer:  viper.GetString("jwt.issuer"),
			Subject: subject,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

// ParseToken 解析 TOKEN
// @Description 这个是能将保存在TOKEN的里的 信息解析出来的方法
// @Author aDuo 2024-08-14 23:03:51
// @Param tokenString  	token
// @Return *jwt.Token
// @Return *Claims		解析出来的保存到 TOKEN的信息
// @Return error
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}

// ReleaseAppToken
// @Description ADMIN 签发 token
// @Author aDuo 2024-08-21 15:25:48
// @Param user
// @Return string

func ReleaseAppToken(ID uint, UUID, subject string) string {

	var token = releaseToken(ID, UUID, "app")

	// 签发完 TOKEN 将TOKEN  存入数据库
	//err := model.User{}.SetUserTokenById(user.ID, user.Token)
	//if err != nil {
	//	panic(err)
	//}
	// 用用户信息 刷新存入REDIS
	//redis.UserInfo{}.ReloadUser(user.UUID)

	return token
}
