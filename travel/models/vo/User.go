package vo

type User struct {
	Name     string `json:"name,omitempty" form:"name" `
	Phone    string `json:"phone,omitempty" form:"phone" `
	Province string `json:"province,omitempty" form:"province" `
	City     string `json:"city,omitempty" form:"city"`
	Address  string `json:"address,omitempty" form:"address" `
	CId      string `json:"cid,omitempty" form:"cid"`
}

type Login struct{
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
}
