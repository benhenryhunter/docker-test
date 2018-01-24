package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	router := NewRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 300 * time.Second,
		ReadTimeout:  300 * time.Second,
	}

	fmt.Println("Listening on port: 8080")

	log.Fatal(srv.ListenAndServe())
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
