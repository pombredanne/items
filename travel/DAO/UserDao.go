package DAO

import (
	"travel/models/vo"
)

type User struct {
	Name     string
	Province string
	City     string
	Address  string
	Phone    string
	Status   int
	Tdate string
}

func SaveUser(user vo.User) (int, error) {
	var id int
	DB.SQL("insert into travelUser(name,phone,province,city,address,cid) values(?,?,?,?,?,?) returning id", user.Name, user.Phone, user.Province, user.City, user.Address, user.CId).Get(&id)
	return id, nil
}

func GetUserByDate(startTime string, endTime string) ([]User, error) {
	users := make([]User, 0)
	err := DB.SQL("select distinct name,province,city,address,phone,status,tdate from travelUser where tdate between ? and ?", startTime, endTime).Find(&users)
	if err != nil {
		return nil, err
	}
	for i:=range users{
	  users[i].Tdate= users[i].Tdate[:10]
	}
	return users, nil
}
