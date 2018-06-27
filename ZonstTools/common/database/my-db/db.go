package my_db

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"fmt"
)

//连接数据库信息
type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Host     string `json:"host"`
	DbName   string `json:"dbName"`
	MaxIdle  int    `json:"maxIdle"`
	MaxConn  int    `json:"maxConn"`
}

var db *sqlx.DB

//使用时需要提前注入好Database数据
func (d *Database) NewDb() {
	var err error
	db, err = sqlx.Open("postgres", d.ConnectString())
	if err != nil {
		panic(err)
	}
	// set max open connections
	if d.MaxConn!=0{
		db.SetMaxOpenConns(d.MaxConn)
	}
	if d.MaxIdle!=0{
		db.SetMaxIdleConns(d.MaxIdle)
	}
	// set max idle connections

	// try connecting to the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

}
func (d *Database) ConnectString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Password, d.DbName)
}

func CloseDb() (err error) {
	return db.Close()
}

func GetDb() *sqlx.DB {
	return db.Unsafe()
}
