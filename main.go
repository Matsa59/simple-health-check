package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func logReq(ip string) {
	msg := fmt.Sprintf("Request receive from %s.", ip)
	log.Println(msg)
}

func main() {
	log.Printf("Start listening (port: %d)\n.", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go logReq(r.RemoteAddr)
		fmt.Fprintf(w, "OK")
	})

	portString := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(portString, nil))
}
