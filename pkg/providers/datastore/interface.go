package datastore

import (
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	"github.com/srijanaravali/spinnaker-servicebroker/pkg/broker"
)

type Interface interface {

	// PutServiceDefinition push catalog service definition to datastore.
	PutServiceDefinition(sd osb.Service) error

	// GetParam fetches parameter from datastore.
	GetParam(paramname string) (value string, err error)

	// PutParam puts parameters into datastore.
	PutParam(paramname string, paramvalue string) error

	// GetServiceDefinition fetches given catalog service definition from datastore.
	GetServiceDefinition(serviceuuid string) (*osb.Service, error)

	// GetServiceInstance fetches given service instance from datastore.
	GetServiceInstance(sid string) (*broker.ServiceInstance, error)

	// PutServiceInstance stores given service instance in datastore.
	PutServiceInstance(si broker.ServiceInstance) error

	// DeleteServiceInstance deletes the service instance.
	DeleteServiceInstance(sid string) error

	// PutServiceBinding stores the service binding.
	PutServiceBinding(sb broker.ServiceBinding)

	// DeleteServiceBinding deletes the service binding.
	DeleteServiceBinding(id string) error
}
