package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	// nil uses default handler DefaultServeMux, configured above
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
