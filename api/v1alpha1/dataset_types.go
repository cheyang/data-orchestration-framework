/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/**
* apiVersion: "data-orchestration.org/v1aphla1"
* kind: "Dataset"
* metadata:
*   name: "imagenet"
* spec:
*   # target: /imagenet # optional
*   mountPoint: oss://imagenet-huabei5/images/
*   options:
*      fs.oss.endpoint: oss-cn-huhehaote-internal.aliyuncs.com
*      fs.oss.accessKeyId: xxx
*      fs.oss.accessKeySecret: yyy
*   minReplicas: 1 # optional
*   maxReplicas: 3
*   affinity:
*     nodeAffinity:
*       preferredDuringSchedulingIgnoredDuringExecution:
*       - weight: 1
*         preference:
*           matchExpressions:
*           - key: another-node-label-key
*             operator: Exists
* status:
*   total: 50GiB
*   ufsTotal: 50GiB
*   cacheStatus:
*       cached: 15GiB
*       cachable: 40GiB
*       needMoreForCache: 10GiB
 */

// DatasetSpec defines the desired state of Dataset
type DatasetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// MountPoint is the mount point of source
	MountPoint string `json:"mountPoint,omitempty"`

	// Options to configure
	Options map[string]string `json:"options,omitempty"`

	// MinReplicas is the min replicas of dataset in the cluster
	// Optional.
	MinReplicas *int32 `json:"minReplicas,omitempty"`

	// maxReplicas is the max replicas of dataset in the cluster
	// Optional.
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`

	// NodeAffinity defines constraints that limit what nodes this dataset can be cached to.
	// This field influences the scheduling of pods that use the cached dataset.
	// +optional
	NodeAffinity *CacheableNodeAffinity
}

// CacheableNodeAffinity defines constraints that limit what nodes this dataset can be cached to.
type CacheableNodeAffinity struct {
	// Required specifies hard node constraints that must be met.
	Required *v1.NodeSelector
}

// CacheStateName is the name identifying various cacheStateName in a CacheStateNameList.
type CacheStateName string

// ResourceList is a set of (resource name, quantity) pairs.
type CacheStateList map[CacheStateName]resource.Quantity

// CacheStateName names must be not more than 63 characters, consisting of upper- or lower-case alphanumeric characters,
// with the -, _, and . characters allowed anywhere, except the first or last character.
// The default convention, matching that for annotations, is to use lower-case names, with dashes, rather than
// camel case, separating compound words.
// Fully-qualified resource typenames are constructed from a DNS-style subdomain, followed by a slash `/` and a name.
const (
	// Cached in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)
	Cached CacheStateName = "cached"
	// Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)
	Cacheable CacheStateName = "cacheable"
	// NeedMoreForCache size, in bytes (e,g. 5Gi = 5GiB = 5 * 1024 * 1024 * 1024)
	NeedMoreForCache CacheStateName = "needMoreForCache"
	// Percentage represents the cache percentage over the total data in the underlayer filesystem.
	// 1.5 = 1500m
	CachedPercentage CacheStateName = "cachedPercentage"
)

// The cache phase indicates whether the loading is behaving
type CachePhase string

const (
	// planning the cache
	// Planning CachePhase = "planning"
	// Loading the cache
	Loading CachePhase = "loading"
	// loaded the cache
	Ready CachePhase = "ready"
)

// DatasetStatus defines the observed state of Dataset
type DatasetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Total in bytes of the dataset to load in the cluster
	Total resource.Quantity `json:"total,omitempty"`

	// Total in bytes of dataset in the cluster
	UfsTotal resource.Quantity `json:"ufsTotal,omitempty"`

	// CacheStatus represents the total resources of the dataset.
	CacheStatus CacheStateList `json:"cacheStatus,omitempty"`

	Phase CachePhase `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true

// Dataset is the Schema for the datasets API
type Dataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatasetSpec   `json:"spec,omitempty"`
	Status DatasetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetList contains a list of Dataset
type DatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dataset{}, &DatasetList{})
}
