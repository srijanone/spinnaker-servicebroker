# Spinnaker Pipeline as Service.

## Introduction
Provide Spinnaker Pipelines as service using OSBAPI. Spinnaker Pipeline Broker(ClusterServiceBroker) offers Pipelines
as Service(ClusterServiceClass) and each of spinnaker pipeline template correspond to a ClusterServicePlan.

Refer: https://kubernetes.io/docs/concepts/extend-kubernetes/service-catalog/ for terminologies used above.

## Naming Conventions
Pipelines templates(plans) are named according to following conventions. Each of the pipeline templates correspond to *ClusterServicePlan*

```
<provider>-<details>
```
* The *provider* is the name of the provider(example: k8s, cloud-foundary, aws-ec2, aws-ecs).
* The *details* is brief details of the pipeline + storage service used(example: bake-deploy-s3, bake-approve-deploy-s3). This
  should provide insights into the stages involved in the pipelines.

## Configurations/Installations

Pipeline templates are not configured in the default installation of Spinnaker and has to be enabled which can be done
as below.

```
hal config features edit --pipeline-templates true
hal deploy apply
```

Templates can be managed using Spin or the UI. To manage templates through the UI, enable the requisite feature flag:
```
hal config features edit --managed-pipeline-templates-v2-ui true
```

## References

https://www.spinnaker.io/guides/user/pipeline/pipeline-templates/