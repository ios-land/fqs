package main

import (
	"fmt"
	"log"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func main() {
	http.HandleFunc("/ping", pingHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Fatal(err)
	}

}
