package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	handler := http.FileServer(http.Dir("./static"))
	log.Println("Listening on port", port)
	http.ListenAndServe(port, handler)
}
