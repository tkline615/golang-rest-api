package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(resp, "Up and running")
	})
	log.Println("Server listning on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
