package config

import (
	"strings"
	"encoding/json"
	"api-libs/database/my-redis"
	"api-libs/database/my-db"
)

var (
	Conf *config
)

type config struct {
	Addr       string         `json:"addr"`
	Redis      my_redis.Redis `json:"redis"`
	DB         my_db.Database `json:"db"`
	EnvModel   string         `json:"envModel"`
	SigningKey string         `json:"signingKey"`
	RoleId     int            `json:"roleId"`
}

var configMap = map[string]map[string]interface{}{
	"develop": {
		"addr": ":1304",
		"redis": map[string]interface{}{
			"address": "redis://127.0.0.1:6379",
			"maxIdle": 10,
			"maxConn": 0,
		},
		"db": map[string]interface{}{
			"user":     "fe",
			"password": "fe7S@QF4cLVwuLBR",
			"host":     "111.231.137.127",
			"port":     5432,
			"dbName":   "test_fe",
		},
		"signingKey": "bestcaicai444444",
		"roleId":     4,
	},
	"testing": {
		"addr": ":1404",
		"redis": map[string]interface{}{
			"address": "redis://127.0.0.1:6379",
			"maxIdle": 10,
			"maxConn": 0,
		},
		"db": map[string]interface{}{
			"user":     "fe",
			"password": "fe7S@QF4cLVwuLBR",
			"host":     "111.231.137.127",
			"port":     5432,
			"dbName":   "test_fe",

		},
		"signingKey": "bestcaicai444444",
		"roleId":     4,
	},
	"production": {
		"addr": ":1304",
		"redis": map[string]interface{}{
			"address": "redis://127.0.0.1:6379",
			"maxIdle": 10,
			"maxConn": 0,

		},
		"db": map[string]interface{}{
			"user":     "fe",
			"password": "fe7S@QF4cLVwuLBR",
			"host":     "127.0.0.1",
			"port":     5432,
			"dbName":   "fe",
			"maxIdle":  10,
			"maxConn":  0,
		},
		"signingKey": "bestcaicai444444",
		"roleId":     4,
	},
}

func Init(envModel string) {
	Conf = &config{}
	configMap[envModel]["envModel"] = envModel
	configJson, _ := json.Marshal(configMap[envModel])
	jsonParser := json.NewDecoder(strings.NewReader(string(configJson)))
	jsonParser.Decode(Conf)
}
