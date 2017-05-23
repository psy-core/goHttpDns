package redis

import (
	"time"

	"github.com/cihub/seelog"
	redigo "github.com/garyburd/redigo/redis"
)

// RedisConfig Type
type RedisConfig struct {
	Host           string `yaml:"host"`
	Network        string `yaml:"network"`
	Database       int    `yaml:"db"`
	ConnectTimeout int64  `yaml:"connectTimeout"`
	ReadTimeout    int64  `yaml:"readTimeout"`
	WriteTimeout   int64  `yaml:"writeTimeout"`
	MaxActive      int    `yaml:"maxActive"`
	MaxIdle        int    `yaml:"maxIdle"`
	IdleTimeout    int    `yaml:"idleTimeout"`
	Wait           bool   `yaml:"wait"`
}

//redisPool 连接池
var redisPool *redigo.Pool

//IsValid Func
func (rc *RedisConfig) IsValid() bool {
	return len(rc.Host) > 0 && len(rc.Network) > 0 && rc.Database >= 0 && rc.Database < 16
}

//Init Func
func Init(redisConf *RedisConfig) {
	redisPool = &redigo.Pool{
		MaxActive:   redisConf.MaxActive,
		MaxIdle:     redisConf.MaxIdle,
		IdleTimeout: time.Duration(redisConf.IdleTimeout) * time.Second,
		Wait:        redisConf.Wait,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", redisConf.Host, redigo.DialDatabase(redisConf.Database),
				redigo.DialConnectTimeout(time.Duration(redisConf.ConnectTimeout)*time.Millisecond),
				redigo.DialReadTimeout(time.Duration(redisConf.ReadTimeout)*time.Millisecond),
				redigo.DialWriteTimeout(time.Duration(redisConf.WriteTimeout)*time.Millisecond))
			if err != nil {
				seelog.Errorf("[Redis] Dial error: %v", err)
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

//Set 设置redis数据, string-string
func Set(key, value string) error {
	conn := redisPool.Get()
	if conn.Err() != nil {
		return conn.Err()
	}
	defer conn.Close()
	if _, err := conn.Do("set", key, value); err != nil {
		return err
	}
	return nil
}

//SetEx 设置redis数据, 带有超时, string-string
func SetEx(key, value string, expire int) error {
	conn := redisPool.Get()
	if conn.Err() != nil {
		return conn.Err()
	}
	defer conn.Close()
	if _, err := conn.Do("setex", key, expire, value); err != nil {
		return err
	}
	return nil
}

//Get 获取redis数据，string-string
func Get(key string) (string, error) {
	conn := redisPool.Get()
	if conn.Err() != nil {
		return "", conn.Err()
	}
	defer conn.Close()
	value, err := redigo.String(conn.Do("get", key))
	if err != nil {
		return "", err
	}
	return value, nil
}

//MGet 插入多个键值数据到redis string-string
func MGet(keys []string) ([]string, error) {
	conn := redisPool.Get()
	if conn.Err() != nil {
		return nil, conn.Err()
	}
	defer conn.Close()
	newKeyArray := make([]interface{}, len(keys))
	for index, key := range keys {
		newKeyArray[index] = interface{}(key)
	}
	values, err := redigo.Strings(conn.Do("mget", newKeyArray...))
	if err != nil {
		return nil, err
	}
	return values, nil
}

//HGet 获取hash类型的redis数据 key-field-value
//注意 可能值经过压缩等变化，redis取值不考虑编码问题，返回的string类型为utf8编码
//编码转换等问题交由业务逻辑处理
func HGet(key, field string) (string, error) {
	conn := redisPool.Get()
	if conn.Err() != nil {
		return "", conn.Err()
	}
	defer conn.Close()
	value, err := redigo.String(conn.Do("hget", key, field))
	if err != nil {
		return "", err
	}
	return value, nil
}

//HGetAll 获取给定key的hash类型的所有redis数据 key-field-value
//注意 可能值经过压缩等变化，redis取值不考虑编码问题，返回的string值类型为utf8编码
//编码转换等问题交由业务逻辑处理
func HGetAll(key string) (map[string]string, error) {
	conn := redisPool.Get()
	if conn.Err() != nil {
		return nil, conn.Err()
	}
	defer conn.Close()
	value, err := redigo.StringMap(conn.Do("hgetall", key))
	if err != nil {
		return nil, err
	}
	return value, nil
}
