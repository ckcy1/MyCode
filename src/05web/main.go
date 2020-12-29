package main

import (
	"html/template"
	"net/http"
)

func  handle(w http.ResponseWriter,r *http.Request)  {
	tem,_:=template.ParseFiles("index.html")
	tem.Execute(w, "hello")
}
func main()  {
	http.HandleFunc("/",handle)
	http.ListenAndServe(":8080", nil)

}