package xormTool

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	id,err:=Update("update class set name=? where id=? returning id","高中n班",1)
	if err!=nil {
		t.Fatal(err)
	}
	t.Log(id)
}
