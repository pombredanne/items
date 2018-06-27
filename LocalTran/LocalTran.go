package LocalTran
type RegistedObject struct{
	F func(args...interface{})
	Args interface{}
}

type TransactionManagerI interface{
	Status() bool				//返回事务自动管理是否开启

}
type TransactionManager struct{
	status bool
}

func GetTransactionManager()*TransactionManager{
	return &TransactionManager{true}
}

func (tm *TransactionManager) BeginTran(registedObjects ...RegistedObject){

}

func (tm *TransactionManager) Status() bool{
	return tm.status
}