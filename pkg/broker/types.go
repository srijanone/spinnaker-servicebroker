package broker

import (
	"sync"

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
	instances map[string]*service.ServiceInstance
}

type pipeline struct {
	Schema          string `json:"schema"`
	Template        `json:"template"`
	Application     string        `json:"application"`
	Name            string        `json:"name"`
	Triggers        []interface{} `json:"triggers"`
	Type            string        `json:"type"`
	Stages          []interface{} `json:"stages"`
	Variables       `json:"variables"`
	Exclude         []interface{} `json:"exclude"`
	ParameterConfig []interface{} `json:"parameterConfig"`
	Notifications   []interface{} `json:"notifications"`
}

type Template struct {
	ArtifactAccount string `json:"artifactAccount"`
	Reference       string `json:"reference"`
	Type            string `json:"type"`
}

type Variables struct {
	Namespace                    string `json:"namespace"`
	DockerRegistry               string `json:"docker_registry"`
	K8SAccount                   string `json:"k8s_account"`
	HelmPackageS3ObjectPath      string `json:"helm_package_s3_object_path"`
	HelmOverrideFileS3ObjectPath string `json:"helm_override_file_s3_object_path"`
	DockerRegistryOrg            string `json:"docker_registry_org"`
	DockerRepository             string `json:"docker_repository"`
	HalS3Account                 string `json:"hal_s3_account"`
	HalDockerRegistryAccount     string `json:"hal_docker_registry_account"`
	DockerImageTag               string `json:"docker_image_tag"`
	SpinnakerApplication         string `json:"spinnaker_application"`
}

//@TODO: Needs rename.
type requestBodyDelete struct {
	Application  string `json:"application"`
	PipelineName string `json:"pipelineName"`
}
