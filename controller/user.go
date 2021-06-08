package controller

import (
	"encoding/json"
	"go-im/service"
	"net/http"
)

type UserController struct {

}

func NewUserController () *UserController{
	return &UserController{}
}

func (uc *UserController) Login(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	password := request.PostForm.Get("password")
	userService := service.NewUserService()
	user, err := userService.Login(mobile, password)
	ret := make(map[string]interface{})
	if err != nil {
		ret["code"] = http.StatusInternalServerError
		ret["msg"] = err.Error()
		ret["data"] = ""
	}
	writer.Header().Set("Content-Type", "application/json")
	ret["code"] = http.StatusOK
	ret["msg"] = ""
	ret["data"] = user
	msg , _ := json.Marshal(ret)
	writer.Write(msg)

}

func (uc *UserController) Register (writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	password := request.PostForm.Get("password")
	rePassword := request.PostForm.Get("re_password")
	nickname := request.PostForm.Get("nickname")
	sex := request.PostForm.Get("sex")

	param := make(map[string]interface{})
	param["mobile"] = mobile
	param["password"] = password
	param["rePassword"] = rePassword
	param["nickname"] = nickname
	param["sex"] = sex

	userService := service.NewUserService()
	userService.Register(param)

}