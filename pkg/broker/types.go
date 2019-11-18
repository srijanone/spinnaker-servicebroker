package broker

import (
	"sync"

	"github.com/srijanaravali/spinnaker-servicebroker/pkg/datastore"
	"github.com/srijanaravali/spinnaker-servicebroker/pkg/service"
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
	storage   datastore.DataStore
	instances map[string]*service.ServiceInstance
}
