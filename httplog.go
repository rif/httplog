package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		b, _ := json.MarshalIndent(r, "", " ")
		log.Print(string(b))
	})
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	log.Printf("Listening on 0.0.0.0:%s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
