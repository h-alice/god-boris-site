package main

import (
	"html/template"
	"log"
	"net/http"
)

func clientHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("/static/html/page_template.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := template.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", clientHandler)
	log.Println("[+] Collecting God Boris Power...")
	log.Println("[+] Server Starting Up.")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
