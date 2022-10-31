package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/task/", hello)

	handler := Logging(mux)

	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}
