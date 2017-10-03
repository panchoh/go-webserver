package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	counter int
	port    = 9000
)

func count(w http.ResponseWriter, r *http.Request) {
	counter++
}

func stats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Visits: %d", counter)))
}

func main() {
	http.HandleFunc("/", count)
	http.HandleFunc("/stats", stats)

	log.Println("Listening at port ", port)
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
