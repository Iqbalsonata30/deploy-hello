package main

import (
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World\n"))
	})
	log.Println("server running on port :3000")
	http.ListenAndServe(":3000", r)

}
