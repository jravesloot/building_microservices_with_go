// ex5.go uses an Encoder to write json to the writer directly, without

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	// nb. not calling Body.Close() as typically needed because ServeHTTP
	// handler automatically closes request stream
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var request helloWorldRequest
	err = json.Unmarshal(body, &request)
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
