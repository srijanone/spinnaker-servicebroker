# Introduction
Pipeline template for K8S(Manifest Based) provider using highlander strategy and S3 as artifact storage.

# Creating a pipeline

Steps to create a pipeline from the pipeline template.

Note: It is assumed that Application in Spinnaker is created for the steps below.

Refer this link which illustrates how to create pipeline from a pipeline template: https://www.spinnaker.io/guides/user/pipeline/pipeline-templates/instantiate/.

# Important resources

Please note that in order to use S3 as artifact storage an artifact account has to be created. Similarly 
an account for Docker Registry has to be created as well. Check the resources below.
* Configuring S3 as artificat storage: https://www.spinnaker.io/setup/artifacts/s3/
* Configuring Docker registry: https://www.spinnaker.io/setup/install/providers/docker-registry/