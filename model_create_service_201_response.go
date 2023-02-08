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

// checks if the CreateService201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateService201Response{}

// CreateService201Response struct for CreateService201Response
type CreateService201Response struct {
	Service *Service `json:"service,omitempty"`
	DeployId *string `json:"deployId,omitempty"`
}

// NewCreateService201Response instantiates a new CreateService201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateService201Response() *CreateService201Response {
	this := CreateService201Response{}
	return &this
}

// NewCreateService201ResponseWithDefaults instantiates a new CreateService201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateService201ResponseWithDefaults() *CreateService201Response {
	this := CreateService201Response{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *CreateService201Response) GetService() Service {
	if o == nil || isNil(o.Service) {
		var ret Service
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateService201Response) GetServiceOk() (*Service, bool) {
	if o == nil || isNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *CreateService201Response) HasService() bool {
	if o != nil && !isNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given Service and assigns it to the Service field.
func (o *CreateService201Response) SetService(v Service) {
	o.Service = &v
}

// GetDeployId returns the DeployId field value if set, zero value otherwise.
func (o *CreateService201Response) GetDeployId() string {
	if o == nil || isNil(o.DeployId) {
		var ret string
		return ret
	}
	return *o.DeployId
}

// GetDeployIdOk returns a tuple with the DeployId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateService201Response) GetDeployIdOk() (*string, bool) {
	if o == nil || isNil(o.DeployId) {
		return nil, false
	}
	return o.DeployId, true
}

// HasDeployId returns a boolean if a field has been set.
func (o *CreateService201Response) HasDeployId() bool {
	if o != nil && !isNil(o.DeployId) {
		return true
	}

	return false
}

// SetDeployId gets a reference to the given string and assigns it to the DeployId field.
func (o *CreateService201Response) SetDeployId(v string) {
	o.DeployId = &v
}

func (o CreateService201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateService201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.DeployId) {
		toSerialize["deployId"] = o.DeployId
	}
	return toSerialize, nil
}

type NullableCreateService201Response struct {
	value *CreateService201Response
	isSet bool
}

func (v NullableCreateService201Response) Get() *CreateService201Response {
	return v.value
}

func (v *NullableCreateService201Response) Set(val *CreateService201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateService201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateService201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateService201Response(val *CreateService201Response) *NullableCreateService201Response {
	return &NullableCreateService201Response{value: val, isSet: true}
}

func (v NullableCreateService201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateService201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


