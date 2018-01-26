package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	port = "8080"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Up")
}

func main() {
	router := NewRouter()
	fmt.Println("Listening on port: " + port)
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 300 * time.Second,
		ReadTimeout:  300 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
