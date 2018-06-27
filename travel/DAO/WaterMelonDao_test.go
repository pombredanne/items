package DAO

import "testing"

func TestInsertWaterMelon(t *testing.T) {
	id,err:=InsertWaterMelon("10.0.203.92",1,2)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(id)
}
