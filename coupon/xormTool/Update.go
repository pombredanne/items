package xormTool

func Update(sql string,args ...interface{})(int,error){
	var id int

	session :=Db.NewSession()
	session.Begin()
	defer session.Close()
	_,err:=session.SQL(sql,args...).Get(&id)
	if err!=nil {
		session.Rollback()
		return -1,err
	}
	session.Commit()
	return id,nil
}
