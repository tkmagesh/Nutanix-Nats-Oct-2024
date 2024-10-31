package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	// Make sure that we have the NATS Go client imported.
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"

	"golang.org/x/exp/slices"
)

func main() {
	// Determine a suitable URL for a connection to a NATS server
	url, exists := os.LookupEnv("NATS_URL")
	if !exists {
		url = nats.DefaultURL
	} else {
		url = strings.TrimSpace(url)
	}

	if strings.TrimSpace(url) == "" {
		url = nats.DefaultURL
	}

	// Connect to the server
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	// ### Defining a Service
	//
	// This will create a service definition. Service definitions are made up of
	// the service name (which can't have things like whitespace in it), a version,
	// and a description. Even with no running endpoints, this service is discoverable
	// via the micro protocol and by service discovery tools like `nats micro`.
	// All of the default background handlers for discovery, PING, and stats are
	// started at this point.
	srv, err := micro.AddService(nc, micro.Config{
		Name:        "minmax",
		Version:     "0.0.1",
		Description: "Returns the min/max number in a request",
		QueueGroup:  "minmax-qg",
	})

	// Each time we create a service, it will be given a new unique identifier. If multiple
	// copies of the `minmax` service are running across a NATS subject space, then tools
	// like `nats micro` will consider them like unique instances of the one service and the
	// endpoint subscriptions are queue subscribed, so requests will only be sent to one
	// endpoint _instance_ at a time.
	fmt.Printf("Created service: %s (%s)\n", srv.Info().Name, srv.Info().ID)

	if err != nil {
		log.Fatal(err)
		return
	}

	// ### Adding endpoints
	//
	// Groups serve as namespaces and are used as a subject prefix when endpoints
	// don't supply fixed subjects. In this case, all endpoints will be listening
	// on a subject that starts with `minmax.`
	root := srv.AddGroup("minmax")

	// Adds two endpoints to the service, one for the `min` operation and one for
	// the `max` operation. Each endpoint represents a subscription. The supplied handlers
	// will respond to `minmax.min` and `minmax.max`, respectively.
	root.AddEndpoint("min", micro.HandlerFunc(handleMin))
	root.AddEndpoint("max", micro.HandlerFunc(handleMax))

	fmt.Printf("Endpoint '%s' requests: %d\n", srv.Stats().Endpoints[0].Name, srv.Stats().Endpoints[0].NumRequests)
	fmt.Printf("Endpoint '%s' requests: %d\n", srv.Stats().Endpoints[1].Name, srv.Stats().Endpoints[1].NumRequests)

	runtime.Goexit()
}

func handleMin(req micro.Request) {
	var arr MinMaxRequest
	_ = json.Unmarshal([]byte(req.Data()), &arr)
	slices.Sort(arr.Values)

	res := ServiceResult{Min: arr.Values[0]}
	req.RespondJSON(res)
}

func handleMax(req micro.Request) {
	var arr MinMaxRequest
	_ = json.Unmarshal([]byte(req.Data()), &arr)
	slices.Sort(arr.Values)

	res := ServiceResult{Max: arr.Values[len(arr.Values)-1]}
	req.RespondJSON(res)
}

func decode(msg *nats.Msg) ServiceResult {
	var res ServiceResult
	json.Unmarshal(msg.Data, &res)
	return res
}

type MinMaxRequest struct {
	Values []int `json:"values"`
}

type ServiceResult struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}
