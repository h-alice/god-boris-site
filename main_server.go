package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Test server. Hello, World!")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Println("[+] Collecting God Boris Power...")
	log.Println("[+] Server Starting Up.")
	// fmt.Println("[+] Server Starting Up... Collecting God Boris Power.")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}