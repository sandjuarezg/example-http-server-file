package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var fs http.Handler = http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	fmt.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
