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

// checks if the ApiAudioInputsReferenceLevelGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAudioInputsReferenceLevelGet200Response{}

// ApiAudioInputsReferenceLevelGet200Response struct for ApiAudioInputsReferenceLevelGet200Response
type ApiAudioInputsReferenceLevelGet200Response struct {
	// RMS level in dBfs of the digital reference input
	Rms *int32 `json:"rms,omitempty"`
}

// NewApiAudioInputsReferenceLevelGet200Response instantiates a new ApiAudioInputsReferenceLevelGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAudioInputsReferenceLevelGet200Response() *ApiAudioInputsReferenceLevelGet200Response {
	this := ApiAudioInputsReferenceLevelGet200Response{}
	return &this
}

// NewApiAudioInputsReferenceLevelGet200ResponseWithDefaults instantiates a new ApiAudioInputsReferenceLevelGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAudioInputsReferenceLevelGet200ResponseWithDefaults() *ApiAudioInputsReferenceLevelGet200Response {
	this := ApiAudioInputsReferenceLevelGet200Response{}
	return &this
}

// GetRms returns the Rms field value if set, zero value otherwise.
func (o *ApiAudioInputsReferenceLevelGet200Response) GetRms() int32 {
	if o == nil || IsNil(o.Rms) {
		var ret int32
		return ret
	}
	return *o.Rms
}

// GetRmsOk returns a tuple with the Rms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsReferenceLevelGet200Response) GetRmsOk() (*int32, bool) {
	if o == nil || IsNil(o.Rms) {
		return nil, false
	}
	return o.Rms, true
}

// HasRms returns a boolean if a field has been set.
func (o *ApiAudioInputsReferenceLevelGet200Response) HasRms() bool {
	if o != nil && !IsNil(o.Rms) {
		return true
	}

	return false
}

// SetRms gets a reference to the given int32 and assigns it to the Rms field.
func (o *ApiAudioInputsReferenceLevelGet200Response) SetRms(v int32) {
	o.Rms = &v
}

func (o ApiAudioInputsReferenceLevelGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAudioInputsReferenceLevelGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rms) {
		toSerialize["rms"] = o.Rms
	}
	return toSerialize, nil
}

type NullableApiAudioInputsReferenceLevelGet200Response struct {
	value *ApiAudioInputsReferenceLevelGet200Response
	isSet bool
}

func (v NullableApiAudioInputsReferenceLevelGet200Response) Get() *ApiAudioInputsReferenceLevelGet200Response {
	return v.value
}

func (v *NullableApiAudioInputsReferenceLevelGet200Response) Set(val *ApiAudioInputsReferenceLevelGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAudioInputsReferenceLevelGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAudioInputsReferenceLevelGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAudioInputsReferenceLevelGet200Response(val *ApiAudioInputsReferenceLevelGet200Response) *NullableApiAudioInputsReferenceLevelGet200Response {
	return &NullableApiAudioInputsReferenceLevelGet200Response{value: val, isSet: true}
}

func (v NullableApiAudioInputsReferenceLevelGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAudioInputsReferenceLevelGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


