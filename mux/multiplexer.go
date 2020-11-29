package mux

import (
	"fmt"
	"log"
	"net/http"
)

type CustomServerMux struct {
}

func (p *CustomServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		log.Println("using q")
		fmt.Fprintf(w, "test")
		return
	}
	fmt.Fprint(w, "404 not found")
}
