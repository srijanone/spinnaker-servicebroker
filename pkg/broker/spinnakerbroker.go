package broker

// Runs at startup and bootstraps the broker.
func NewSpinnakerBroker(o Options) (*SpinnakerBroker, error) {
	bl := SpinnakerBroker{
		async:     o.Async,
		GateUrl:   o.GateUrl,
		instances: make(map[string]*ServiceInstance, 10),
	}
	return &bl, nil
}
