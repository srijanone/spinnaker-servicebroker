package broker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/glog"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	"github.com/pmorie/osb-broker-lib/pkg/broker"
)

func (b *SpinnakerBroker) GetCatalog(c *broker.RequestContext) (*broker.CatalogResponse, error) {

	// @TODO: Persist this in some sort of storage.
	response := &broker.CatalogResponse{}
	osbResponse := &osb.CatalogResponse{
		Services: []osb.Service{
			{
				Name:          "spinnaker-pipeline-as-service",
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

// Provision is executed when the OSB API receives `PUT /v2/service_instances/:instance_id`
func (b *SpinnakerBroker) Provision(request *osb.ProvisionRequest, c *broker.RequestContext) (*broker.ProvisionResponse, error) {

	url := b.GateUrl + "pipelines"

	response := broker.ProvisionResponse{}

	serviceInstance := &serviceInstance{
		ID:        request.InstanceID,
		ServiceID: request.ServiceID,
		PlanID:    request.PlanID,
		Params:    request.Parameters,
	}

	PlanID := "k8s-bake-deploy-s3"

	switch PlanID {
	case "k8s-bake-deploy-s3":
		pipeline := &pipeline{
			Schema: "v2",
			Template: Template{
				ArtifactAccount: "front50ArtifactCredentials",
				Reference:       "spinnaker://k8s-bake-approve-deploy-s3-23-oct",
				Type:            "front50/pipelineTemplate",
			},
			Application: "v2poc",
			Name:        "test-11",
			Type:        "templatedPipeline",
			Triggers:    make([]interface{}, 0),
			Stages:      make([]interface{}, 0),
			Variables: Variables{
				Namespace:                    "default",
				DockerRegistry:               "docker.io",
				K8SAccount:                   "my-k8-account",
				HelmPackageS3ObjectPath:      "s3://spin-helm/node-1.0.0.tgz",
				HelmOverrideFileS3ObjectPath: "s3://spin-helm/values.yaml",
				DockerRegistryOrg:            "athakur",
				DockerRepository:             "athakur/node",
				HalS3Account:                 "my-s3-account",
				HalDockerRegistryAccount:     "my-docker-registry",
				DockerImageTag:               "0.1.0",
				SpinnakerApplication:         "v2poc",
			},
			Exclude:         make([]interface{}, 0),
			ParameterConfig: make([]interface{}, 0),
			Notifications:   make([]interface{}, 0),
		}
		requestBody, _ := json.Marshal(pipeline)
		fmt.Println(requestBody)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(body))
	case "k8s-bake-approve-deploy-s3":
		// Logic goes here.

	}
	// Check to see if this is the same instance
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
