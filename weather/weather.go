package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	r := Record{
		Station: "Babylon 5",
		Time:    time.Now(),       // field type of time
		Temp:    Value{22.3, "c"}, // compound type
	}

	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	if err := enc.Encode(r); err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

// REF: https://golang.org/pkg/encoding/json/#RawMessage.MarshalJSON
func (v Value) MarshalJSON() ([]byte, error) {
	// Step 1: Convert to a JSONable type
	value := fmt.Sprintf("%f%s", v.Amount, v.Unit)
	// Step 2: Use encoding/json to marshal it
	return json.Marshal(value)
}

type Value struct {
	Amount float64
	Unit   string
}

type Record struct {
	Station string
	Time    time.Time
	Temp    Value
}
