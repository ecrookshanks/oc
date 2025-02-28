// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// VSpherePlatformFailureDomainSpecApplyConfiguration represents a declarative configuration of the VSpherePlatformFailureDomainSpec type for use
// with apply.
type VSpherePlatformFailureDomainSpecApplyConfiguration struct {
	Name           *string                                               `json:"name,omitempty"`
	Region         *string                                               `json:"region,omitempty"`
	Zone           *string                                               `json:"zone,omitempty"`
	RegionAffinity *VSphereFailureDomainRegionAffinityApplyConfiguration `json:"regionAffinity,omitempty"`
	ZoneAffinity   *VSphereFailureDomainZoneAffinityApplyConfiguration   `json:"zoneAffinity,omitempty"`
	Server         *string                                               `json:"server,omitempty"`
	Topology       *VSpherePlatformTopologyApplyConfiguration            `json:"topology,omitempty"`
}

// VSpherePlatformFailureDomainSpecApplyConfiguration constructs a declarative configuration of the VSpherePlatformFailureDomainSpec type for use with
// apply.
func VSpherePlatformFailureDomainSpec() *VSpherePlatformFailureDomainSpecApplyConfiguration {
	return &VSpherePlatformFailureDomainSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *VSpherePlatformFailureDomainSpecApplyConfiguration) WithName(value string) *VSpherePlatformFailureDomainSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithRegion sets the Region field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Region field is set to the value of the last call.
func (b *VSpherePlatformFailureDomainSpecApplyConfiguration) WithRegion(value string) *VSpherePlatformFailureDomainSpecApplyConfiguration {
	b.Region = &value
	return b
}

// WithZone sets the Zone field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Zone field is set to the value of the last call.
func (b *VSpherePlatformFailureDomainSpecApplyConfiguration) WithZone(value string) *VSpherePlatformFailureDomainSpecApplyConfiguration {
	b.Zone = &value
	return b
}

// WithRegionAffinity sets the RegionAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RegionAffinity field is set to the value of the last call.
func (b *VSpherePlatformFailureDomainSpecApplyConfiguration) WithRegionAffinity(value *VSphereFailureDomainRegionAffinityApplyConfiguration) *VSpherePlatformFailureDomainSpecApplyConfiguration {
	b.RegionAffinity = value
	return b
}

// WithZoneAffinity sets the ZoneAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ZoneAffinity field is set to the value of the last call.
func (b *VSpherePlatformFailureDomainSpecApplyConfiguration) WithZoneAffinity(value *VSphereFailureDomainZoneAffinityApplyConfiguration) *VSpherePlatformFailureDomainSpecApplyConfiguration {
	b.ZoneAffinity = value
	return b
}

// WithServer sets the Server field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Server field is set to the value of the last call.
func (b *VSpherePlatformFailureDomainSpecApplyConfiguration) WithServer(value string) *VSpherePlatformFailureDomainSpecApplyConfiguration {
	b.Server = &value
	return b
}

// WithTopology sets the Topology field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Topology field is set to the value of the last call.
func (b *VSpherePlatformFailureDomainSpecApplyConfiguration) WithTopology(value *VSpherePlatformTopologyApplyConfiguration) *VSpherePlatformFailureDomainSpecApplyConfiguration {
	b.Topology = value
	return b
}
