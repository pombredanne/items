package mydb

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"time"
	"log"
)

//func GetPg(url string) *dbr.Session {
//	conn, err := dbr.Open("postgres", url, nil)
//	if err != nil {
//		panic(err)
//	}
//	return conn.NewSession(nil)
//}

func GetPgx(url string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		panic(err)
	}
	return db.Unsafe()

}

func GetRedis(url string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 200,
		//MaxActive:   0,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(url)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}


