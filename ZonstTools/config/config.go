package config

import (
	"strings"
	"encoding/json"
	"ZonstTools/common/database/my-redis"
	"ZonstTools/common/database/my-db"
)

type Config struct {
	Addr  string         `json:"addr"`
	Redis my_redis.Redis `json:"redis"`
	DB    my_db.Database `json:"db"`
}

var configMap = map[string]map[string]interface{}{
	"develop": {
		"addr": ":1303",
		"redis": map[string]interface{}{
			"address": "",
			"maxIdle": 20,
			"maxConn": 10,
		},
		"db": map[string]interface{}{
			"user":     "fe",
			"password": "fe7S@QF4cLVwuLBR",
			"host":     "111.231.137.127",
			"port":     5432,
			"dbName":   "test_fe",
			"maxIdle":  20,
			"maxConn":  10,
		},
	},
	"testing": {
		"addr": ":1403",
		"redis": map[string]interface{}{
			"address": "redis://127.0.0.1:6379",
			"maxIdle": 20,
			"maxConn": 10,
		},
		"db": map[string]interface{}{
			"user":     "fe",
			"password": "fe7S@QF4cLVwuLBR",
			"host":     "111.231.137.127",
			"port":     5432,
			"dbName":   "test_fe",
			"maxIdle":  20,
			"maxConn":  10,
		},
	},
	"production": {
		"addr": ":1303",
		"redis": map[string]interface{}{
			"address": "redis://127.0.0.1:6379",
			"maxIdle": 20,
			"maxConn": 10,
		},
		"db": map[string]interface{}{
			"user":     "fe",
			"password": "fe7S@QF4cLVwuLBR",
			"host":     "111.231.137.127",
			"port":     5432,
			"dbName":   "fe",
			"maxIdle":  20,
			"maxConn":  10,
		},
	},
}

func Load(envModel string) *Config {
	config := &Config{}
	configMap[envModel]["envModel"] = envModel
	//将enModel对应的addr,db,redis配置序列化，再反序列注入config体种
	configJson, _ := json.Marshal(configMap[envModel])
	jsonParser := json.NewDecoder(strings.NewReader(string(configJson)))
	jsonParser.Decode(config)
	return config
}
