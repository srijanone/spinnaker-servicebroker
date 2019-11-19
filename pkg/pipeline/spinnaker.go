package pipeline

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Construct a pipeline object.
func NewSpinnakerPipeline(params map[string]interface{}) (*SpinnakerPipeline, error) {

	pipeline := &SpinnakerPipeline{
		Schema: "v2",
		Template: Template{
			ArtifactAccount: "front50ArtifactCredentials",
			Reference:       "spinnaker://" + params["pipeline_template"].(string),
			Type:            "front50/pipelineTemplate",
		},
		Application: params["spinnaker_application"].(string),
		Name:        params["pipeline_name"].(string),
		Type:        "templatedPipeline",
		Triggers:    make([]interface{}, 0),
		Stages:      make([]interface{}, 0),
		Variables: Variables{
			Namespace:                    params["namespace"].(string),
			DockerRegistry:               params["docker_registry"].(string),
			K8SAccount:                   params["k8s_account"].(string),
			HelmPackageS3ObjectPath:      params["helm_package_s3_object_path"].(string),
			HelmOverrideFileS3ObjectPath: params["helm_override_file_s3_object_path"].(string),
			DockerRegistryOrg:            params["docker_registry_org"].(string),
			DockerRepository:             params["docker_repository"].(string),
			HalS3Account:                 params["hal_s3_account"].(string),
			HalDockerRegistryAccount:     params["hal_docker_registry_account"].(string),
			DockerImageTag:               params["docker_image_tag"].(string),
			SpinnakerApplication:         params["spinnaker_application"].(string),
		},
		Exclude:         make([]interface{}, 0),
		ParameterConfig: make([]interface{}, 0),
		Notifications:   make([]interface{}, 0),
	}

	return pipeline, nil
}

func CreatePipeline(restEndpoint string, pipeline *SpinnakerPipeline) bool {
	requestBody, _ := json.Marshal(pipeline)
	resp, err := http.Post(restEndpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	log.Println(string(body))
	return true
}

func DeletePipeline(restEndpoint string, payload *DeletePayload) bool {
	requestBody, _ := json.Marshal(payload)
	req, err := http.NewRequest("DELETE", restEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	return true
}
