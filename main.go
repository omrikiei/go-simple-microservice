package main

import (
	"log"
	"net/http"

	// load and register all the controllers
	_ "github.com/omrikiei/go-simple-microservice/controllers/monitor"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
