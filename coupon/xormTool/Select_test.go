package xormTool

import (
	"testing"
	"time"
)

func TestSelectCount(t *testing.T) {
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	count, err := SelectCount("select * from class")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
}

func TestSelect(t *testing.T) {
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	type Class struct {
		Id   int
		Name string
	}
	classes := make([]Class, 0)
	err := Select(&classes, "select * from class")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(classes)

}
func TestSelectOne(t *testing.T) {
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	type Class struct {
		Id   int
		Name string
	}
	class := Class{}
	err := SelectOne(&class, "select * from class where id=?", 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(class)
}

func TestDynamicSelect(t *testing.T) {
	type User struct {
		Id          int
		Name        string
		ClassId     int `db:"class_id"`
		Description string
	}
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	users := make([]User, 0)
	orderBy := make([]string, 0)
	orderBy = append(orderBy, "id")
	whereMap := make([][]string, 0)
	whereMap = append(whereMap, []string{
		"", "name", "=",
	})
	whereMap = append(whereMap, []string{
		"and", "id", ">",
	})
	whereMap = append(whereMap, []string{
		"and", "id", "<>",
	})
	whereMap = append(whereMap, []string{
		"and", "created", ">",
	})
	err := DynamicSelect(&users, "select * from public.user", whereMap, orderBy, "desc", 2, 0, "ft4", "1035", "1","2012/09/09")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(users)
}

func TestDynamicSelectOne(t *testing.T) {
	type User struct {
		Id          int
		Name        string
		ClassId     int `db:"class_id"`
		Description string
	}
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	user := User{}
	orderBy := make([]string, 0)
	orderBy = append(orderBy, "id")
	whereMap := make([][]string, 0)
	whereMap = append(whereMap, []string{
		"", "id", "=",
	})
	err := DynamicSelectOne(&user, "select * from public.user", whereMap, nil, "desc", 2, 0, 1036)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestDynamicSelectCount(t *testing.T) {
	type User struct {
		Id          int
		Name        string
		ClassId     int `db:"class_id"`
		Description string
	}
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	orderBy := make([]string, 0)
	orderBy = append(orderBy, "id")
	whereMap := make([][]string, 0)
	whereMap = append(whereMap, []string{
		"", "name", "=",
	})
	whereMap = append(whereMap, []string{
		"and", "id", ">",
	})
	whereMap = append(whereMap, []string{
		"and", "id", "<>",
	})
	count, err := DynamicSelectCount("select * from public.user", whereMap, nil, "desc", 2, 0, "ft4", 1035, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
}

func TestRemove(t *testing.T) {
	ti, er := time.Parse("2006-01-02 15:04:05", "2018-04-13 15:30:33")
	if er != nil {
		t.Fatal(er)
	}
	args := []interface{}{
		1, ti,1,
	}
	t.Log("原始数据:",args)
	args = remove(args, 1)
	t.Log("删除后:",args)
}
func BenchmarkRemove(b *testing.B){
	ti, er := time.Parse("2006-01-02 15:04:05", "2018-04-13 15:30:33")
	if er != nil {
		b.Fatal(er)
	}
	args := []interface{}{
		1, ti,1,
	}
	b.Log("原始数据:",args)
	for i:=0;i<b.N;i++{
		args = remove(args, 1)
	}
	b.Log("删除后:",args)
}
func TestRemoveZero(t *testing.T) {
	ti, er := time.Parse("2006-01-02 15:04:05", "2018-04-13 15:30:33")
	if er != nil {
		t.Fatal(er)
	}
	args := []interface{}{
		1, 0.5, "test", ti,"","%%",
	}
	args=removeZero(args)
	t.Log(args)
}

func BenchmarkRemoveZero(b *testing.B){
	ti, er := time.Parse("2006-01-02 15:04:05", "2018-04-13 15:30:33")
	args := []interface{}{
		1, 0.5, "test", ti,"","%%",
	}
	if er != nil {
		b.Fatal(er)
	}
	for i:=0;i<b.N;i++{
		args=removeZero(args)
	}
	b.Log(args)
}