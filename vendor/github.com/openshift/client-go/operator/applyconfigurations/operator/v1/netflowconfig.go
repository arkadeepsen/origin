// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
)

// NetFlowConfigApplyConfiguration represents a declarative configuration of the NetFlowConfig type for use
// with apply.
type NetFlowConfigApplyConfiguration struct {
	Collectors []operatorv1.IPPort `json:"collectors,omitempty"`
}

// NetFlowConfigApplyConfiguration constructs a declarative configuration of the NetFlowConfig type for use with
// apply.
func NetFlowConfig() *NetFlowConfigApplyConfiguration {
	return &NetFlowConfigApplyConfiguration{}
}

// WithCollectors adds the given value to the Collectors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Collectors field.
func (b *NetFlowConfigApplyConfiguration) WithCollectors(values ...operatorv1.IPPort) *NetFlowConfigApplyConfiguration {
	for i := range values {
		b.Collectors = append(b.Collectors, values[i])
	}
	return b
}
