package libs

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	RedisPool *redis.Pool
	redisConf redisConfig
)

type redisConfig struct {
	host      string
	port      string
	password  string
	database  int
	maxIdle   int
	maxActive int
}

func InitRedisPool(host, port, password string, database, maxIdle, maxActive int) {
	log.Info("初始化redis连接池")
	redisConf = redisConfig{
		host:      host,
		port:      port,
		password:  password,
		database:  database,
		maxIdle:   maxIdle,
		maxActive: maxActive,
	}
	address := redisConf.host + ":" + redisConf.port
	optPassword := redis.DialPassword(redisConf.password)
	optDatabase := redis.DialDatabase(redisConf.database)
	RedisPool = &redis.Pool{
		MaxIdle:     redisConf.maxIdle,
		MaxActive:   redisConf.maxActive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (c redis.Conn, err error) {

			c, err = redis.Dial("tcp", address, optDatabase, optPassword)
			if err != nil {
				// handle error
				log.Error("redis:", err)
			}
			return

		},
		// check the health of an idle connection before the connection is returned to the application
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func GetRedisConn() (conn redis.Conn) {
	conn = RedisPool.Get()
	return
}

func GetRedisConnDB(database int) (conn redis.Conn, err error) {
	conn = GetRedisConn()
	if database != redisConf.database {
		_, err = conn.Do("SELECT", database)
	}
	return
}

// ------------------------------------------------------------------------

//MyRedisCache
type RedisCache struct {
}

//NewRedisCache
func NewRedisCache() *RedisCache {
	return &RedisCache{}
}

//Get return cached value
func (mem *RedisCache) Get(key string) (val interface{}) {
	conn := GetRedisConn()
	_bytes, _ := conn.Do("GET", key)
	if _bytes == nil {
		_bytes = []uint8{}
	}
	_ = json.Unmarshal(_bytes.([]byte), &val)
	return
}

//IsExist
func (mem *RedisCache) IsExist(key string) (exist bool) {
	conn := GetRedisConn()
	exist, _ = redis.Bool(conn.Do("EXISTS", key))
	return
}

//Set cached value with key and expire time.
func (mem *RedisCache) Set(key string, val interface{}, timeout time.Duration) (err error) {
	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return err
	}
	conn := GetRedisConn()
	_, err = conn.Do("SETEX", key, int32(timeout/time.Second), data)
	return
}

//Del
func (mem *RedisCache) Delete(key string) (err error) {
	conn := GetRedisConn()
	_, err = conn.Do("DEL", key)
	return
}
