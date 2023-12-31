/*
TCC M 3rd party API Release

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiDeviceIdentificationPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDeviceIdentificationPutRequest{}

// ApiDeviceIdentificationPutRequest struct for ApiDeviceIdentificationPutRequest
type ApiDeviceIdentificationPutRequest struct {
	Visual *bool `json:"visual,omitempty"`
}

// NewApiDeviceIdentificationPutRequest instantiates a new ApiDeviceIdentificationPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDeviceIdentificationPutRequest() *ApiDeviceIdentificationPutRequest {
	this := ApiDeviceIdentificationPutRequest{}
	return &this
}

// NewApiDeviceIdentificationPutRequestWithDefaults instantiates a new ApiDeviceIdentificationPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDeviceIdentificationPutRequestWithDefaults() *ApiDeviceIdentificationPutRequest {
	this := ApiDeviceIdentificationPutRequest{}
	return &this
}

// GetVisual returns the Visual field value if set, zero value otherwise.
func (o *ApiDeviceIdentificationPutRequest) GetVisual() bool {
	if o == nil || IsNil(o.Visual) {
		var ret bool
		return ret
	}
	return *o.Visual
}

// GetVisualOk returns a tuple with the Visual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDeviceIdentificationPutRequest) GetVisualOk() (*bool, bool) {
	if o == nil || IsNil(o.Visual) {
		return nil, false
	}
	return o.Visual, true
}

// HasVisual returns a boolean if a field has been set.
func (o *ApiDeviceIdentificationPutRequest) HasVisual() bool {
	if o != nil && !IsNil(o.Visual) {
		return true
	}

	return false
}

// SetVisual gets a reference to the given bool and assigns it to the Visual field.
func (o *ApiDeviceIdentificationPutRequest) SetVisual(v bool) {
	o.Visual = &v
}

func (o ApiDeviceIdentificationPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDeviceIdentificationPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Visual) {
		toSerialize["visual"] = o.Visual
	}
	return toSerialize, nil
}

type NullableApiDeviceIdentificationPutRequest struct {
	value *ApiDeviceIdentificationPutRequest
	isSet bool
}

func (v NullableApiDeviceIdentificationPutRequest) Get() *ApiDeviceIdentificationPutRequest {
	return v.value
}

func (v *NullableApiDeviceIdentificationPutRequest) Set(val *ApiDeviceIdentificationPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDeviceIdentificationPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDeviceIdentificationPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDeviceIdentificationPutRequest(val *ApiDeviceIdentificationPutRequest) *NullableApiDeviceIdentificationPutRequest {
	return &NullableApiDeviceIdentificationPutRequest{value: val, isSet: true}
}

func (v NullableApiDeviceIdentificationPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDeviceIdentificationPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


