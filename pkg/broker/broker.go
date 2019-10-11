package broker

import (
	"fmt"
)

// Runs at startup and bootstraps the broker.
func NewSpinnakerBroker(o Options) (*SpinnakerBroker, error) {
	fmt.Print("I am here")
	bl := SpinnakerBroker{
		gateUrl: "localhost:8084",
	}
	return &bl, nil
}
