package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
   fmt.Println(r.PostForm)
	_, _ = fmt.Fprintln(w, r.Form)
}
func main() {
	http.HandleFunc("/process", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
