package honey

import (
	"log"
	"net/http"
	"testing"
)

const address = ":8000"

func TestMain(t *testing.T) {
	route := New()
	log.Printf("honey server is started at : %s\n", address)
	log.Fatal(http.ListenAndServe(address, route))
}

func Hello(c *Context) {
	c.writermem.WriteString("honey is happy\n")
	c.writermem.WriteString(c.Request.RequestURI)
	log.Println("honey is happy")
}
