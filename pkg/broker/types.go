package broker

import (
	"sync"
)

type Options struct {
	GateUrl     string
	CatalogPath string
	Async       bool
}

type SpinnakerBroker struct {
	async bool
	sync.RWMutex
	GateUrl   string
	instances map[string]*exampleInstance
}

// exampleInstance is intended as an example of a type that holds information about a service instance
type exampleInstance struct {
	ID        string
	ServiceID string
	PlanID    string
	Params    map[string]interface{}
}
