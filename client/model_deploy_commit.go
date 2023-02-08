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
	"time"
)

// DeployCommit struct for DeployCommit
type DeployCommit struct {
	Id *string `json:"id,omitempty"`
	Message *string `json:"message,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewDeployCommit instantiates a new DeployCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployCommit() *DeployCommit {
	this := DeployCommit{}
	return &this
}

// NewDeployCommitWithDefaults instantiates a new DeployCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployCommitWithDefaults() *DeployCommit {
	this := DeployCommit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeployCommit) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployCommit) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeployCommit) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeployCommit) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeployCommit) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployCommit) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeployCommit) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeployCommit) SetMessage(v string) {
	o.Message = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DeployCommit) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployCommit) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DeployCommit) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DeployCommit) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o DeployCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableDeployCommit struct {
	value *DeployCommit
	isSet bool
}

func (v NullableDeployCommit) Get() *DeployCommit {
	return v.value
}

func (v *NullableDeployCommit) Set(val *DeployCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployCommit(val *DeployCommit) *NullableDeployCommit {
	return &NullableDeployCommit{value: val, isSet: true}
}

func (v NullableDeployCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


