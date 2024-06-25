package main

import (
	"fmt"
	"groupie-tracker/server"
	"log"
	"net/http"
)

const port = ":6060"

func main() {

	fmt.Println("(http://localhost:6060) - Server started on port", port)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./frontend/assets"))))
	http.HandleFunc("/", server.PathHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}
