package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web/splive/dist/spl/")))	
	http.ListenAndServe(":5200", nil)
}