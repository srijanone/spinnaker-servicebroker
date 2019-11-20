package broker

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/google/uuid"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	"github.com/pmorie/osb-broker-lib/pkg/broker"
	spinnaker "github.com/srijanaravali/spinnaker-servicebroker/pkg/pipeline"
	"github.com/srijanaravali/spinnaker-servicebroker/pkg/service"
)

func (b *SpinnakerBroker) GetCatalog(c *broker.RequestContext) (*broker.CatalogResponse, error) {

	restEndpoint := b.GateUrl + "/v2/pipelineTemplates"

	resp, err := http.Get(restEndpoint)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var pipelineTemplates []interface{}

	json.Unmarshal([]byte(body), &pipelineTemplates)
	var plans []osb.Plan
	set := make(map[string]bool)
	for _, pipelineTemplate := range pipelineTemplates {
		data := pipelineTemplate.(map[string]interface{})
		name := data["id"].(string)
		metadata := data["metadata"].(map[string]interface{})
		if !set[name] {
			plan := osb.Plan{
				Name:        name,
				ID:          uuid.New().String(),
				Description: metadata["description"].(string),
				Free:        truePtr(),
				Bindable:    falsePtr(),
			}
			plans = append(plans, plan)
		}
		set[name] = true
	}

	response := &broker.CatalogResponse{}
	osbResponse := &osb.CatalogResponse{
		Services: []osb.Service{
			{
				Name:          "spinnaker-pipeline-as-service",
				ID:            uuid.New().String(),
				Description:   "Spinnaker Pipeline as Service.",
				Bindable:      false,
				PlanUpdatable: truePtr(),
				Plans:         plans,
			},
		},
	}

	glog.Infof("catalog response: %#+v", osbResponse)

	response.CatalogResponse = *osbResponse

	return response, nil
}

// Provision is executed when the OSB API receives `PUT /v2/service_instances/:instance_id`
func (b *SpinnakerBroker) Provision(request *osb.ProvisionRequest, c *broker.RequestContext) (*broker.ProvisionResponse, error) {

	restEndpoint := b.GateUrl + "/pipelines"

	response := broker.ProvisionResponse{}

	serviceInstance := &service.ServiceInstance{
		ID:        request.InstanceID,
		ServiceID: request.ServiceID,
		PlanID:    request.PlanID,
		Params:    request.Parameters,
	}

	params := request.Parameters

	pipeline, err := spinnaker.NewSpinnakerPipeline(params)
	if err != nil {
		log.Fatalln(err)
	}

	if spinnaker.CreatePipeline(restEndpoint, pipeline) {
		b.storage.WriteInstance(*serviceInstance)
	}

	// Check to see if this is the same instance.
	// @TODO: Needs fix. Need to get persistence.
	if i := b.instances[request.InstanceID]; i != nil {
		if i.Match(serviceInstance) {
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
	b.instances[request.InstanceID] = serviceInstance
	if request.AcceptsIncomplete {
		response.Async = b.async
	}

	return &response, nil
}

func (b *SpinnakerBroker) Deprovision(request *osb.DeprovisionRequest, c *broker.RequestContext) (*broker.DeprovisionResponse, error) {

	serviceInstance, _, err := b.storage.GetInstance(request.InstanceID)

	if err != nil {
		log.Fatalln(err)
	}

	applicatioName := serviceInstance.Params["spinnaker_application"].(string)
	pipelineName := serviceInstance.Params["pipeline_name"].(string)
	// @TODO: This is test code. Needs to be deleted.
	restEndpoint := b.GateUrl + "/pipelines/" + applicatioName + "/" + pipelineName
	requestBody := &spinnaker.DeletePayload{
		Application:  applicatioName,
		PipelineName: pipelineName,
	}
	response := broker.DeprovisionResponse{}

	if spinnaker.DeletePipeline(restEndpoint, requestBody) {
		b.storage.DeleteInstance(request.InstanceID)
	}
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
