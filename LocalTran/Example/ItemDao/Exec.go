package ItemDao

import "fmt"

type ItemDaoI interface {
	Insert(i int)
}

type ItemDaoImpl struct {
	Count int
}

func Get() ItemDaoI{
	return ItemDaoImpl{}
}

func (ud ItemDaoImpl) Insert(i int) {
	fmt.Println(fmt.Sprintf("插入%d条Item成功", i))
	ud.CountAdd(i)
}
func (ud *ItemDaoImpl) CountAdd(i int){
	ud.Count+=i
}