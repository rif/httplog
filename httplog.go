package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		b, _ := json.MarshalIndent(r, "", " ")
		log.Print(string(b))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
