package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExternalMetric describes a ExternalMetric resource
type ExternalMetric struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	meta_v1.TypeMeta `json:",inline"`

	// ObjectMeta contains the metadata for the particular object (name, namespace, self link, labels, etc)
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the custom resource spec
	Spec ExternalMetricSpec `json:"spec"`
}

// ExternalMetricSpec is the spec for a ExternalMetric resource
type ExternalMetricSpec struct {
	// Message and SomeValue are example custom spec fields
	//
	// this is where you would put your custom resource data
	Message   string `json:"message"`
	SomeValue *int32 `json:"someValue"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExternalMetricsList is a list of ExternalMetric resources
type ExternalMetricList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []ExternalMetric `json:"items"`
}
