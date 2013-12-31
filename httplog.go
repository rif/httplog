package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	address = flag.String("address", "0.0.0.0", "listening interface address")
	port    = flag.String("port", "8080", "listening port")
	format  = flag.Bool("format", false, "pretty print json request info")
	echo    = flag.Bool("echo", false, "send request info back to client")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var b []byte
		if *format {
			b, _ = json.MarshalIndent(r, "", " ")
		} else {
			b, _ = json.Marshal(r)
		}
		s := string(b)
		log.Print(s)
		if *echo {
			fmt.Fprint(w, s)
		}
	})
	log.Printf("Listening on %s:%s...", *address, *port)
	log.Fatal(http.ListenAndServe(*address+":"+*port, nil))
}
