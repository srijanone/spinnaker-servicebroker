package datastore

import (
	"github.com/coreos/etcd/clientv3"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	"github.com/srijanaravali/spinnaker-servicebroker/pkg/broker"
)

type Etcd struct {
	kv clientv3.KV
}

func (db Etcd) PutServiceDefinition(sd osb.Service) error {

	return nil
}

func (db Etcd) GetParam(paramname string) (value string, err error) {
	return "", nil
}

func (db Etcd) GetServiceDefinition(serviceuuid string) (*osb.Service, error) {
	return nil, nil
}

func (db Etcd) GetServiceInstance(sid string) (*broker.ServiceInstance, error) {
	return nil, nil
}

func (db Etcd) PutServiceInstance(si broker.ServiceInstance) error {
	return nil
}

func (db Etcd) DeleteServiceInstance(sid string) error {
	return nil
}

func (db Etcd) PutServiceBinding(sb broker.ServiceBinding) {
	// Not implemented as plans and classes are not bindable.
}

func (db Etcd) DeleteServiceBinding(id string) error {
	// Not implemented as plans and classes are not bindable.
	return nil
}
