# Package
- [autoscaling.k8s.io/v1](#autoscaling.k8s.io/v1)
- [autoscaling.k8s.io/v1beta1](#autoscaling.k8s.io/v1beta1)
- [autoscaling.k8s.io/v1beta2](#autoscaling.k8s.io/v1beta2)
- [poc.autoscaling.k8s.io/v1alpha1](#poc.autoscaling.k8s.io/v1alpha1)
## autoscaling.k8s.io/v1 package display name
> <p>Package v1 contains definitions of Vertical Pod Autoscaler related objects.</p>
### Resource Types:
- [VerticalPodAutoscaler](#autoscaling.k8s.io/v1.VerticalPodAutoscaler)
- [VerticalPodAutoscalerCheckpoint](#autoscaling.k8s.io/v1.VerticalPodAutoscalerCheckpoint)
### VerticalPodAutoscaler
<a id="autoscaling.k8s.io/v1.VerticalPodAutoscaler"></a>
> <p>VerticalPodAutoscaler is the configuration for a vertical pod
autoscaler, which automatically manages pod resources based on historical and
real time resource utilization.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `apiVersion` | `autoscaling.k8s.io/v1` |
| `kind`        | `VerticalPodAutoscaler` |
| `metadata` | [Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#objectmeta-v1-meta) |
| `spec` | [VerticalPodAutoscalerSpec](#autoscaling.k8s.io/v1.VerticalPodAutoscalerSpec) |
| `status` | [VerticalPodAutoscalerStatus](#autoscaling.k8s.io/v1.VerticalPodAutoscalerStatus) |
### VerticalPodAutoscalerCheckpoint
<a id="autoscaling.k8s.io/v1.VerticalPodAutoscalerCheckpoint"></a>
> <p>VerticalPodAutoscalerCheckpoint is the checkpoint of the internal state of VPA that
is used for recovery after recommender&rsquo;s restart.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `apiVersion` | `autoscaling.k8s.io/v1` |
| `kind`        | `VerticalPodAutoscalerCheckpoint` |
| `metadata` | [Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#objectmeta-v1-meta) |
| `spec` | [VerticalPodAutoscalerCheckpointSpec](#autoscaling.k8s.io/v1.VerticalPodAutoscalerCheckpointSpec) |
| `status` | [VerticalPodAutoscalerCheckpointStatus](#autoscaling.k8s.io/v1.VerticalPodAutoscalerCheckpointStatus) |
### ContainerControlledValues
<a id="autoscaling.k8s.io/v1.ContainerControlledValues"></a>
(*Alias*: `string`)
(*Appears on*: 
[ContainerResourcePolicy](#autoscaling.k8s.io/v1.ContainerResourcePolicy)
)
> <p>ContainerControlledValues controls which resource value should be autoscaled.</p>
#### Constants
| Value-Type             | Description         |
|--------------------|---------------------|
| `&#34;RequestsAndLimits&#34;` | <p>ContainerControlledValuesRequestsAndLimits means resource request and limits
are scaled automatically. The limit is scaled proportionally to the request.</p>
|
| `&#34;RequestsOnly&#34;` | <p>ContainerControlledValuesRequestsOnly means only requested resource is autoscaled.</p>
|
### ContainerResourcePolicy
<a id="autoscaling.k8s.io/v1.ContainerResourcePolicy"></a>
(*Appears on*: 
[PodResourcePolicy](#autoscaling.k8s.io/v1.PodResourcePolicy)
)
> <p>ContainerResourcePolicy controls how autoscaler computes the recommended
resources for a specific container.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerName` | string
|
| `mode` | [ContainerScalingMode](#autoscaling.k8s.io/v1.ContainerScalingMode) |
| `minAllowed` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `maxAllowed` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `controlledResources` | [[]Kubernetes core/v1.ResourceName](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcename-v1-core) |
| `controlledValues` | [ContainerControlledValues](#autoscaling.k8s.io/v1.ContainerControlledValues) |
### ContainerScalingMode
<a id="autoscaling.k8s.io/v1.ContainerScalingMode"></a>
(*Alias*: `string`)
(*Appears on*: 
[ContainerResourcePolicy](#autoscaling.k8s.io/v1.ContainerResourcePolicy)
)
> <p>ContainerScalingMode controls whether autoscaler is enabled for a specific
container.</p>
#### Constants
| Value-Type             | Description         |
|--------------------|---------------------|
| `&#34;Auto&#34;` | <p>ContainerScalingModeAuto means autoscaling is enabled for a container.</p>
|
| `&#34;Off&#34;` | <p>ContainerScalingModeOff means autoscaling is disabled for a container.</p>
|
### EvictionChangeRequirement
<a id="autoscaling.k8s.io/v1.EvictionChangeRequirement"></a>
(*Alias*: `string`)
(*Appears on*: 
[EvictionRequirement](#autoscaling.k8s.io/v1.EvictionRequirement)
)
> <p>EvictionChangeRequirement refers to the relationship between the new target recommendation for a Pod and its current requests, what kind of change is necessary for the Pod to be evicted</p>
#### Constants
| Value-Type             | Description         |
|--------------------|---------------------|
| `&#34;TargetHigherThanRequests&#34;` | <p>TargetHigherThanRequests means the new target recommendation for a Pod is higher than its current requests, i.e. the Pod is scaled up</p>
|
| `&#34;TargetLowerThanRequests&#34;` | <p>TargetLowerThanRequests means the new target recommendation for a Pod is lower than its current requests, i.e. the Pod is scaled down</p>
|
### EvictionRequirement
<a id="autoscaling.k8s.io/v1.EvictionRequirement"></a>
(*Appears on*: 
[PodUpdatePolicy](#autoscaling.k8s.io/v1.PodUpdatePolicy)
)
> <p>EvictionRequirement defines a single condition which needs to be true in
order to evict a Pod</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `resources` | [[]Kubernetes core/v1.ResourceName](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcename-v1-core) |
| `changeRequirement` | [EvictionChangeRequirement](#autoscaling.k8s.io/v1.EvictionChangeRequirement) |
### HistogramCheckpoint
<a id="autoscaling.k8s.io/v1.HistogramCheckpoint"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpointStatus](#autoscaling.k8s.io/v1.VerticalPodAutoscalerCheckpointStatus)
)
> <p>HistogramCheckpoint contains data needed to reconstruct the histogram.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `referenceTimestamp` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `bucketWeights` | map[int]uint32
|
| `totalWeight` | float64
|
### PodResourcePolicy
<a id="autoscaling.k8s.io/v1.PodResourcePolicy"></a>
(*Appears on*: 
[VerticalPodAutoscalerSpec](#autoscaling.k8s.io/v1.VerticalPodAutoscalerSpec)
)
> <p>PodResourcePolicy controls how autoscaler computes the recommended resources
for containers belonging to the pod. There can be at most one entry for every
named container and optionally a single wildcard entry with <code>containerName</code> = &lsquo;*&rsquo;,
which handles all containers that don&rsquo;t have individual policies.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerPolicies` | [[]ContainerResourcePolicy](#autoscaling.k8s.io/v1.ContainerResourcePolicy) |
### PodUpdatePolicy
<a id="autoscaling.k8s.io/v1.PodUpdatePolicy"></a>
(*Appears on*: 
[VerticalPodAutoscalerSpec](#autoscaling.k8s.io/v1.VerticalPodAutoscalerSpec)
)
> <p>PodUpdatePolicy describes the rules on how changes are applied to the pods.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `updateMode` | [UpdateMode](#autoscaling.k8s.io/v1.UpdateMode) |
| `minReplicas` | int32
|
| `evictionRequirements` | [[]EvictionRequirement](#autoscaling.k8s.io/v1.EvictionRequirement) |
### RecommendedContainerResources
<a id="autoscaling.k8s.io/v1.RecommendedContainerResources"></a>
(*Appears on*: 
[RecommendedPodResources](#autoscaling.k8s.io/v1.RecommendedPodResources)
)
> <p>RecommendedContainerResources is the recommendation of resources computed by
autoscaler for a specific container. Respects the container resource policy
if present in the spec. In particular the recommendation is not produced for
containers with <code>ContainerScalingMode</code> set to &lsquo;Off&rsquo;.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerName` | string
|
| `target` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `lowerBound` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `upperBound` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `uncappedTarget` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
### RecommendedPodResources
<a id="autoscaling.k8s.io/v1.RecommendedPodResources"></a>
(*Appears on*: 
[VerticalPodAutoscalerStatus](#autoscaling.k8s.io/v1.VerticalPodAutoscalerStatus)
)
> <p>RecommendedPodResources is the recommendation of resources computed by
autoscaler. It contains a recommendation for each container in the pod
(except for those with <code>ContainerScalingMode</code> set to &lsquo;Off&rsquo;).</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerRecommendations` | [[]RecommendedContainerResources](#autoscaling.k8s.io/v1.RecommendedContainerResources) |
### UpdateMode
<a id="autoscaling.k8s.io/v1.UpdateMode"></a>
(*Alias*: `string`)
(*Appears on*: 
[PodUpdatePolicy](#autoscaling.k8s.io/v1.PodUpdatePolicy)
)
> <p>UpdateMode controls when autoscaler applies changes to the pod resources.</p>
#### Constants
| Value-Type             | Description         |
|--------------------|---------------------|
| `&#34;Auto&#34;` | <p>UpdateModeAuto means that autoscaler assigns resources on pod creation
and additionally can update them during the lifetime of the pod,
using any available update method. Currently this is equivalent to
Recreate, which is the only available update method.</p>
|
| `&#34;Initial&#34;` | <p>UpdateModeInitial means that autoscaler only assigns resources on pod
creation and does not change them during the lifetime of the pod.</p>
|
| `&#34;Off&#34;` | <p>UpdateModeOff means that autoscaler never changes Pod resources.
The recommender still sets the recommended resources in the
VerticalPodAutoscaler object. This can be used for a &ldquo;dry run&rdquo;.</p>
|
| `&#34;Recreate&#34;` | <p>UpdateModeRecreate means that autoscaler assigns resources on pod
creation and additionally can update them during the lifetime of the
pod by deleting and recreating the pod.</p>
|
### VerticalPodAutoscalerCheckpointSpec
<a id="autoscaling.k8s.io/v1.VerticalPodAutoscalerCheckpointSpec"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpoint](#autoscaling.k8s.io/v1.VerticalPodAutoscalerCheckpoint)
)
> <p>VerticalPodAutoscalerCheckpointSpec is the specification of the checkpoint object.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `vpaObjectName` | string
|
| `containerName` | string
|
### VerticalPodAutoscalerCheckpointStatus
<a id="autoscaling.k8s.io/v1.VerticalPodAutoscalerCheckpointStatus"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpoint](#autoscaling.k8s.io/v1.VerticalPodAutoscalerCheckpoint)
)
> <p>VerticalPodAutoscalerCheckpointStatus contains data of the checkpoint.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `lastUpdateTime` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `version` | string
|
| `cpuHistogram` | [HistogramCheckpoint](#autoscaling.k8s.io/v1.HistogramCheckpoint) |
| `memoryHistogram` | [HistogramCheckpoint](#autoscaling.k8s.io/v1.HistogramCheckpoint) |
| `firstSampleStart` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `lastSampleStart` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `totalSamplesCount` | int
|
### VerticalPodAutoscalerCondition
<a id="autoscaling.k8s.io/v1.VerticalPodAutoscalerCondition"></a>
(*Appears on*: 
[VerticalPodAutoscalerStatus](#autoscaling.k8s.io/v1.VerticalPodAutoscalerStatus)
)
> <p>VerticalPodAutoscalerCondition describes the state of
a VerticalPodAutoscaler at a certain point.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `type` | [VerticalPodAutoscalerConditionType](#autoscaling.k8s.io/v1.VerticalPodAutoscalerConditionType) |
| `status` | [Kubernetes core/v1.ConditionStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#conditionstatus-v1-core) |
| `lastTransitionTime` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `reason` | string
|
| `message` | string
|
### VerticalPodAutoscalerConditionType
<a id="autoscaling.k8s.io/v1.VerticalPodAutoscalerConditionType"></a>
(*Alias*: `string`)
(*Appears on*: 
[VerticalPodAutoscalerCondition](#autoscaling.k8s.io/v1.VerticalPodAutoscalerCondition)
)
> <p>VerticalPodAutoscalerConditionType are the valid conditions of
a VerticalPodAutoscaler.</p>
### VerticalPodAutoscalerRecommenderSelector
<a id="autoscaling.k8s.io/v1.VerticalPodAutoscalerRecommenderSelector"></a>
(*Appears on*: 
[VerticalPodAutoscalerSpec](#autoscaling.k8s.io/v1.VerticalPodAutoscalerSpec)
)
> <p>VerticalPodAutoscalerRecommenderSelector points to a specific Vertical Pod Autoscaler recommender.
In the future it might pass parameters to the recommender.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `name` | string
|
### VerticalPodAutoscalerSpec
<a id="autoscaling.k8s.io/v1.VerticalPodAutoscalerSpec"></a>
(*Appears on*: 
[VerticalPodAutoscaler](#autoscaling.k8s.io/v1.VerticalPodAutoscaler)
)
> <p>VerticalPodAutoscalerSpec is the specification of the behavior of the autoscaler.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `targetRef` | [Kubernetes autoscaling/v1.CrossVersionObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#crossversionobjectreference-v1-autoscaling) |
| `updatePolicy` | [PodUpdatePolicy](#autoscaling.k8s.io/v1.PodUpdatePolicy) |
| `resourcePolicy` | [PodResourcePolicy](#autoscaling.k8s.io/v1.PodResourcePolicy) |
| `recommenders` | [[]VerticalPodAutoscalerRecommenderSelector](#autoscaling.k8s.io/v1.VerticalPodAutoscalerRecommenderSelector) |
### VerticalPodAutoscalerStatus
<a id="autoscaling.k8s.io/v1.VerticalPodAutoscalerStatus"></a>
(*Appears on*: 
[VerticalPodAutoscaler](#autoscaling.k8s.io/v1.VerticalPodAutoscaler)
)
> <p>VerticalPodAutoscalerStatus describes the runtime state of the autoscaler.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `recommendation` | [RecommendedPodResources](#autoscaling.k8s.io/v1.RecommendedPodResources) |
| `conditions` | [[]VerticalPodAutoscalerCondition](#autoscaling.k8s.io/v1.VerticalPodAutoscalerCondition) |
---
## autoscaling.k8s.io/v1beta1 package display name
> <p>Package v1beta1 contains definitions of Vertical Pod Autoscaler related objects.</p>
### Resource Types:
- [VerticalPodAutoscaler](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscaler)
- [VerticalPodAutoscalerCheckpoint](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCheckpoint)
### VerticalPodAutoscaler
<a id="autoscaling.k8s.io/v1beta1.VerticalPodAutoscaler"></a>
> <p>VerticalPodAutoscaler is the configuration for a vertical pod
autoscaler, which automatically manages pod resources based on historical and
real time resource utilization.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `apiVersion` | `autoscaling.k8s.io/v1beta1` |
| `kind`        | `VerticalPodAutoscaler` |
| `metadata` | [Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#objectmeta-v1-meta) |
| `spec` | [VerticalPodAutoscalerSpec](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerSpec) |
| `status` | [VerticalPodAutoscalerStatus](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerStatus) |
### VerticalPodAutoscalerCheckpoint
<a id="autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCheckpoint"></a>
> <p>VerticalPodAutoscalerCheckpoint is the checkpoint of the internal state of VPA that
is used for recovery after recommender&rsquo;s restart.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `apiVersion` | `autoscaling.k8s.io/v1beta1` |
| `kind`        | `VerticalPodAutoscalerCheckpoint` |
| `metadata` | [Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#objectmeta-v1-meta) |
| `spec` | [VerticalPodAutoscalerCheckpointSpec](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCheckpointSpec) |
| `status` | [VerticalPodAutoscalerCheckpointStatus](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCheckpointStatus) |
### ContainerResourcePolicy
<a id="autoscaling.k8s.io/v1beta1.ContainerResourcePolicy"></a>
(*Appears on*: 
[PodResourcePolicy](#autoscaling.k8s.io/v1beta1.PodResourcePolicy)
)
> <p>ContainerResourcePolicy controls how autoscaler computes the recommended
resources for a specific container.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerName` | string
|
| `mode` | [ContainerScalingMode](#autoscaling.k8s.io/v1beta1.ContainerScalingMode) |
| `minAllowed` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `maxAllowed` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
### ContainerScalingMode
<a id="autoscaling.k8s.io/v1beta1.ContainerScalingMode"></a>
(*Alias*: `string`)
(*Appears on*: 
[ContainerResourcePolicy](#autoscaling.k8s.io/v1beta1.ContainerResourcePolicy)
)
> <p>ContainerScalingMode controls whether autoscaler is enabled for a specific
container.</p>
#### Constants
| Value-Type             | Description         |
|--------------------|---------------------|
| `&#34;Auto&#34;` | <p>ContainerScalingModeAuto means autoscaling is enabled for a container.</p>
|
| `&#34;Off&#34;` | <p>ContainerScalingModeOff means autoscaling is disabled for a container.</p>
|
### HistogramCheckpoint
<a id="autoscaling.k8s.io/v1beta1.HistogramCheckpoint"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpointStatus](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCheckpointStatus)
)
> <p>HistogramCheckpoint contains data needed to reconstruct the histogram.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `referenceTimestamp` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `bucketWeights` | map[int]uint32
|
| `totalWeight` | float64
|
### PodResourcePolicy
<a id="autoscaling.k8s.io/v1beta1.PodResourcePolicy"></a>
(*Appears on*: 
[VerticalPodAutoscalerSpec](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerSpec)
)
> <p>PodResourcePolicy controls how autoscaler computes the recommended resources
for containers belonging to the pod. There can be at most one entry for every
named container and optionally a single wildcard entry with <code>containerName</code> = &lsquo;*&rsquo;,
which handles all containers that don&rsquo;t have individual policies.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerPolicies` | [[]ContainerResourcePolicy](#autoscaling.k8s.io/v1beta1.ContainerResourcePolicy) |
### PodUpdatePolicy
<a id="autoscaling.k8s.io/v1beta1.PodUpdatePolicy"></a>
(*Appears on*: 
[VerticalPodAutoscalerSpec](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerSpec)
)
> <p>PodUpdatePolicy describes the rules on how changes are applied to the pods.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `updateMode` | [UpdateMode](#autoscaling.k8s.io/v1beta1.UpdateMode) |
### RecommendedContainerResources
<a id="autoscaling.k8s.io/v1beta1.RecommendedContainerResources"></a>
(*Appears on*: 
[RecommendedPodResources](#autoscaling.k8s.io/v1beta1.RecommendedPodResources)
)
> <p>RecommendedContainerResources is the recommendation of resources computed by
autoscaler for a specific container. Respects the container resource policy
if present in the spec. In particular the recommendation is not produced for
containers with <code>ContainerScalingMode</code> set to &lsquo;Off&rsquo;.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerName` | string
|
| `target` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `lowerBound` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `upperBound` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `uncappedTarget` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
### RecommendedPodResources
<a id="autoscaling.k8s.io/v1beta1.RecommendedPodResources"></a>
(*Appears on*: 
[VerticalPodAutoscalerStatus](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerStatus)
)
> <p>RecommendedPodResources is the recommendation of resources computed by
autoscaler. It contains a recommendation for each container in the pod
(except for those with <code>ContainerScalingMode</code> set to &lsquo;Off&rsquo;).</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerRecommendations` | [[]RecommendedContainerResources](#autoscaling.k8s.io/v1beta1.RecommendedContainerResources) |
### UpdateMode
<a id="autoscaling.k8s.io/v1beta1.UpdateMode"></a>
(*Alias*: `string`)
(*Appears on*: 
[PodUpdatePolicy](#autoscaling.k8s.io/v1beta1.PodUpdatePolicy)
)
> <p>UpdateMode controls when autoscaler applies changes to the pod resources.</p>
#### Constants
| Value-Type             | Description         |
|--------------------|---------------------|
| `&#34;Auto&#34;` | <p>UpdateModeAuto means that autoscaler assigns resources on pod creation
and additionally can update them during the lifetime of the pod,
using any available update method. Currently this is equivalent to
Recreate, which is the only available update method.</p>
|
| `&#34;Initial&#34;` | <p>UpdateModeInitial means that autoscaler only assigns resources on pod
creation and does not change them during the lifetime of the pod.</p>
|
| `&#34;Off&#34;` | <p>UpdateModeOff means that autoscaler never changes Pod resources.
The recommender still sets the recommended resources in the
VerticalPodAutoscaler object. This can be used for a &ldquo;dry run&rdquo;.</p>
|
| `&#34;Recreate&#34;` | <p>UpdateModeRecreate means that autoscaler assigns resources on pod
creation and additionally can update them during the lifetime of the
pod by deleting and recreating the pod.</p>
|
### VerticalPodAutoscalerCheckpointSpec
<a id="autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCheckpointSpec"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpoint](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCheckpoint)
)
> <p>VerticalPodAutoscalerCheckpointSpec is the specification of the checkpoint object.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `vpaObjectName` | string
|
| `containerName` | string
|
### VerticalPodAutoscalerCheckpointStatus
<a id="autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCheckpointStatus"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpoint](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCheckpoint)
)
> <p>VerticalPodAutoscalerCheckpointStatus contains data of the checkpoint.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `lastUpdateTime` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `version` | string
|
| `cpuHistogram` | [HistogramCheckpoint](#autoscaling.k8s.io/v1beta1.HistogramCheckpoint) |
| `memoryHistogram` | [HistogramCheckpoint](#autoscaling.k8s.io/v1beta1.HistogramCheckpoint) |
| `firstSampleStart` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `lastSampleStart` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `totalSamplesCount` | int
|
### VerticalPodAutoscalerCondition
<a id="autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCondition"></a>
(*Appears on*: 
[VerticalPodAutoscalerStatus](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerStatus)
)
> <p>VerticalPodAutoscalerCondition describes the state of
a VerticalPodAutoscaler at a certain point.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `type` | [VerticalPodAutoscalerConditionType](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerConditionType) |
| `status` | [Kubernetes core/v1.ConditionStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#conditionstatus-v1-core) |
| `lastTransitionTime` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `reason` | string
|
| `message` | string
|
### VerticalPodAutoscalerConditionType
<a id="autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerConditionType"></a>
(*Alias*: `string`)
(*Appears on*: 
[VerticalPodAutoscalerCondition](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCondition)
)
> <p>VerticalPodAutoscalerConditionType are the valid conditions of
a VerticalPodAutoscaler.</p>
### VerticalPodAutoscalerSpec
<a id="autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerSpec"></a>
(*Appears on*: 
[VerticalPodAutoscaler](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscaler)
)
> <p>VerticalPodAutoscalerSpec is the specification of the behavior of the autoscaler.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `selector` | [Kubernetes meta/v1.LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#labelselector-v1-meta) |
| `updatePolicy` | [PodUpdatePolicy](#autoscaling.k8s.io/v1beta1.PodUpdatePolicy) |
| `resourcePolicy` | [PodResourcePolicy](#autoscaling.k8s.io/v1beta1.PodResourcePolicy) |
### VerticalPodAutoscalerStatus
<a id="autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerStatus"></a>
(*Appears on*: 
[VerticalPodAutoscaler](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscaler)
)
> <p>VerticalPodAutoscalerStatus describes the runtime state of the autoscaler.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `recommendation` | [RecommendedPodResources](#autoscaling.k8s.io/v1beta1.RecommendedPodResources) |
| `conditions` | [[]VerticalPodAutoscalerCondition](#autoscaling.k8s.io/v1beta1.VerticalPodAutoscalerCondition) |
---
## autoscaling.k8s.io/v1beta2 package display name
> <p>Package v1beta2 contains definitions of Vertical Pod Autoscaler related objects.</p>
### Resource Types:
- [VerticalPodAutoscaler](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscaler)
- [VerticalPodAutoscalerCheckpoint](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCheckpoint)
### VerticalPodAutoscaler
<a id="autoscaling.k8s.io/v1beta2.VerticalPodAutoscaler"></a>
> <p>VerticalPodAutoscaler is the configuration for a vertical pod
autoscaler, which automatically manages pod resources based on historical and
real time resource utilization.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `apiVersion` | `autoscaling.k8s.io/v1beta2` |
| `kind`        | `VerticalPodAutoscaler` |
| `metadata` | [Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#objectmeta-v1-meta) |
| `spec` | [VerticalPodAutoscalerSpec](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerSpec) |
| `status` | [VerticalPodAutoscalerStatus](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerStatus) |
### VerticalPodAutoscalerCheckpoint
<a id="autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCheckpoint"></a>
> <p>VerticalPodAutoscalerCheckpoint is the checkpoint of the internal state of VPA that
is used for recovery after recommender&rsquo;s restart.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `apiVersion` | `autoscaling.k8s.io/v1beta2` |
| `kind`        | `VerticalPodAutoscalerCheckpoint` |
| `metadata` | [Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#objectmeta-v1-meta) |
| `spec` | [VerticalPodAutoscalerCheckpointSpec](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCheckpointSpec) |
| `status` | [VerticalPodAutoscalerCheckpointStatus](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCheckpointStatus) |
### ContainerResourcePolicy
<a id="autoscaling.k8s.io/v1beta2.ContainerResourcePolicy"></a>
(*Appears on*: 
[PodResourcePolicy](#autoscaling.k8s.io/v1beta2.PodResourcePolicy)
)
> <p>ContainerResourcePolicy controls how autoscaler computes the recommended
resources for a specific container.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerName` | string
|
| `mode` | [ContainerScalingMode](#autoscaling.k8s.io/v1beta2.ContainerScalingMode) |
| `minAllowed` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `maxAllowed` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
### ContainerScalingMode
<a id="autoscaling.k8s.io/v1beta2.ContainerScalingMode"></a>
(*Alias*: `string`)
(*Appears on*: 
[ContainerResourcePolicy](#autoscaling.k8s.io/v1beta2.ContainerResourcePolicy)
)
> <p>ContainerScalingMode controls whether autoscaler is enabled for a specific
container.</p>
#### Constants
| Value-Type             | Description         |
|--------------------|---------------------|
| `&#34;Auto&#34;` | <p>ContainerScalingModeAuto means autoscaling is enabled for a container.</p>
|
| `&#34;Off&#34;` | <p>ContainerScalingModeOff means autoscaling is disabled for a container.</p>
|
### HistogramCheckpoint
<a id="autoscaling.k8s.io/v1beta2.HistogramCheckpoint"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpointStatus](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCheckpointStatus)
)
> <p>HistogramCheckpoint contains data needed to reconstruct the histogram.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `referenceTimestamp` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `bucketWeights` | map[int]uint32
|
| `totalWeight` | float64
|
### PodResourcePolicy
<a id="autoscaling.k8s.io/v1beta2.PodResourcePolicy"></a>
(*Appears on*: 
[VerticalPodAutoscalerSpec](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerSpec)
)
> <p>PodResourcePolicy controls how autoscaler computes the recommended resources
for containers belonging to the pod. There can be at most one entry for every
named container and optionally a single wildcard entry with <code>containerName</code> = &lsquo;*&rsquo;,
which handles all containers that don&rsquo;t have individual policies.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerPolicies` | [[]ContainerResourcePolicy](#autoscaling.k8s.io/v1beta2.ContainerResourcePolicy) |
### PodUpdatePolicy
<a id="autoscaling.k8s.io/v1beta2.PodUpdatePolicy"></a>
(*Appears on*: 
[VerticalPodAutoscalerSpec](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerSpec)
)
> <p>PodUpdatePolicy describes the rules on how changes are applied to the pods.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `updateMode` | [UpdateMode](#autoscaling.k8s.io/v1beta2.UpdateMode) |
### RecommendedContainerResources
<a id="autoscaling.k8s.io/v1beta2.RecommendedContainerResources"></a>
(*Appears on*: 
[RecommendedPodResources](#autoscaling.k8s.io/v1beta2.RecommendedPodResources)
)
> <p>RecommendedContainerResources is the recommendation of resources computed by
autoscaler for a specific container. Respects the container resource policy
if present in the spec. In particular the recommendation is not produced for
containers with <code>ContainerScalingMode</code> set to &lsquo;Off&rsquo;.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerName` | string
|
| `target` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `lowerBound` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `upperBound` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `uncappedTarget` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
### RecommendedPodResources
<a id="autoscaling.k8s.io/v1beta2.RecommendedPodResources"></a>
(*Appears on*: 
[VerticalPodAutoscalerStatus](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerStatus)
)
> <p>RecommendedPodResources is the recommendation of resources computed by
autoscaler. It contains a recommendation for each container in the pod
(except for those with <code>ContainerScalingMode</code> set to &lsquo;Off&rsquo;).</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerRecommendations` | [[]RecommendedContainerResources](#autoscaling.k8s.io/v1beta2.RecommendedContainerResources) |
### UpdateMode
<a id="autoscaling.k8s.io/v1beta2.UpdateMode"></a>
(*Alias*: `string`)
(*Appears on*: 
[PodUpdatePolicy](#autoscaling.k8s.io/v1beta2.PodUpdatePolicy)
)
> <p>UpdateMode controls when autoscaler applies changes to the pod resources.</p>
#### Constants
| Value-Type             | Description         |
|--------------------|---------------------|
| `&#34;Auto&#34;` | <p>UpdateModeAuto means that autoscaler assigns resources on pod creation
and additionally can update them during the lifetime of the pod,
using any available update method. Currently this is equivalent to
Recreate, which is the only available update method.</p>
|
| `&#34;Initial&#34;` | <p>UpdateModeInitial means that autoscaler only assigns resources on pod
creation and does not change them during the lifetime of the pod.</p>
|
| `&#34;Off&#34;` | <p>UpdateModeOff means that autoscaler never changes Pod resources.
The recommender still sets the recommended resources in the
VerticalPodAutoscaler object. This can be used for a &ldquo;dry run&rdquo;.</p>
|
| `&#34;Recreate&#34;` | <p>UpdateModeRecreate means that autoscaler assigns resources on pod
creation and additionally can update them during the lifetime of the
pod by deleting and recreating the pod.</p>
|
### VerticalPodAutoscalerCheckpointSpec
<a id="autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCheckpointSpec"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpoint](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCheckpoint)
)
> <p>VerticalPodAutoscalerCheckpointSpec is the specification of the checkpoint object.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `vpaObjectName` | string
|
| `containerName` | string
|
### VerticalPodAutoscalerCheckpointStatus
<a id="autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCheckpointStatus"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpoint](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCheckpoint)
)
> <p>VerticalPodAutoscalerCheckpointStatus contains data of the checkpoint.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `lastUpdateTime` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `version` | string
|
| `cpuHistogram` | [HistogramCheckpoint](#autoscaling.k8s.io/v1beta2.HistogramCheckpoint) |
| `memoryHistogram` | [HistogramCheckpoint](#autoscaling.k8s.io/v1beta2.HistogramCheckpoint) |
| `firstSampleStart` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `lastSampleStart` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `totalSamplesCount` | int
|
### VerticalPodAutoscalerCondition
<a id="autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCondition"></a>
(*Appears on*: 
[VerticalPodAutoscalerStatus](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerStatus)
)
> <p>VerticalPodAutoscalerCondition describes the state of
a VerticalPodAutoscaler at a certain point.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `type` | [VerticalPodAutoscalerConditionType](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerConditionType) |
| `status` | [Kubernetes core/v1.ConditionStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#conditionstatus-v1-core) |
| `lastTransitionTime` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `reason` | string
|
| `message` | string
|
### VerticalPodAutoscalerConditionType
<a id="autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerConditionType"></a>
(*Alias*: `string`)
(*Appears on*: 
[VerticalPodAutoscalerCondition](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCondition)
)
> <p>VerticalPodAutoscalerConditionType are the valid conditions of
a VerticalPodAutoscaler.</p>
### VerticalPodAutoscalerSpec
<a id="autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerSpec"></a>
(*Appears on*: 
[VerticalPodAutoscaler](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscaler)
)
> <p>VerticalPodAutoscalerSpec is the specification of the behavior of the autoscaler.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `targetRef` | [Kubernetes autoscaling/v1.CrossVersionObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#crossversionobjectreference-v1-autoscaling) |
| `updatePolicy` | [PodUpdatePolicy](#autoscaling.k8s.io/v1beta2.PodUpdatePolicy) |
| `resourcePolicy` | [PodResourcePolicy](#autoscaling.k8s.io/v1beta2.PodResourcePolicy) |
### VerticalPodAutoscalerStatus
<a id="autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerStatus"></a>
(*Appears on*: 
[VerticalPodAutoscaler](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscaler)
)
> <p>VerticalPodAutoscalerStatus describes the runtime state of the autoscaler.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `recommendation` | [RecommendedPodResources](#autoscaling.k8s.io/v1beta2.RecommendedPodResources) |
| `conditions` | [[]VerticalPodAutoscalerCondition](#autoscaling.k8s.io/v1beta2.VerticalPodAutoscalerCondition) |
---
## poc.autoscaling.k8s.io/v1alpha1 package display name
> <p>Package v1alpha1 contains definitions of Vertical Pod Autoscaler related objects.</p>
### Resource Types:
- [VerticalPodAutoscaler](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscaler)
- [VerticalPodAutoscalerCheckpoint](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCheckpoint)
### VerticalPodAutoscaler
<a id="poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscaler"></a>
> <p>VerticalPodAutoscaler is the configuration for a vertical pod
autoscaler, which automatically manages pod resources based on historical and
real time resource utilization.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `apiVersion` | `poc.autoscaling.k8s.io/v1alpha1` |
| `kind`        | `VerticalPodAutoscaler` |
| `metadata` | [Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#objectmeta-v1-meta) |
| `spec` | [VerticalPodAutoscalerSpec](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerSpec) |
| `status` | [VerticalPodAutoscalerStatus](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerStatus) |
### VerticalPodAutoscalerCheckpoint
<a id="poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCheckpoint"></a>
> <p>VerticalPodAutoscalerCheckpoint is the checkpoint of the internal state of VPA that
is used for recovery after recommender&rsquo;s restart.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `apiVersion` | `poc.autoscaling.k8s.io/v1alpha1` |
| `kind`        | `VerticalPodAutoscalerCheckpoint` |
| `metadata` | [Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#objectmeta-v1-meta) |
| `spec` | [VerticalPodAutoscalerCheckpointSpec](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCheckpointSpec) |
| `status` | [VerticalPodAutoscalerCheckpointStatus](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCheckpointStatus) |
### ContainerResourcePolicy
<a id="poc.autoscaling.k8s.io/v1alpha1.ContainerResourcePolicy"></a>
(*Appears on*: 
[PodResourcePolicy](#poc.autoscaling.k8s.io/v1alpha1.PodResourcePolicy)
)
> <p>ContainerResourcePolicy controls how autoscaler computes the recommended
resources for a specific container.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerName` | string
|
| `mode` | [ContainerScalingMode](#poc.autoscaling.k8s.io/v1alpha1.ContainerScalingMode) |
| `minAllowed` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `maxAllowed` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
### ContainerScalingMode
<a id="poc.autoscaling.k8s.io/v1alpha1.ContainerScalingMode"></a>
(*Alias*: `string`)
(*Appears on*: 
[ContainerResourcePolicy](#poc.autoscaling.k8s.io/v1alpha1.ContainerResourcePolicy)
)
> <p>ContainerScalingMode controls whether autoscaler is enabled for a specific
container.</p>
#### Constants
| Value-Type             | Description         |
|--------------------|---------------------|
| `&#34;Auto&#34;` | <p>ContainerScalingModeAuto means autoscaling is enabled for a container.</p>
|
| `&#34;Off&#34;` | <p>ContainerScalingModeOff means autoscaling is disabled for a container.</p>
|
### HistogramCheckpoint
<a id="poc.autoscaling.k8s.io/v1alpha1.HistogramCheckpoint"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpointStatus](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCheckpointStatus)
)
> <p>HistogramCheckpoint contains data needed to reconstruct the histogram.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `referenceTimestamp` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `bucketWeights` | map[int]uint32
|
| `totalWeight` | float64
|
### PodResourcePolicy
<a id="poc.autoscaling.k8s.io/v1alpha1.PodResourcePolicy"></a>
(*Appears on*: 
[VerticalPodAutoscalerSpec](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerSpec)
)
> <p>PodResourcePolicy controls how autoscaler computes the recommended resources
for containers belonging to the pod. There can be at most one entry for every
named container and optionally a single wildcard entry with <code>containerName</code> = &lsquo;*&rsquo;,
which handles all containers that don&rsquo;t have individual policies.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerPolicies` | [[]ContainerResourcePolicy](#poc.autoscaling.k8s.io/v1alpha1.ContainerResourcePolicy) |
### PodUpdatePolicy
<a id="poc.autoscaling.k8s.io/v1alpha1.PodUpdatePolicy"></a>
(*Appears on*: 
[VerticalPodAutoscalerSpec](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerSpec)
)
> <p>PodUpdatePolicy describes the rules on how changes are applied to the pods.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `updateMode` | [UpdateMode](#poc.autoscaling.k8s.io/v1alpha1.UpdateMode) |
### RecommendedContainerResources
<a id="poc.autoscaling.k8s.io/v1alpha1.RecommendedContainerResources"></a>
(*Appears on*: 
[RecommendedPodResources](#poc.autoscaling.k8s.io/v1alpha1.RecommendedPodResources)
)
> <p>RecommendedContainerResources is the recommendation of resources computed by
autoscaler for a specific container. Respects the container resource policy
if present in the spec. In particular the recommendation is not produced for
containers with <code>ContainerScalingMode</code> set to &lsquo;Off&rsquo;.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerName` | string
|
| `target` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `lowerBound` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
| `upperBound` | [Kubernetes core/v1.ResourceList](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#resourcelist-v1-core) |
### RecommendedPodResources
<a id="poc.autoscaling.k8s.io/v1alpha1.RecommendedPodResources"></a>
(*Appears on*: 
[VerticalPodAutoscalerStatus](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerStatus)
)
> <p>RecommendedPodResources is the recommendation of resources computed by
autoscaler. It contains a recommendation for each container in the pod
(except for those with <code>ContainerScalingMode</code> set to &lsquo;Off&rsquo;).</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `containerRecommendations` | [[]RecommendedContainerResources](#poc.autoscaling.k8s.io/v1alpha1.RecommendedContainerResources) |
### UpdateMode
<a id="poc.autoscaling.k8s.io/v1alpha1.UpdateMode"></a>
(*Alias*: `string`)
(*Appears on*: 
[PodUpdatePolicy](#poc.autoscaling.k8s.io/v1alpha1.PodUpdatePolicy)
)
> <p>UpdateMode controls when autoscaler applies changes to the pod resources.</p>
#### Constants
| Value-Type             | Description         |
|--------------------|---------------------|
| `&#34;Auto&#34;` | <p>UpdateModeAuto means that autoscaler assigns resources on pod creation
and additionally can update them during the lifetime of the pod,
using any available update method. Currently this is equivalent to
Recreate, which is the only available update method.</p>
|
| `&#34;Initial&#34;` | <p>UpdateModeInitial means that autoscaler only assigns resources on pod
creation and does not change them during the lifetime of the pod.</p>
|
| `&#34;Off&#34;` | <p>UpdateModeOff means that autoscaler never changes Pod resources.
The recommender still sets the recommended resources in the
VerticalPodAutoscaler object. This can be used for a &ldquo;dry run&rdquo;.</p>
|
| `&#34;Recreate&#34;` | <p>UpdateModeRecreate means that autoscaler assigns resources on pod
creation and additionally can update them during the lifetime of the
pod by deleting and recreating the pod.</p>
|
### VerticalPodAutoscalerCheckpointSpec
<a id="poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCheckpointSpec"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpoint](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCheckpoint)
)
> <p>VerticalPodAutoscalerCheckpointSpec is the specification of the checkpoint object.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `vpaObjectName` | string
|
| `containerName` | string
|
### VerticalPodAutoscalerCheckpointStatus
<a id="poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCheckpointStatus"></a>
(*Appears on*: 
[VerticalPodAutoscalerCheckpoint](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCheckpoint)
)
> <p>VerticalPodAutoscalerCheckpointStatus contains data of the checkpoint.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `lastUpdateTime` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `version` | string
|
| `cpuHistogram` | [HistogramCheckpoint](#poc.autoscaling.k8s.io/v1alpha1.HistogramCheckpoint) |
| `memoryHistogram` | [HistogramCheckpoint](#poc.autoscaling.k8s.io/v1alpha1.HistogramCheckpoint) |
| `firstSampleStart` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `lastSampleStart` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `totalSamplesCount` | int
|
### VerticalPodAutoscalerCondition
<a id="poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCondition"></a>
(*Appears on*: 
[VerticalPodAutoscalerStatus](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerStatus)
)
> <p>VerticalPodAutoscalerCondition describes the state of
a VerticalPodAutoscaler at a certain point.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `type` | [VerticalPodAutoscalerConditionType](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerConditionType) |
| `status` | [Kubernetes core/v1.ConditionStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#conditionstatus-v1-core) |
| `lastTransitionTime` | [Kubernetes meta/v1.Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#time-v1-meta) |
| `reason` | string
|
| `message` | string
|
### VerticalPodAutoscalerConditionType
<a id="poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerConditionType"></a>
(*Alias*: `string`)
(*Appears on*: 
[VerticalPodAutoscalerCondition](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCondition)
)
> <p>VerticalPodAutoscalerConditionType are the valid conditions of
a VerticalPodAutoscaler.</p>
### VerticalPodAutoscalerSpec
<a id="poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerSpec"></a>
(*Appears on*: 
[VerticalPodAutoscaler](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscaler)
)
> <p>VerticalPodAutoscalerSpec is the specification of the behavior of the autoscaler.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `selector` | [Kubernetes meta/v1.LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.32/#labelselector-v1-meta) |
| `updatePolicy` | [PodUpdatePolicy](#poc.autoscaling.k8s.io/v1alpha1.PodUpdatePolicy) |
| `resourcePolicy` | [PodResourcePolicy](#poc.autoscaling.k8s.io/v1alpha1.PodResourcePolicy) |
### VerticalPodAutoscalerStatus
<a id="poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerStatus"></a>
(*Appears on*: 
[VerticalPodAutoscaler](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscaler)
)
> <p>VerticalPodAutoscalerStatus describes the runtime state of the autoscaler.</p>
#### Members
| Field-Type            | Description         |
|-------------------|---------------------|
| `recommendation` | [RecommendedPodResources](#poc.autoscaling.k8s.io/v1alpha1.RecommendedPodResources) |
| `conditions` | [[]VerticalPodAutoscalerCondition](#poc.autoscaling.k8s.io/v1alpha1.VerticalPodAutoscalerCondition) |
---
*Generated with `gen-crd-api-reference-docs` on git commit `cab7b83f6`.*
