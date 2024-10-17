package db

import (
	"fmt"
	"github.com/spf13/viper"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 链接后的数据库常量
var (
	DB *gorm.DB
)

func MysqlInit() error {
	_, err := Mysql(
		viper.GetString("db.hostname"),
		viper.GetInt("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
	)
	return err
}

// Mysql
// @Description  MYSQL 链接方法
// @Author aDuo 2024-08-14 22:31:33
// @Param hostname  链接地址
// @Param port 		端口
// @Param username 	用户名
// @Param password 	密码
// @Param dbname 	库名
// @Return *gorm.DB	gorm类
// @Return error	链接错误信息
func Mysql(hostname string, port int, username string, password string, dbname string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			username, password, hostname, port, dbname))
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	db.DB().SetMaxIdleConns(50)
	// SetMaxOpenConns 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(1000)
	// SetConnMaxLifetime 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(time.Hour)
	DB = db
	return db, nil
}
