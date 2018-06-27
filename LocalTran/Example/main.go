package main

import (
	"LocalTran"
	"LocalTran/Example/UserDao"
	"LocalTran/Example/ItemDao"
)
var userDao UserDao.UserDaoI
var itemDao ItemDao.ItemDaoI
func main() {
	tranManager:=LocalTran.GetTransactionManager()
	userDao = UserDao.Get()
	itemDao = ItemDao.Get()
	r1 :=LocalTran.RegistedObject{userDao.Insert,4}
	r2 :=LocalTran.RegistedObject{itemDao.Insert,3}
	tranManager.BeginTran(r1,r2)
}