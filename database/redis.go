package database

import (
	"Dream/conf"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var RedisDB redis.Conn

func init() {
	ip := conf.Config.Redis.IP
	port := conf.Config.Redis.Port
	password := conf.Config.Redis.Password
	var err error
	RedisDB, err = redis.Dial("tcp", ip+":"+port)
	RedisDB.Do("AUTH", password)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect redis")
	}
	if RedisDB.Err() != nil {
		panic("redis error")
	}
}
