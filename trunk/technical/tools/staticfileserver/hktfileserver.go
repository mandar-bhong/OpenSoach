package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web/splive/dist/hkt/")))	
	http.ListenAndServe(":5201", nil)
}