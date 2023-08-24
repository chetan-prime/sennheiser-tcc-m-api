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

// checks if the ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth{}

// ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth struct for ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth
type ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth struct {
	// Angle in degrees
	Min *int32 `json:"min,omitempty"`
	// Angle in degrees
	Max *int32 `json:"max,omitempty"`
}

// NewApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth instantiates a new ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth() *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth {
	this := ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth{}
	return &this
}

// NewApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuthWithDefaults instantiates a new ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuthWithDefaults() *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth {
	this := ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) SetMax(v int32) {
	o.Max = &v
}

func (o ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth struct {
	value *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth
	isSet bool
}

func (v NullableApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) Get() *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth {
	return v.value
}

func (v *NullableApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) Set(val *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth(val *ApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) *NullableApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth {
	return &NullableApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth{value: val, isSet: true}
}

func (v NullableApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAudioInputsMicrophoneExclusionZonesGet200ResponseInnerAzimuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


