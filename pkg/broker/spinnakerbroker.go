package broker

import "github.com/srijanaravali/spinnaker-servicebroker/pkg/service"

// Runs at startup and bootstraps the broker.
func NewSpinnakerBroker(o Options) (*SpinnakerBroker, error) {
	bl := SpinnakerBroker{
		async:     o.Async,
		GateUrl:   o.GateUrl,
		instances: make(map[string]*service.ServiceInstance, 10),
	}
	return &bl, nil
}
