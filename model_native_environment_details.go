/*
Render Public API

Manage everything about your Render services

API version: 1.0.0
Contact: support@render.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package render

import (
	"encoding/json"
)

// checks if the NativeEnvironmentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NativeEnvironmentDetails{}

// NativeEnvironmentDetails struct for NativeEnvironmentDetails
type NativeEnvironmentDetails struct {
	BuildCommand *string `json:"buildCommand,omitempty"`
	StartCommand *string `json:"startCommand,omitempty"`
}

// NewNativeEnvironmentDetails instantiates a new NativeEnvironmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeEnvironmentDetails() *NativeEnvironmentDetails {
	this := NativeEnvironmentDetails{}
	return &this
}

// NewNativeEnvironmentDetailsWithDefaults instantiates a new NativeEnvironmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeEnvironmentDetailsWithDefaults() *NativeEnvironmentDetails {
	this := NativeEnvironmentDetails{}
	return &this
}

// GetBuildCommand returns the BuildCommand field value if set, zero value otherwise.
func (o *NativeEnvironmentDetails) GetBuildCommand() string {
	if o == nil || isNil(o.BuildCommand) {
		var ret string
		return ret
	}
	return *o.BuildCommand
}

// GetBuildCommandOk returns a tuple with the BuildCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeEnvironmentDetails) GetBuildCommandOk() (*string, bool) {
	if o == nil || isNil(o.BuildCommand) {
		return nil, false
	}
	return o.BuildCommand, true
}

// HasBuildCommand returns a boolean if a field has been set.
func (o *NativeEnvironmentDetails) HasBuildCommand() bool {
	if o != nil && !isNil(o.BuildCommand) {
		return true
	}

	return false
}

// SetBuildCommand gets a reference to the given string and assigns it to the BuildCommand field.
func (o *NativeEnvironmentDetails) SetBuildCommand(v string) {
	o.BuildCommand = &v
}

// GetStartCommand returns the StartCommand field value if set, zero value otherwise.
func (o *NativeEnvironmentDetails) GetStartCommand() string {
	if o == nil || isNil(o.StartCommand) {
		var ret string
		return ret
	}
	return *o.StartCommand
}

// GetStartCommandOk returns a tuple with the StartCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeEnvironmentDetails) GetStartCommandOk() (*string, bool) {
	if o == nil || isNil(o.StartCommand) {
		return nil, false
	}
	return o.StartCommand, true
}

// HasStartCommand returns a boolean if a field has been set.
func (o *NativeEnvironmentDetails) HasStartCommand() bool {
	if o != nil && !isNil(o.StartCommand) {
		return true
	}

	return false
}

// SetStartCommand gets a reference to the given string and assigns it to the StartCommand field.
func (o *NativeEnvironmentDetails) SetStartCommand(v string) {
	o.StartCommand = &v
}

func (o NativeEnvironmentDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NativeEnvironmentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BuildCommand) {
		toSerialize["buildCommand"] = o.BuildCommand
	}
	if !isNil(o.StartCommand) {
		toSerialize["startCommand"] = o.StartCommand
	}
	return toSerialize, nil
}

type NullableNativeEnvironmentDetails struct {
	value *NativeEnvironmentDetails
	isSet bool
}

func (v NullableNativeEnvironmentDetails) Get() *NativeEnvironmentDetails {
	return v.value
}

func (v *NullableNativeEnvironmentDetails) Set(val *NativeEnvironmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNativeEnvironmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNativeEnvironmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNativeEnvironmentDetails(val *NativeEnvironmentDetails) *NullableNativeEnvironmentDetails {
	return &NullableNativeEnvironmentDetails{value: val, isSet: true}
}

func (v NullableNativeEnvironmentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNativeEnvironmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


