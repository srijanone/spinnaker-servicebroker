# Spinnaker Service Broker

## Introduction
Provide Spinnaker Pipelines as service using OSBAPI. Spinnaker Pipeline Broker(ClusterServiceBroker) offers Pipelines
as Service(ClusterServiceClass) and each of spinnaker pipeline template correspond to a ClusterServicePlan.

Refer: https://kubernetes.io/docs/concepts/extend-kubernetes/service-catalog/ for terminologies used above.

## Installation
The current release assumes that spinnaker and spinnaker-servicebroker should be installed in the same K8S cluster(shared services cluster) and Spinnaker Servicebroker would be able to communicate to spinnaker(gate).

Spinnaker ServiceBroker ships with helm charts to install it on a K8S cluster.

### Kubernetes Service Catalog
Install Kubernetes [Service Catalog](https://kubernetes.io/docs/tasks/service-catalog/install-service-catalog-using-helm/). *This is a pre-requisite*

### Spinnaker
It is assumed that Spinnaker is installed on the same cluster using steps mentioned in https://www.spinnaker.io/setup/. As of now we only provide pipeline templates to deploy workloads on K8S cluster.

### Spinnaker ServiceBroker

  * Change default GateUr in values.yaml. In the current release it is assumed that Spinnaker and Spinnaker Service Broker run in the
    same cluster and Gate is accessible via K8S service url.
  ```
  spinnaker:
    gate_url: http://spin-gate.spinnaker.svc.cluster.local:8084
  ```
  * Install spinnaker-servicebroker
```
helm install charts/spinnaker-servicebroker --name spinnakar-servicebroker --namespace <namespace>
```
  In case spinnaker-servicebroker is installed correctly, you should see following output
  ```
  #$ kubectl get ClusterServiceBroker
  NAME                      URL                                                                                                STATUS   AGE
  spinnaker-servicebroker   http://spinnaker-servicebroker-spinnaker-servicebroker.spinnaker-servicebroker.svc.cluster.local   Ready    10d
  ```
  In case you have [svcat](https://svc-cat.io/docs/install/) installed you should see the service offerings.
  ```
  #:~ Ashish_Thakur$ svcat marketplace
              CLASS                         PLANS                       DESCRIPTION
  +-------------------------------+----------------------------+--------------------------------+
  spinnaker-pipeline-as-service   k8s-bake-approve-deploy-s3   Spinnaker Pipeline as Service.
                                  k8s-bake-deploy-s3
  ```
## Pipeline templates
This service provides spinnaker pipeline templates as a service offering(service class). The supported templates and example manifest files can be found at ```templates``` and ```examples``` directory. Currently we provision the pipeline templates using *spin-cli* or *Gate API*. We would like to provision these templates automatically when the broker is deployed.