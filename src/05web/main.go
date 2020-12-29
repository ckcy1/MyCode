package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "time"
)

func handler (w http.ResponseWriter,r *http.Request){
  tem,err:= template.ParseFiles("/http)/index.html")
  fmt.Println(err)
  log.Fatal(tem.Execute(w,"hello"))
}
func main(){
//mux:=http.NewServeMux()
http.HandleFunc("/",handler)
log.Fatal(http.ListenAndServe(":8080",nil))//wo


}