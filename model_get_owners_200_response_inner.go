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

// checks if the GetOwners200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOwners200ResponseInner{}

// GetOwners200ResponseInner struct for GetOwners200ResponseInner
type GetOwners200ResponseInner struct {
	Owner *Owner `json:"owner,omitempty"`
	Cursor *string `json:"cursor,omitempty"`
}

// NewGetOwners200ResponseInner instantiates a new GetOwners200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOwners200ResponseInner() *GetOwners200ResponseInner {
	this := GetOwners200ResponseInner{}
	return &this
}

// NewGetOwners200ResponseInnerWithDefaults instantiates a new GetOwners200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOwners200ResponseInnerWithDefaults() *GetOwners200ResponseInner {
	this := GetOwners200ResponseInner{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *GetOwners200ResponseInner) GetOwner() Owner {
	if o == nil || isNil(o.Owner) {
		var ret Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOwners200ResponseInner) GetOwnerOk() (*Owner, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *GetOwners200ResponseInner) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Owner and assigns it to the Owner field.
func (o *GetOwners200ResponseInner) SetOwner(v Owner) {
	o.Owner = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *GetOwners200ResponseInner) GetCursor() string {
	if o == nil || isNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOwners200ResponseInner) GetCursorOk() (*string, bool) {
	if o == nil || isNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *GetOwners200ResponseInner) HasCursor() bool {
	if o != nil && !isNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *GetOwners200ResponseInner) SetCursor(v string) {
	o.Cursor = &v
}

func (o GetOwners200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOwners200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	return toSerialize, nil
}

type NullableGetOwners200ResponseInner struct {
	value *GetOwners200ResponseInner
	isSet bool
}

func (v NullableGetOwners200ResponseInner) Get() *GetOwners200ResponseInner {
	return v.value
}

func (v *NullableGetOwners200ResponseInner) Set(val *GetOwners200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOwners200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOwners200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOwners200ResponseInner(val *GetOwners200ResponseInner) *NullableGetOwners200ResponseInner {
	return &NullableGetOwners200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOwners200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOwners200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


