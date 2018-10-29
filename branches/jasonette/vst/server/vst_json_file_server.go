package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./json_files")))
	err := http.ListenAndServe(":81", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
