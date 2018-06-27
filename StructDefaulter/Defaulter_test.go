package StructDefaulter

import (
	"testing"
	"fmt"
)

func TestDefaulter_Init(t *testing.T) {
	type User struct{
		Name string `Defaulter:"ft"`
	}
	user :=User{"undefined"}
	defaulter :=Defaulter{}
	defaulter.Init(user)
	fmt.Println(user)
}