package main

import (
  "fmt"
  "log"
  "net/http"
	"time"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
    log.Println("received request")
    fmt.Fprintf(w, "Hello Docker!!")
  })



  log.Println("start server")
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())


}
