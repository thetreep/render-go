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

// checks if the Route type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Route{}

// Route struct for Route
type Route struct {
	Type string `json:"type"`
	Source string `json:"source"`
	Destination string `json:"destination"`
}

// NewRoute instantiates a new Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute(type_ string, source string, destination string) *Route {
	this := Route{}
	this.Type = type_
	this.Source = source
	this.Destination = destination
	return &this
}

// NewRouteWithDefaults instantiates a new Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteWithDefaults() *Route {
	this := Route{}
	return &this
}

// GetType returns the Type field value
func (o *Route) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Route) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Route) SetType(v string) {
	o.Type = v
}

// GetSource returns the Source field value
func (o *Route) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *Route) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *Route) SetSource(v string) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *Route) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *Route) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *Route) SetDestination(v string) {
	o.Destination = v
}

func (o Route) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Route) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["source"] = o.Source
	toSerialize["destination"] = o.Destination
	return toSerialize, nil
}

type NullableRoute struct {
	value *Route
	isSet bool
}

func (v NullableRoute) Get() *Route {
	return v.value
}

func (v *NullableRoute) Set(val *Route) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute(val *Route) *NullableRoute {
	return &NullableRoute{value: val, isSet: true}
}

func (v NullableRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


