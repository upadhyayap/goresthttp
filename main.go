package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello world"))
	})

	// scope of err is limited in if statement
	if err := http.ListenAndServe(":6061", nil); err != nil {
		panic(err)
	}
}
