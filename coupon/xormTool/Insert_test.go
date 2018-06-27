package xormTool

import (
	"testing"

)

func TestInsert(t *testing.T) {
	sql := "insert into class(name) values(?) "
	id,err:=Insert(sql,"高中6班")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(id)
}
func TestLocalSessionInsert(t *testing.T) {
	sql := "insert into class(name) values(?) returning id"
	id,err:=LocalSessionInsert(sql,"高中6班")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(id)
}
