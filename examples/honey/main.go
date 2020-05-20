package main

import (
	"github.com/daymenu/honey"
	"log"
)

const address = ":8000"

func main () {
	app := honey.New()
	app.GET("/hello", Hello)
	app.Run(address)
}

func Hello(c *honey.Context) {
	c.String("honey is happy")
	log.Println("honey is happy")
}