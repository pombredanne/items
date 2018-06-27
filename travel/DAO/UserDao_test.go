package DAO

import (
	"testing"
	"travel/models/vo"
	)


func TestSaveUser(t *testing.T) {
	userVO :=vo.User{"ft","139382974738","江西省","南昌市","中至大厦","2"}
	id,err:=SaveUser(userVO)
	if err!=nil {
		t.Fatal(err)
	}
	t.Log(id)
}


func TestGetUserByDate(t *testing.T) {
	users,err:=GetUserByDate("2018-4-2","2018-4-2")
	if err!=nil {
		t.Fatal(err)
	}
	t.Log(users)
}