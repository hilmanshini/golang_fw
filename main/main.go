package main

import (
	"app/middleware"
	"fmt"
	"log"
	"net/http"
)

type CustomServerMux struct {
}

func (p *CustomServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("SERVING")
	middleware.MustAuth(func(h1 http.ResponseWriter, h2 *http.Request) {
		fmt.Fprintf(h1, "OK")
	})(w, r)
}

func main() {
	appmux := &CustomServerMux{}
	http.ListenAndServe(":8000", appmux)

}
