package broker

import (
	"net/http"
	"reflect"

	"github.com/golang/glog"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	"github.com/pmorie/osb-broker-lib/pkg/broker"
)

func truePtr() *bool {
	b := true
	return &b
}

func falsePtr() *bool {
	b := false
	return &b
}

func (b *SpinnakerBroker) GetCatalog(c *broker.RequestContext) (*broker.CatalogResponse, error) {
	// Your catalog business logic goes here.
	response := &broker.CatalogResponse{}
	osbResponse := &osb.CatalogResponse{
		Services: []osb.Service{
			{
				Name:          "Spinnaker Pipeline as Service.",
				ID:            "4f6e6cf6-ffdd-425f-a2c7-3c9258ad246a",
				Description:   "Spinnaker Pipeline as Service.",
				Bindable:      false,
				PlanUpdatable: truePtr(),
				Plans: []osb.Plan{
					{
						Name:        "k8s-bake-approve-deploy-s3",
						ID:          "86064792-7ea2-467b-af93-ac9694d96d5b",
						Description: "Pipeline template for K8S(Manifest Based) provider using highlander strategy and S3 as artifact storage.",
						Free:        truePtr(),
						Bindable:    falsePtr(),
					},
					{
						Name:        "k8s-bake-deploy-s3",
						ID:          "86064792-7ea2-467b-af93-bc9694d96d5b",
						Description: "Pipeline template for K8S(Manifest Based) provider using highlander strategy and S3 as artifact storage.",
						Free:        truePtr(),
						Bindable:    falsePtr(),
					},
				},
			},
		},
	}

	glog.Infof("catalog response: %#+v", osbResponse)

	response.CatalogResponse = *osbResponse

	return response, nil
}

func (b *SpinnakerBroker) Provision(request *osb.ProvisionRequest, c *broker.RequestContext) (*broker.ProvisionResponse, error) {
	// Your provision business logic goes here

	// Send request to spinnaker to create the pipeline.
	// example implementation:
	b.Lock()
	defer b.Unlock()

	response := broker.ProvisionResponse{}

	exampleInstance := &exampleInstance{
		ID:        request.InstanceID,
		ServiceID: request.ServiceID,
		PlanID:    request.PlanID,
		Params:    request.Parameters,
	}

	// Check to see if this is the same instance
	if i := b.instances[request.InstanceID]; i != nil {
		if i.Match(exampleInstance) {
			response.Exists = true
			return &response, nil
		} else {
			// Instance ID in use, this is a conflict.
			description := "InstanceID in use"
			return nil, osb.HTTPStatusCodeError{
				StatusCode:  http.StatusConflict,
				Description: &description,
			}
		}
	}
	b.instances[request.InstanceID] = exampleInstance

	if request.AcceptsIncomplete {
		response.Async = b.async
	}

	return &response, nil
}

func (b *SpinnakerBroker) Deprovision(request *osb.DeprovisionRequest, c *broker.RequestContext) (*broker.DeprovisionResponse, error) {
	// Your deprovision business logic goes here

	// example implementation:
	b.Lock()
	defer b.Unlock()

	response := broker.DeprovisionResponse{}

	delete(b.instances, request.InstanceID)

	if request.AcceptsIncomplete {
		response.Async = b.async
	}

	return &response, nil
}

func (b *SpinnakerBroker) LastOperation(request *osb.LastOperationRequest, c *broker.RequestContext) (*broker.LastOperationResponse, error) {
	// Your last-operation business logic goes here

	return nil, nil
}

// Not used as Services and ServicePlans are non-bindable.
func (b *SpinnakerBroker) Bind(request *osb.BindRequest, c *broker.RequestContext) (*broker.BindResponse, error) {
	// Your bind business logic goes here

	// example implementation:
	b.Lock()
	defer b.Unlock()

	instance, ok := b.instances[request.InstanceID]
	if !ok {
		return nil, osb.HTTPStatusCodeError{
			StatusCode: http.StatusNotFound,
		}
	}

	response := broker.BindResponse{
		BindResponse: osb.BindResponse{
			Credentials: instance.Params,
		},
	}
	if request.AcceptsIncomplete {
		response.Async = b.async
	}

	return &response, nil
}

// Not used as Services and ServicePlans are non-bindable.
func (b *SpinnakerBroker) Unbind(request *osb.UnbindRequest, c *broker.RequestContext) (*broker.UnbindResponse, error) {
	// Your unbind business logic goes here
	return &broker.UnbindResponse{}, nil
}

func (b *SpinnakerBroker) Update(request *osb.UpdateInstanceRequest, c *broker.RequestContext) (*broker.UpdateInstanceResponse, error) {
	// Your logic for updating a service goes here.
	response := broker.UpdateInstanceResponse{}
	if request.AcceptsIncomplete {
		response.Async = b.async
	}

	return &response, nil
}

func (b *SpinnakerBroker) ValidateBrokerAPIVersion(version string) error {
	return nil
}

func (i *exampleInstance) Match(other *exampleInstance) bool {
	return reflect.DeepEqual(i, other)
}
