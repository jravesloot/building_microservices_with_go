// ex3.go uses an Encoder to write json to the writer directly, without
// first storing all json in a byte array, which is much more efficient

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
	// change output key to be camel case 'message'
	Message string `json:"message"`
	// "-" means field is not marshalled
	Author string `json:"-"`
	// doesn't marshal if field is empty; note the , (fieldname,directive)
	Date string `json:",omitempty"`
	// convert output to a string and rename the field
	Id int `json:"id, string"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{
		Message: "Coming to directly from the encoder",
		Id:      77,
		Author:  "Jimbo",
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
