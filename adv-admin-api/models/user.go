package models

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/garyburd/redigo/redis"
	"github.com/dgrijalva/jwt-go"
	"time"
	"api-libs/my-jwt"
	"corm"
	"adv-admin-api/config"
	"api-libs/rsp"
)

type User struct {
	Id        *int    `db:"id"`
	Email     *string `json:"email" db:"email" validate:"required,gt=0,regexp=email"`
	Password  *string `json:"password" db:"password" validate:"required,gte=6,lte=20"`
	Role      *int    `json:"role" db:"role"`
	IpAddress *string `json:"ip_address" db:"ip_address"`
	Status    *int    `json:"status" db:"status"`
	RegDate   *string `json:"reg_date" db:"reg_date"`
	Test      *bool   `json:"test" db:"test"`
}

// 更改当前token 版本
func UpdateTokenVersion(userId interface{}) error {
	client := config.Conf.Redis.GetPool().Get()
	if _, err := client.Do("SET", fmt.Sprintf("login_user:%v", userId), uuid.NewV4().String()); err != nil {
		return rsp.HandlerError(err)
	}
	defer client.Close()
	return nil
}

func GetTokenVersion(userId interface{}) (string, error) {
	client := config.Conf.Redis.GetPool().Get()
	version, _ := redis.String(client.Do("GET", fmt.Sprintf("login_user:%v", userId)))
	if version == "" {
		version = uuid.NewV4().String()
		if _, err := client.Do("SET", fmt.Sprintf("login_user:%v", userId), version); err != nil {
			return "", rsp.HandlerError(err)
		}
	}
	defer client.Close()
	return version, nil
}

func CreateToken(userId interface{}, version, signingKey string) (string, error) {
	claims := jwt.MapClaims{
		"id":      userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"version": version,
	}
	j := my_jwt.New(signingKey)
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", rsp.HandlerError(err)
	}
	return token, nil
}

// 验证用户是否存在
func CheckUserExist(email string, roleId int, result *User) error {
	if err := corm.Db.Get(result, `select * from user_user where
	email=$1 AND status!=0 AND role=$2`, email, roleId); err != nil {
		return rsp.HandlerError(err)
	}
	return nil
}
