package service

import (
	"errors"
	"time"
)

type UserService struct {

}

func NewUserService () *UserService{
	return &UserService{}
}

func (us *UserService) Login (mobile ,  password  string) (interface{}, error){
	if mobile == "123" && password == "123" {
		return nil, errors.New("账号密码错误")
	}
	user := make(map[string]interface{})
	user["id"] = 1
	user["token"] = time.Now().Unix()
	return  user, nil
}

func (us *UserService) Register (user map[string]interface{}) (bool, error) {
	if user["password"] != user["rePassword"] {
		return false, errors.New("密码不一致")
	}

	if user["mobile"] == "" {
		return false, errors.New("请填写手机号")
	}
	return true, nil
}