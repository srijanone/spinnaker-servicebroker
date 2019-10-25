package broker

// Construct a pipeline object.
func NewSpinnakerPipeline(params map[string]interface{}) (*pipeline, error) {

	pipeline := &pipeline{
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
