package main

import (
	"encoding/json"
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

type helloWorldResponse struct {
	Message string
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello za warudo"}
	data, err := json.Marshal(response)

	if err != nil {
		panic("Oops!")
	}

	fmt.Fprint(w, string(data))
}
