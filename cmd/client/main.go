package main

import (
	"corpuser/client/router"
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Println("...")
	log.Println("Starting the rest server")
	r := router.GetRouter()
	log.Fatal(http.ListenAndServe(":8080", removeTrailingSlash(r)))
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
