package DAO

import (

	"travel/models/vo"
)

func SaveVisitor(visitorVO vo.Visitor) (int,error){
	var id int
	DB.SQL("insert into visitorLog(ip,cid) values(?,?) returning id",visitorVO.Ip,visitorVO.CId).Get(&id)
	return id,nil
}

