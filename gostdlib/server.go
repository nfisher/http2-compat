package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	srv := &http2.Server{}
	server := &http.Server{
		Addr:    "localhost:8000",
		Handler: h2c.NewHandler(mux, srv),
	}

	log.Println("listening on", server.Addr)
	server.ListenAndServe()
}
