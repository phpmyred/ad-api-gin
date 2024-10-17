// Package db
// @Description: REDIS 相关的类
package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var (
	Rdb      *redis.Client
	RdbStore *RedisStore
)

// RedisInit
// @Description 初始化Redis链接方法
// @Author aDuo 2024-08-14 22:46:10
// @Param addr 	链接地址带端口
// @Param password  密码
// @Param db
// @Return *redis.Client
// @Return error
func RedisInit(addr string, password string, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // 如果没有密码，则留空
		DB:       db,       // 使用默认数据库
	})
	Rdb = rdb
	RdbStore = NewDefaultRedisStore()

	// 使用Ping redis测试连接
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Redis链接出错，错误信息：", err)
		return nil, err
	}
	fmt.Println("Redis连接成功，Ping结果：", pong)

	return rdb, nil
}

// NewDefaultRedisStore
// @Description  自己封装 Redis初始化自定义
// @Author aDuo 2024-08-14 22:44:56
// @Return *RedisStore
func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * time.Duration(viper.GetInt("redis.expiration")),
		PreKey:     viper.GetString("redis.preKey"),
		Context:    context.TODO(),
	}
}

// RedisStore 自定义Redis的 初始化结构体
// @Description:
type RedisStore struct {
	//
	// Expiration
	// @Description: 过期时间 默认为秒
	//
	Expiration time.Duration
	//
	// PreKey
	// @Description: 拼接KEY的前缀
	//
	PreKey  string
	Context context.Context
}

// Set
// @Description 带配置前缀的redis set 方法 带默认过期时间 过期时间在配置文件里
// @Author aDuo 2024-08-14 22:40:33
// @Param key 	 键
// @Param value  值
// @Return error
func (rs *RedisStore) Set(key string, value string) error {
	err := Rdb.Set(rs.Context, rs.PreKey+key, value, rs.Expiration).Err()
	if err != nil {
		//Z.LoggerObj.Error("RedisStoreSetError!", err)
		return err
	}
	return nil
}

// SetT  set
// @Description  自动带key前缀的Redis自定义过期时间的set 方法
// @Author aDuo 2024-08-14 22:39:12
// @Param key 	 键
// @Param value  值
// @Param s 	 过期时间 单位 秒
// @Return error
func (rs *RedisStore) SetT(key string, value string, s int) error {

	err := Rdb.Set(rs.Context, rs.PreKey+key, value, time.Second*time.Duration(s)).Err()
	if err != nil {
		//Z.LoggerObj.Error("RedisStoreSetError!", err)

		return err
	}
	return nil
}

// Get
// @Description  自动带key前缀的Redis自定义get 方法
// @Author aDuo 2024-08-14 22:47:11
// @Param key 	 键
// @Return string 值
func (rs *RedisStore) Get(key string) string {
	key = rs.PreKey + key
	val, err := Rdb.Get(rs.Context, key).Result()
	if err != nil {

		//Z.LoggerObj.Error("RedisStoreGetError!", err)

		return ""
	}

	return val
}

// Del
// @Description  自动带key前缀的Redis自定义del 方法
// @Author aDuo 2024-08-14 22:48:03
// @Param key
// @Return bool 成功返回 true 失败返回 false
func (rs *RedisStore) Del(key string) bool {
	err := Rdb.Del(rs.Context, rs.PreKey+key).Err()
	if err != nil {
		//Z.LoggerObj.Error("RedisStoreClearError!", err)
		return false
	}
	return true
}

// Verify
// @Description 判断redis 该KEY获取到的值是否跟你的值相等
// @Author aDuo 2024-08-14 22:49:06
// @Param key    要判断值的键
// @Param answer 要判断的值
// @Return bool	 成功返回 TRUE 失败返回 False
func (rs *RedisStore) Verify(key, answer string) bool {
	v := rs.Get(key)
	return v == answer
}

func (rs RedisStore) IsExist(key string) bool {
	res := rs.Get(key)

	return res != ""
}

func (rs RedisStore) Expire(key string, s int) error {
	err := Rdb.Expire(rs.Context, rs.PreKey+key, time.Second*time.Duration(s)).Err()
	if err != nil {
		//Z.LoggerObj.Error("RedisStoreSetError!", err)

		return err
	}
	return nil
}
