package datastore

import (
	"github.com/srijanaravali/spinnaker-servicebroker/pkg/service"
)

// Store is an interface to be implemented by types capable of handling
// persistence for other broker-related types
type DataStore interface {

	// WriteInstance persists the given instance to the underlying storage
	WriteInstance(instance service.ServiceInstance) error

	// GetInstance retrieves a persisted instance from the underlying storage by
	// instance id
	GetInstance(instanceID string) (service.ServiceInstance, bool, error)

	// GetInstanceByID retrieves a persisted instance from the underlying storage
	// by alias
	// GetInstanceByAlias(alias string) (broker.ServiceInstance, bool, error)

	// GetInstanceChildCountByAlias returns the number of child instances
	// GetInstanceChildCountByAlias(alias string) (int64, error)

	// DeleteInstance deletes a persisted instance from the underlying storage by
	// instance id
	DeleteInstance(instanceID string) (bool, error)

	// WriteBinding persists the given binding to the underlying storage
	// WriteBinding(binding broker.ServiceBinding) error

	// GetBinding retrieves a persisted instance from the underlying storage by
	// binding id
	// GetBinding(bindingID string) (broker.ServiceBinding, bool, error)

	// DeleteBinding deletes a persisted binding from the underlying storage by
	// binding id
	// DeleteBinding(bindingID string) (bool, error)

	// TestConnection tests the connection to the underlying database (if there
	// is one)
	TestConnection() error
}
