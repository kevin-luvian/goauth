package gredis

import (
	"errors"
	gtime "time"

	"github.com/gomodule/redigo/redis"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	DefaultTTL = 10 * gtime.Minute
)

var (
	ErrNoKey  = errors.New("redis missing key")
	redisPool *redis.Pool
)

// Setup Initialize the Redis instance
func Setup() {
	s := setting.Redis

	redisPool = &redis.Pool{
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

func Ping() error {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	return err
}

// Set a key/value
func Set(key string, value string, time ...gtime.Duration) error {
	if len(time) == 0 {
		time = append(time, DefaultTTL)
	}

	conn := redisPool.Get()
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

func SetStruct(key string, value interface{}, time ...gtime.Duration) error {
	if len(time) == 0 {
		time = append(time, DefaultTTL)
	}

	conn := redisPool.Get()
	defer conn.Close()

	enc, err := msgpack.Marshal(value)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, enc)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time[0].Seconds())
	if err != nil {
		return err
	}

	return nil
}

// Get get a key
func Get(key string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err == redis.ErrNil {
		return "", ErrNoKey
	} else if err != nil {
		return "", err
	}

	return reply, nil
}

// GetStruct get a key
func GetStruct(key string, target interface{}) error {
	conn := redisPool.Get()
	defer conn.Close()

	buf, err := redis.Bytes(conn.Do("GET", key))
	if err == redis.ErrNil {
		return ErrNoKey
	} else if err != nil {
		return err
	}

	err = msgpack.Unmarshal(buf, target)
	return err
}

// Exists check a key
func Exists(key string) bool {
	conn := redisPool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := redisPool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := redisPool.Get()
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
