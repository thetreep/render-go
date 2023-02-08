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

// checks if the Deploy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Deploy{}

// Deploy struct for Deploy
type Deploy struct {
	Id string `json:"id"`
	Commit *DeployCommit `json:"commit,omitempty"`
	Status *string `json:"status,omitempty"`
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewDeploy instantiates a new Deploy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploy(id string) *Deploy {
	this := Deploy{}
	this.Id = id
	return &this
}

// NewDeployWithDefaults instantiates a new Deploy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployWithDefaults() *Deploy {
	this := Deploy{}
	return &this
}

// GetId returns the Id field value
func (o *Deploy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Deploy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Deploy) SetId(v string) {
	o.Id = v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *Deploy) GetCommit() DeployCommit {
	if o == nil || isNil(o.Commit) {
		var ret DeployCommit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deploy) GetCommitOk() (*DeployCommit, bool) {
	if o == nil || isNil(o.Commit) {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *Deploy) HasCommit() bool {
	if o != nil && !isNil(o.Commit) {
		return true
	}

	return false
}

// SetCommit gets a reference to the given DeployCommit and assigns it to the Commit field.
func (o *Deploy) SetCommit(v DeployCommit) {
	o.Commit = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Deploy) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deploy) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Deploy) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Deploy) SetStatus(v string) {
	o.Status = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *Deploy) GetFinishedAt() time.Time {
	if o == nil || isNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deploy) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *Deploy) HasFinishedAt() bool {
	if o != nil && !isNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *Deploy) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Deploy) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deploy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Deploy) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Deploy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Deploy) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deploy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Deploy) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Deploy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Deploy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Deploy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !isNil(o.Commit) {
		toSerialize["commit"] = o.Commit
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.FinishedAt) {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableDeploy struct {
	value *Deploy
	isSet bool
}

func (v NullableDeploy) Get() *Deploy {
	return v.value
}

func (v *NullableDeploy) Set(val *Deploy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploy(val *Deploy) *NullableDeploy {
	return &NullableDeploy{value: val, isSet: true}
}

func (v NullableDeploy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


