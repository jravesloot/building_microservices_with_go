// ex6.go adds a json Encoder to send encoded data on the fly

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

type helloWorldRequest struct {
	// more explicit, and faster, to specify 'name' vs letting Unmarshal
	// try case-insensitive matching
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`
	Date    string `json:",omitempty"`
	Id      int    `json:"id, string"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// nb. not calling Body.Close() as typically needed because ServeHTTP
	// handler automatically closes request stream
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{
		Message: fmt.Sprintf("Hello %s, I was a %T", request.Name, &request),
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
