package main

import "fmt"

type DataBaseSlice struct{
	Id int
	Name string
	Money float64
}

type MemorySession struct{
}
func (ms MemorySession) Update(){

}


func main() {
	var datas = make([]DataBaseSlice,0)
	datas = GetSlice()
	go func(){
		SynToDatabase(datas)
	}()

	Update(datas[0],100)

	Update(datas[1],200)

	defer SynToDatabase(datas)
}

//模仿从数据库获取数据
func GetSlice() []DataBaseSlice{
	return []DataBaseSlice{
		DataBaseSlice{4,"ft4",1},
		DataBaseSlice{1,"ft1",2},
	}
}

//模仿同步数据库数据
func SynToDatabase(datas []DataBaseSlice){
	fmt.Println("成功同步")
}

//模仿加钱
func Update(data DataBaseSlice ,off float64){
	data.Money+=off
}