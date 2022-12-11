package gredis

import (
	gtime "time"

	sredis "github.com/gin-contrib/sessions/redis"

	"github.com/gomodule/redigo/redis"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup() {
	s := setting.Redis

	RedisConn = &redis.Pool{
		MaxIdle:     s.MaxIdle,
		MaxActive:   s.MaxActive,
		IdleTimeout: s.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", s.Host)
			if err != nil {
				return nil, err
			}
			if s.Password != "" {
				if _, err := c.Do("AUTH", s.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t gtime.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	if err := Ping(); err != nil {
		logging.Fatalln("Error redis setup", err)
	}
}

func NewStore() (sredis.Store, error) {
	s := setting.Redis

	return sredis.NewStore(s.MaxIdle, "tcp", s.Host, s.Password, []byte(s.Salt))
}

func Ping() error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	return err
}

// Set a key/value
func Set(key string, value string, time ...gtime.Duration) error {
	if len(time) == 0 {
		time = append(time, 10*gtime.Minute)
	}

	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time[0].Seconds())
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
