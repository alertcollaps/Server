package main

import (
	"Server/httpservice"
	"Server/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(sql.GetLastElement())
	httpservice.NewHandler()
	log.Println("Started!!!")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln("ListenServe", err)
	}

}
