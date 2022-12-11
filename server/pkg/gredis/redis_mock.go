package gredis

import (
	gtime "time"

	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"github.com/vmihailenco/msgpack/v5"
)

func MockRedis() *redigomock.Conn {
	conn := redigomock.NewConn()

	redisPool = &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return conn, nil
		},
	}

	return conn
}

func MockExpectPing(conn *redigomock.Conn, err error) {
	cmd := conn.Command("PING")
	if err == nil {
		cmd.Expect("")
	} else {
		cmd.ExpectError(err)
	}
}

func MockExpectSet(conn *redigomock.Conn, key string, value string, time gtime.Duration, err error) {
	cmd := conn.Command("SET", key, value)
	if err != nil {
		cmd.ExpectError(err)
	} else {
		cmd.Expect("")
	}

	conn.Command("EXPIRE", key, time.Seconds()).Expect("")
}

func MockExpectSetStruct(conn *redigomock.Conn, key string, value interface{}, time gtime.Duration, err error) {
	buf, merr := msgpack.Marshal(value)
	if merr != nil {
		panic(merr)
	}

	cmd := conn.Command("SET", key, buf)
	if err != nil {
		cmd.ExpectError(err)
	} else {
		cmd.Expect("")
	}

	conn.Command("EXPIRE", key, time.Seconds()).Expect("")
}

func MockExpectGet(conn *redigomock.Conn, key string, value string, err error) {
	cmd := conn.Command("GET", key)
	if err != nil {
		cmd.ExpectError(err)
	} else {
		cmd.Expect(value)
	}
}

func MockExpectGetStruct(conn *redigomock.Conn, key string, value interface{}, err error) {
	cmd := conn.Command("GET", key)
	if err != nil {
		cmd.ExpectError(err)
		return
	}

	buf, merr := msgpack.Marshal(value)
	if merr != nil {
		panic(merr)
	}

	cmd.Expect(buf)
}
