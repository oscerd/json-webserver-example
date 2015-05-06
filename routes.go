package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Result struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl/index.html")
	t.Execute(w, "pippo")
	fmt.Fprintf(w, "") // send data to client side
}

func HandleJSON(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("Nome")
	cognome := r.FormValue("Cognome")
	result, _ := json.Marshal(Result{string(nome), string(cognome)})
	t, _ := template.ParseFiles("tmpl/json.html")
	t.Execute(w, string(result))
	fmt.Fprintf(w, "") // send data to client side
}

func HandleInverseJSON(w http.ResponseWriter, r *http.Request) {
	data := []byte(string(r.FormValue("json")))
	var pi Result
	_ = json.Unmarshal(data, &pi)
	t, _ := template.ParseFiles("tmpl/decode.html")
	t.Execute(w, pi)
	fmt.Fprintf(w, "") // send data to client side
}
