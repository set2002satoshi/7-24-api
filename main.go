package main

import (
	"fmt"
	"net/http"

	"github.com/set2002satoshi/7-24-api/db"
)

func main() {
	db.DbConnect()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}