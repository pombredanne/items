package DAO


func InsertWaterMelon(ip string,cid int,actionType int) (int, error) {
	var id int
	DB.SQL("insert into watermelon(ip,cid,actionType) values(?,?,?) returning id",ip,cid,actionType ).Get(&id)
	return id, nil
}