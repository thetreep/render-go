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
	"fmt"
)

// ServicePATCHServiceDetails - struct for ServicePATCHServiceDetails
type ServicePATCHServiceDetails struct {
	BackgroundWorkerDetailsPATCH *BackgroundWorkerDetailsPATCH
	CronJobDetailsPATCH *CronJobDetailsPATCH
	PrivateServiceDetailsPATCH *PrivateServiceDetailsPATCH
	StaticSiteDetailsPATCH *StaticSiteDetailsPATCH
	WebServiceDetailsPATCH *WebServiceDetailsPATCH
}

// BackgroundWorkerDetailsPATCHAsServicePATCHServiceDetails is a convenience function that returns BackgroundWorkerDetailsPATCH wrapped in ServicePATCHServiceDetails
func BackgroundWorkerDetailsPATCHAsServicePATCHServiceDetails(v *BackgroundWorkerDetailsPATCH) ServicePATCHServiceDetails {
	return ServicePATCHServiceDetails{
		BackgroundWorkerDetailsPATCH: v,
	}
}

// CronJobDetailsPATCHAsServicePATCHServiceDetails is a convenience function that returns CronJobDetailsPATCH wrapped in ServicePATCHServiceDetails
func CronJobDetailsPATCHAsServicePATCHServiceDetails(v *CronJobDetailsPATCH) ServicePATCHServiceDetails {
	return ServicePATCHServiceDetails{
		CronJobDetailsPATCH: v,
	}
}

// PrivateServiceDetailsPATCHAsServicePATCHServiceDetails is a convenience function that returns PrivateServiceDetailsPATCH wrapped in ServicePATCHServiceDetails
func PrivateServiceDetailsPATCHAsServicePATCHServiceDetails(v *PrivateServiceDetailsPATCH) ServicePATCHServiceDetails {
	return ServicePATCHServiceDetails{
		PrivateServiceDetailsPATCH: v,
	}
}

// StaticSiteDetailsPATCHAsServicePATCHServiceDetails is a convenience function that returns StaticSiteDetailsPATCH wrapped in ServicePATCHServiceDetails
func StaticSiteDetailsPATCHAsServicePATCHServiceDetails(v *StaticSiteDetailsPATCH) ServicePATCHServiceDetails {
	return ServicePATCHServiceDetails{
		StaticSiteDetailsPATCH: v,
	}
}

