package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type MinMaxRequest struct {
	Values []int `json:"values"`
}

func main() {

	// Connect to the server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
		return
	}

	// requestSlice := []int{-1, 2, 100, -2000}
	// Creating a raw set of bytes to send to the service endpoints
	req := MinMaxRequest{
		Values: []int{-1, 2, 100, -2000},
	}
	for range 500 {
		time.Sleep(500 * time.Millisecond)
		requestData, _ := json.Marshal(req)
		// fmt.Println(string(requestData))
		// Make a request of the `min` endpoint of the `minmax` service, within the `minmax` group.
		// Note that there's nothing special about this request, it's just a regular NATS
		// request.
		msg, _ := nc.Request("minmax.min", requestData, 2*time.Second)
		// Decode is just a convenience method that unmarshals the JSON response into the
		// `ServiceResult` type.
		result := decode(msg)
		fmt.Printf("Requested min value, got %d\n", result.Min)

		// Make a request of the `max` endpoint of the `minmax` service, within the `minmax` group.
		msg, _ = nc.Request("minmax.max", requestData, 2*time.Second)
		result = decode(msg)
		fmt.Printf("Requested max value, got %d\n", result.Max)
	}
}

func decode(msg *nats.Msg) ServiceResult {
	var res ServiceResult
	json.Unmarshal(msg.Data, &res)
	return res
}

type ServiceResult struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}
