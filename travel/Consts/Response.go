package Consts

import "travel/models/vo"

var Secret = "Travel"
var InfoLacking = vo.FeedBack{
	Status:-1,
	Msg:"请输入足够的信息！",
}
var OKInsert = vo.FeedBack{
	Status:1,
	Msg:"插入成功",
}
var ResponseTokenSetWrong=vo.FeedBack{
	Status:-1,
	Msg:"未找到token",
}

var ResponseTokenValidateError=vo.FeedBack{
	Status:-1,
	Msg:"Token验证error",
}
var ResponseTokenValidateWrong=vo.FeedBack{
	Status:-1,
	Msg:"Token验证错误",
}
var ResponseTokenGenerateError=vo.FeedBack{
	Status:-1,
	Msg:"Token验证错误",
}
var ResponseWrongAcount = vo.FeedBack{
	Status:-1,
	Msg:"账户错误",
}
var ResponseTokenNotFound=vo.FeedBack{
	Status:-1,
	Msg:"未找到合法token令牌",
}