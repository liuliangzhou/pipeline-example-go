package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const webContent = "Hello World!"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	fmt.Fprint(w, webContent, t)
}
