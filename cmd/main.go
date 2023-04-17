package main

import (
	"ascii-art-web-export-file/handlers"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/ascii-art", handlers.ProcessorHandler)
	mux.Handle("/ui/static/", http.StripPrefix("/ui/static/", http.FileServer(http.Dir("./ui/static/"))))
	fmt.Println("Server launched at http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("this host is already run")
	}
}
