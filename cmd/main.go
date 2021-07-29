package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	srv := &http.Server{
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		Addr:         ":8777",
		Handler:      Routes(),
	}

	fmt.Printf("Running on port 8777, friend.")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
