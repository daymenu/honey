package main

import (
	"log"
	"net/http"

	"github.com/daymenu/honey/router"
)

func main() {
	route := router.New()
	log.Fatal(http.ListenAndServe(":8000", route))
}
