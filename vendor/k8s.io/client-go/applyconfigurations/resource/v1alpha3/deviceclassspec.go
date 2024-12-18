/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha3

// DeviceClassSpecApplyConfiguration represents a declarative configuration of the DeviceClassSpec type for use
// with apply.
type DeviceClassSpecApplyConfiguration struct {
	Selectors []DeviceSelectorApplyConfiguration           `json:"selectors,omitempty"`
	Config    []DeviceClassConfigurationApplyConfiguration `json:"config,omitempty"`
}

// DeviceClassSpecApplyConfiguration constructs a declarative configuration of the DeviceClassSpec type for use with
// apply.
func DeviceClassSpec() *DeviceClassSpecApplyConfiguration {
	return &DeviceClassSpecApplyConfiguration{}
}

// WithSelectors adds the given value to the Selectors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Selectors field.
func (b *DeviceClassSpecApplyConfiguration) WithSelectors(values ...*DeviceSelectorApplyConfiguration) *DeviceClassSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSelectors")
		}
		b.Selectors = append(b.Selectors, *values[i])
	}
	return b
}

// WithConfig adds the given value to the Config field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Config field.
func (b *DeviceClassSpecApplyConfiguration) WithConfig(values ...*DeviceClassConfigurationApplyConfiguration) *DeviceClassSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConfig")
		}
		b.Config = append(b.Config, *values[i])
	}
	return b
}
