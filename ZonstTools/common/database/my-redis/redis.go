package my_redis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type Redis struct {
	//Protocol string `json:"protocol"`
	Address string `json:"address"`
	MaxIdle int    `json:"maxIdle"`
	MaxConn int    `json:"maxConn"`
}

//Pool不直接对外，可以通过GetConn()拿到单个connection
var redisPool *redis.Pool

// creates a new pool
//使用前提前注入好Redis数据
func (r *Redis) NewRedis() {
	redisPool = &redis.Pool{
		MaxIdle:     r.MaxIdle,
		MaxActive:   r.MaxConn,
		IdleTimeout: 240 * time.Second,
		Dial: func() (c redis.Conn, err error) {
			c, err = redis.DialURL(r.Address)
			if err != nil {
				return
			}
			return
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return
}

func GetConn() redis.Conn {
	return redisPool.Get()
}
