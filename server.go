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

	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Fatal(err)
	}

}
