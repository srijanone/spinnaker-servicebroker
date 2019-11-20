package service

import (
	"encoding/json"
	"reflect"
)

type ServiceInstance struct {
	ID        string
	ServiceID string
	PlanID    string
	Params    map[string]interface{}
}

type ServiceBinding struct {
	ID         string
	InstanceID string
	PolicyArn  string
	RoleName   string
	Scope      string
}

// NewInstanceFromJSON returns a new Instance unmarshalled from the provided
// JSON []byte.
func NewInstanceFromJSON(jsonBytes []byte) (ServiceInstance, error) {
	instance := ServiceInstance{}
	err := json.Unmarshal(jsonBytes, &instance)
	return instance, err
}

// ToJSON returns a []byte containing a JSON representation of the
// instance.
func (i ServiceInstance) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}

func (i *ServiceInstance) Match(other *ServiceInstance) bool {
	return reflect.DeepEqual(i, other)
}
