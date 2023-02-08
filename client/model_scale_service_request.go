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

// ScaleServiceRequest struct for ScaleServiceRequest
type ScaleServiceRequest struct {
	NumInstances int32 `json:"numInstances"`
}

// NewScaleServiceRequest instantiates a new ScaleServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScaleServiceRequest(numInstances int32) *ScaleServiceRequest {
	this := ScaleServiceRequest{}
	this.NumInstances = numInstances
	return &this
}

// NewScaleServiceRequestWithDefaults instantiates a new ScaleServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScaleServiceRequestWithDefaults() *ScaleServiceRequest {
	this := ScaleServiceRequest{}
	return &this
}

// GetNumInstances returns the NumInstances field value
func (o *ScaleServiceRequest) GetNumInstances() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumInstances
}

// GetNumInstancesOk returns a tuple with the NumInstances field value
// and a boolean to check if the value has been set.
func (o *ScaleServiceRequest) GetNumInstancesOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NumInstances, true
}

// SetNumInstances sets field value
func (o *ScaleServiceRequest) SetNumInstances(v int32) {
	o.NumInstances = v
}

func (o ScaleServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["numInstances"] = o.NumInstances
	}
	return json.Marshal(toSerialize)
}

type NullableScaleServiceRequest struct {
	value *ScaleServiceRequest
	isSet bool
}

func (v NullableScaleServiceRequest) Get() *ScaleServiceRequest {
	return v.value
}

func (v *NullableScaleServiceRequest) Set(val *ScaleServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScaleServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScaleServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScaleServiceRequest(val *ScaleServiceRequest) *NullableScaleServiceRequest {
	return &NullableScaleServiceRequest{value: val, isSet: true}
}

func (v NullableScaleServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScaleServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


