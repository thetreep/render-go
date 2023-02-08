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

// checks if the ServerPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerPort{}

// ServerPort struct for ServerPort
type ServerPort struct {
	Port *int32 `json:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
}

// NewServerPort instantiates a new ServerPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerPort() *ServerPort {
	this := ServerPort{}
	return &this
}

// NewServerPortWithDefaults instantiates a new ServerPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerPortWithDefaults() *ServerPort {
	this := ServerPort{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ServerPort) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerPort) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ServerPort) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ServerPort) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ServerPort) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerPort) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ServerPort) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ServerPort) SetProtocol(v string) {
	o.Protocol = &v
}

func (o ServerPort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	return toSerialize, nil
}

type NullableServerPort struct {
	value *ServerPort
	isSet bool
}

func (v NullableServerPort) Get() *ServerPort {
	return v.value
}

func (v *NullableServerPort) Set(val *ServerPort) {
	v.value = val
	v.isSet = true
}

func (v NullableServerPort) IsSet() bool {
	return v.isSet
}

func (v *NullableServerPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerPort(val *ServerPort) *NullableServerPort {
	return &NullableServerPort{value: val, isSet: true}
}

func (v NullableServerPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


