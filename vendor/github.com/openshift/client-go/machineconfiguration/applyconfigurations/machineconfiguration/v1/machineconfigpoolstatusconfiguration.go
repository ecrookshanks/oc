// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	types "k8s.io/apimachinery/pkg/types"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// MachineConfigPoolStatusConfigurationApplyConfiguration represents a declarative configuration of the MachineConfigPoolStatusConfiguration type for use
// with apply.
type MachineConfigPoolStatusConfigurationApplyConfiguration struct {
	corev1.ObjectReferenceApplyConfiguration `json:",inline"`
	Source                                   []corev1.ObjectReferenceApplyConfiguration `json:"source,omitempty"`
}

// MachineConfigPoolStatusConfigurationApplyConfiguration constructs a declarative configuration of the MachineConfigPoolStatusConfiguration type for use with
// apply.
func MachineConfigPoolStatusConfiguration() *MachineConfigPoolStatusConfigurationApplyConfiguration {
	return &MachineConfigPoolStatusConfigurationApplyConfiguration{}
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *MachineConfigPoolStatusConfigurationApplyConfiguration) WithKind(value string) *MachineConfigPoolStatusConfigurationApplyConfiguration {
	b.ObjectReferenceApplyConfiguration.Kind = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *MachineConfigPoolStatusConfigurationApplyConfiguration) WithNamespace(value string) *MachineConfigPoolStatusConfigurationApplyConfiguration {
	b.ObjectReferenceApplyConfiguration.Namespace = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *MachineConfigPoolStatusConfigurationApplyConfiguration) WithName(value string) *MachineConfigPoolStatusConfigurationApplyConfiguration {
	b.ObjectReferenceApplyConfiguration.Name = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *MachineConfigPoolStatusConfigurationApplyConfiguration) WithUID(value types.UID) *MachineConfigPoolStatusConfigurationApplyConfiguration {
	b.ObjectReferenceApplyConfiguration.UID = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *MachineConfigPoolStatusConfigurationApplyConfiguration) WithAPIVersion(value string) *MachineConfigPoolStatusConfigurationApplyConfiguration {
	b.ObjectReferenceApplyConfiguration.APIVersion = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *MachineConfigPoolStatusConfigurationApplyConfiguration) WithResourceVersion(value string) *MachineConfigPoolStatusConfigurationApplyConfiguration {
	b.ObjectReferenceApplyConfiguration.ResourceVersion = &value
	return b
}

// WithFieldPath sets the FieldPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FieldPath field is set to the value of the last call.
func (b *MachineConfigPoolStatusConfigurationApplyConfiguration) WithFieldPath(value string) *MachineConfigPoolStatusConfigurationApplyConfiguration {
	b.ObjectReferenceApplyConfiguration.FieldPath = &value
	return b
}

// WithSource adds the given value to the Source field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Source field.
func (b *MachineConfigPoolStatusConfigurationApplyConfiguration) WithSource(values ...*corev1.ObjectReferenceApplyConfiguration) *MachineConfigPoolStatusConfigurationApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSource")
		}
		b.Source = append(b.Source, *values[i])
	}
	return b
}
