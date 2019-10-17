package broker

// Runs at startup and bootstraps the broker.
func NewSpinnakerBroker(o Options) (*SpinnakerBroker, error) {
	bl := SpinnakerBroker{
		async:     o.Async,
		gateUrl:   "localhost:8084",
		instances: make(map[string]*exampleInstance, 10),
	}
	return &bl, nil
}
