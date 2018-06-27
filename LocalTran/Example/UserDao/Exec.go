package UserDao

import "fmt"

type UserDaoI interface {
	Insert(i int)
}

type UserDaoImpl struct {
	Count int
}

func Get() UserDaoI{
	return UserDaoImpl{}
}

func (ud UserDaoImpl) Insert(i int) {
	fmt.Println(fmt.Sprintf("插入%d条User成功", i))
	ud.CountAdd(3)
}
func (ud *UserDaoImpl) CountAdd(i int){
	ud.Count=ud.Count+i
}