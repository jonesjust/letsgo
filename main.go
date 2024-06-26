package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    ":4000",
		Handler: http.HandlerFunc(basicHandler),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("something went wrong", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
