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

// checks if the ApiAudioNoiseGatePutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAudioNoiseGatePutRequest{}

// ApiAudioNoiseGatePutRequest struct for ApiAudioNoiseGatePutRequest
type ApiAudioNoiseGatePutRequest struct {
	// Threshold in dB
	Threshold *int32 `json:"threshold,omitempty"`
	// Hold time in ms
	HoldTime *int32 `json:"holdTime,omitempty"`
}

// NewApiAudioNoiseGatePutRequest instantiates a new ApiAudioNoiseGatePutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAudioNoiseGatePutRequest() *ApiAudioNoiseGatePutRequest {
	this := ApiAudioNoiseGatePutRequest{}
	return &this
}

// NewApiAudioNoiseGatePutRequestWithDefaults instantiates a new ApiAudioNoiseGatePutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAudioNoiseGatePutRequestWithDefaults() *ApiAudioNoiseGatePutRequest {
	this := ApiAudioNoiseGatePutRequest{}
	return &this
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *ApiAudioNoiseGatePutRequest) GetThreshold() int32 {
	if o == nil || IsNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioNoiseGatePutRequest) GetThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *ApiAudioNoiseGatePutRequest) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *ApiAudioNoiseGatePutRequest) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetHoldTime returns the HoldTime field value if set, zero value otherwise.
func (o *ApiAudioNoiseGatePutRequest) GetHoldTime() int32 {
	if o == nil || IsNil(o.HoldTime) {
		var ret int32
		return ret
	}
	return *o.HoldTime
}

// GetHoldTimeOk returns a tuple with the HoldTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioNoiseGatePutRequest) GetHoldTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.HoldTime) {
		return nil, false
	}
	return o.HoldTime, true
}

// HasHoldTime returns a boolean if a field has been set.
func (o *ApiAudioNoiseGatePutRequest) HasHoldTime() bool {
	if o != nil && !IsNil(o.HoldTime) {
		return true
	}

	return false
}

// SetHoldTime gets a reference to the given int32 and assigns it to the HoldTime field.
func (o *ApiAudioNoiseGatePutRequest) SetHoldTime(v int32) {
	o.HoldTime = &v
}

func (o ApiAudioNoiseGatePutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAudioNoiseGatePutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.HoldTime) {
		toSerialize["holdTime"] = o.HoldTime
	}
	return toSerialize, nil
}

type NullableApiAudioNoiseGatePutRequest struct {
	value *ApiAudioNoiseGatePutRequest
	isSet bool
}

func (v NullableApiAudioNoiseGatePutRequest) Get() *ApiAudioNoiseGatePutRequest {
	return v.value
}

func (v *NullableApiAudioNoiseGatePutRequest) Set(val *ApiAudioNoiseGatePutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAudioNoiseGatePutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAudioNoiseGatePutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAudioNoiseGatePutRequest(val *ApiAudioNoiseGatePutRequest) *NullableApiAudioNoiseGatePutRequest {
	return &NullableApiAudioNoiseGatePutRequest{value: val, isSet: true}
}

func (v NullableApiAudioNoiseGatePutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAudioNoiseGatePutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

