package broker

import "fmt"

// Runs at startup and bootstraps the broker.
func NewSpinnakerBroker(o Options) (*SpinnakerBroker, error) {
	fmt.Println(o.GateUrl)
	bl := SpinnakerBroker{
		async:     o.Async,
		GateUrl:   o.GateUrl,
		instances: make(map[string]*serviceInstance, 10),
	}
	return &bl, nil
}
