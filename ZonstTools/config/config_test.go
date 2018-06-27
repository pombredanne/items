package config

import "testing"

func TestLoad(t *testing.T) {
	config:=Load("develop")
	t.Log(config.Redis,config.DB,config.Addr)
}