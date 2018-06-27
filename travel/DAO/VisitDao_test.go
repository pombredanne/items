package DAO

import (
	"testing"
	"travel/models/vo"
)

func TestSaveVisitor(t *testing.T) {
	visitorVO := vo.Visitor{"127.0.0.1","3"}
	id,err:=SaveVisitor(visitorVO)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(id)
}