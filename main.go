package main

import (

	"go-im/controller"
	"html/template"
	"log"
	"net/http"
)

// 加载全部模板
func LoadTemplate () {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err)
	}
	for _, v:= range tpl.Templates() {
		tName := v.Name()
		http.HandleFunc(tName, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tName, nil)
		})
	}
}

func main () {

	user := controller.NewUserController()
	http.HandleFunc("/login", user.Login)
	http.HandleFunc("/register", user.Register)

	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	LoadTemplate()

	http.ListenAndServe(":8122", nil)
}
