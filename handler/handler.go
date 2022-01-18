package handler

import (
	"go-web/entity"
	"html/template"
	"log"
	"net/http"
	"path"
)

func HelloHandler(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(writer, req)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(writer, "Error Goblok!!!",http.StatusInternalServerError)
		return
	}
	// data := entity.Product{ID:1, Name:"Book", Price:10000, Stock:5}
	data := []entity.Product{
		{ID:1, Name:"Book", Price:7000, Stock:1},
		{ID:2, Name:"Pencil", Price:5000, Stock:45},
		{ID:3, Name:"Ruler", Price:15000, Stock:25},
		{ID:4, Name:"Eraser", Price:1000, Stock:55},
	}
	err = tmpl.Execute(writer, data)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Error Goblok!!!",http.StatusInternalServerError)
		return
	}
}
func AboutHandler(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/about" {
		http.NotFound(writer, req)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "about.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(writer, "Error Goblok!!!",http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"title": "Go Web ",
		"content": "ABout Page MADYA WEB GOLANG",
	}
	tmpl.Execute(writer, data)
}

func MethodHandler(writer http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/testMethod" {
		http.NotFound(writer, req)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "methodChecker.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(writer, "Error Goblok!!!",http.StatusInternalServerError)
		return
	}
	var m string
	switch req.Method{
	case "GET": 
		m = "GET"
		break
	case "POST":
		m = "POST"
		break
	default:
		m = "Method Lain"
	}
	data := map[string]interface{}{
		"title": "Go Web ",
		"content": m,
	}
	tmpl.Execute(writer, data)
}

func FormHandler(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/form" {
		http.NotFound(writer, req)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(writer, "Error Goblok!!!",http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"title": "Go Web ",
	}
	tmpl.Execute(writer, data)
}

func ProcessHandler(writer http.ResponseWriter, req *http.Request){
	if req.Method == "POST"{
		err := req.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(writer, "Error Goblok!!!",http.StatusInternalServerError)
			return
		}
		name := req.Form.Get("name");
		age := req.Form.Get("age");

		writer.Write([]byte(name))
		writer.Write([]byte(age))
	}
}