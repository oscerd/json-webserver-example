package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/jsonencode", HandleJSON)
	http.HandleFunc("/jsondecode", HandleInverseJSON)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
