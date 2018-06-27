package  main

import (
//	"test_toml/test"
	"fmt"
	"github.com/BurntSushi/toml"
)

//订制配置文件解析载体
type Config struct{
	Database *Database
	SQL *SQL
}

//订制Database块
type Database struct {
	Driver    string
	Username  string `toml:"us"` //表示该属性对应toml里的us
	Password string
}
//订制SQL语句结构
type SQL struct{
	Sql1 string `toml:"sql_1"`
	Sql2 string `toml:"sql_2"`
	Sql3 string `toml:"sql_3"`
	Sql4 string `toml:"sql_4"`
}

var config *Config=new (Config)
func init(){
	//读取配置文件
	_, err := toml.DecodeFile("test.toml",config)
	if err!=nil{
		fmt.Println(err)
	}
}
func main() {
	  fmt.Println(config)
	  fmt.Println(config.Database.Password)
	  fmt.Println(config.SQL.Sql1)
}