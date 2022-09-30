package main

import (
	"Server/httpservice"
	"log"
	"net/http"
)

func main() {
	httpservice.NewHandler()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("ListenServe", err)
	}

}
