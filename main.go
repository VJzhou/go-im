package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

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

	http.HandleFunc("/user/login", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		mobile := request.PostForm.Get("mobile")
		password := request.PostForm.Get("password")
		fmt.Println(mobile)
		fmt.Println(password)
	})

	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	LoadTemplate()

	http.ListenAndServe(":8122", nil)
}