// WebServiceDetailsPATCHAsServicePATCHServiceDetails is a convenience function that returns WebServiceDetailsPATCH wrapped in ServicePATCHServiceDetails
func WebServiceDetailsPATCHAsServicePATCHServiceDetails(v *WebServiceDetailsPATCH) ServicePATCHServiceDetails {
	return ServicePATCHServiceDetails{
		WebServiceDetailsPATCH: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServicePATCHServiceDetails) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BackgroundWorkerDetailsPATCH
	err = newStrictDecoder(data).Decode(&dst.BackgroundWorkerDetailsPATCH)
	if err == nil {
		jsonBackgroundWorkerDetailsPATCH, _ := json.Marshal(dst.BackgroundWorkerDetailsPATCH)
		if string(jsonBackgroundWorkerDetailsPATCH) == "{}" { // empty struct
			dst.BackgroundWorkerDetailsPATCH = nil
		} else {
			match++
		}
	} else {
		dst.BackgroundWorkerDetailsPATCH = nil
	}

	// try to unmarshal data into CronJobDetailsPATCH
	err = newStrictDecoder(data).Decode(&dst.CronJobDetailsPATCH)
	if err == nil {
		jsonCronJobDetailsPATCH, _ := json.Marshal(dst.CronJobDetailsPATCH)
		if string(jsonCronJobDetailsPATCH) == "{}" { // empty struct
			dst.CronJobDetailsPATCH = nil
		} else {
			match++
		}
	} else {
		dst.CronJobDetailsPATCH = nil
	}

	// try to unmarshal data into PrivateServiceDetailsPATCH
	err = newStrictDecoder(data).Decode(&dst.PrivateServiceDetailsPATCH)
	if err == nil {
		jsonPrivateServiceDetailsPATCH, _ := json.Marshal(dst.PrivateServiceDetailsPATCH)
		if string(jsonPrivateServiceDetailsPATCH) == "{}" { // empty struct
			dst.PrivateServiceDetailsPATCH = nil
		} else {
			match++
		}
	} else {
		dst.PrivateServiceDetailsPATCH = nil
	}

	// try to unmarshal data into StaticSiteDetailsPATCH
	err = newStrictDecoder(data).Decode(&dst.StaticSiteDetailsPATCH)
	if err == nil {
		jsonStaticSiteDetailsPATCH, _ := json.Marshal(dst.StaticSiteDetailsPATCH)
		if string(jsonStaticSiteDetailsPATCH) == "{}" { // empty struct
			dst.StaticSiteDetailsPATCH = nil
		} else {
			match++
		}
	} else {
		dst.StaticSiteDetailsPATCH = nil
	}

	// try to unmarshal data into WebServiceDetailsPATCH
	err = newStrictDecoder(data).Decode(&dst.WebServiceDetailsPATCH)
	if err == nil {
		jsonWebServiceDetailsPATCH, _ := json.Marshal(dst.WebServiceDetailsPATCH)
		if string(jsonWebServiceDetailsPATCH) == "{}" { // empty struct
			dst.WebServiceDetailsPATCH = nil
		} else {
			match++
		}
	} else {
		dst.WebServiceDetailsPATCH = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BackgroundWorkerDetailsPATCH = nil
		dst.CronJobDetailsPATCH = nil
		dst.PrivateServiceDetailsPATCH = nil
		dst.StaticSiteDetailsPATCH = nil
		dst.WebServiceDetailsPATCH = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ServicePATCHServiceDetails)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ServicePATCHServiceDetails)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServicePATCHServiceDetails) MarshalJSON() ([]byte, error) {
	if src.BackgroundWorkerDetailsPATCH != nil {
		return json.Marshal(&src.BackgroundWorkerDetailsPATCH)
	}

	if src.CronJobDetailsPATCH != nil {
		return json.Marshal(&src.CronJobDetailsPATCH)
	}

	if src.PrivateServiceDetailsPATCH != nil {
		return json.Marshal(&src.PrivateServiceDetailsPATCH)
	}

	if src.StaticSiteDetailsPATCH != nil {
		return json.Marshal(&src.StaticSiteDetailsPATCH)
	}

	if src.WebServiceDetailsPATCH != nil {
		return json.Marshal(&src.WebServiceDetailsPATCH)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServicePATCHServiceDetails) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BackgroundWorkerDetailsPATCH != nil {
		return obj.BackgroundWorkerDetailsPATCH
	}

	if obj.CronJobDetailsPATCH != nil {
		return obj.CronJobDetailsPATCH
	}

	if obj.PrivateServiceDetailsPATCH != nil {
		return obj.PrivateServiceDetailsPATCH
	}

	if obj.StaticSiteDetailsPATCH != nil {
		return obj.StaticSiteDetailsPATCH
	}

	if obj.WebServiceDetailsPATCH != nil {
		return obj.WebServiceDetailsPATCH
	}

	// all schemas are nil
	return nil
}

type NullableServicePATCHServiceDetails struct {
	value *ServicePATCHServiceDetails
	isSet bool
}

func (v NullableServicePATCHServiceDetails) Get() *ServicePATCHServiceDetails {
	return v.value
}

func (v *NullableServicePATCHServiceDetails) Set(val *ServicePATCHServiceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePATCHServiceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePATCHServiceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePATCHServiceDetails(val *ServicePATCHServiceDetails) *NullableServicePATCHServiceDetails {
	return &NullableServicePATCHServiceDetails{value: val, isSet: true}
}

func (v NullableServicePATCHServiceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePATCHServiceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


