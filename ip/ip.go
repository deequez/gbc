package main

/* REST: JSON over HTTP
CRUD: Create, Retrieve, Update, Delete
C: POST
R: GET
U: PUT/PATCH
D: DELETE
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer resp.Body.Close()

	// var reply ipReply
	var reply struct { // anonymous struct. used when there's no need to contaminate the global namespace.
		// Origin string
		IP string `json:"origin"` //field tag structured per package. here we're specifying for the json package.
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&reply); err != nil {
		log.Fatalf("error: %s", err)
	}
	// fmt.Println("IP:", reply.Origin)
	fmt.Println("IP:", reply.IP)
}

// IPReply returns the origin IP address of our request.
// The type itself does not have to be exported.
// What's important is that the fields are exported.
// type ipReply struct {
// 	// encoding/json will only populate exported fields (strating with uppercase)
// 	Origin string
// }

// Working with JSON, the map[string]{}interface case
func getIPMap() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer resp.Body.Close()

	// High-level of how we get the data
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
	fmt.Println("Body:") //, resp.Body)
	// io.Copy(os.Stdout, resp.Body)

	// How will we work with the data?
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(data)
	var reply map[string]interface{}
	if err := json.Unmarshal(data, &reply); err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println("reply:", reply)
	ip := reply["origin"]
	fmt.Println("IP: ", ip)
}
