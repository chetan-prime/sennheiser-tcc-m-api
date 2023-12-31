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

// checks if the ApiAudioInputsDanteReferenceGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAudioInputsDanteReferenceGet200Response{}

// ApiAudioInputsDanteReferenceGet200Response struct for ApiAudioInputsDanteReferenceGet200Response
type ApiAudioInputsDanteReferenceGet200Response struct {
	// The manual gain for the far end detection in dB
	Gain *int32 `json:"gain,omitempty"`
	// Indication whether automatic threshold adjustment based on noise floor estimation is enabled
	FarEndAutoAdjustEnabled *bool `json:"farEndAutoAdjustEnabled,omitempty"`
}

// NewApiAudioInputsDanteReferenceGet200Response instantiates a new ApiAudioInputsDanteReferenceGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAudioInputsDanteReferenceGet200Response() *ApiAudioInputsDanteReferenceGet200Response {
	this := ApiAudioInputsDanteReferenceGet200Response{}
	var gain int32 = 0
	this.Gain = &gain
	var farEndAutoAdjustEnabled bool = true
	this.FarEndAutoAdjustEnabled = &farEndAutoAdjustEnabled
	return &this
}

// NewApiAudioInputsDanteReferenceGet200ResponseWithDefaults instantiates a new ApiAudioInputsDanteReferenceGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAudioInputsDanteReferenceGet200ResponseWithDefaults() *ApiAudioInputsDanteReferenceGet200Response {
	this := ApiAudioInputsDanteReferenceGet200Response{}
	var gain int32 = 0
	this.Gain = &gain
	var farEndAutoAdjustEnabled bool = true
	this.FarEndAutoAdjustEnabled = &farEndAutoAdjustEnabled
	return &this
}

// GetGain returns the Gain field value if set, zero value otherwise.
func (o *ApiAudioInputsDanteReferenceGet200Response) GetGain() int32 {
	if o == nil || IsNil(o.Gain) {
		var ret int32
		return ret
	}
	return *o.Gain
}

// GetGainOk returns a tuple with the Gain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsDanteReferenceGet200Response) GetGainOk() (*int32, bool) {
	if o == nil || IsNil(o.Gain) {
		return nil, false
	}
	return o.Gain, true
}

// HasGain returns a boolean if a field has been set.
func (o *ApiAudioInputsDanteReferenceGet200Response) HasGain() bool {
	if o != nil && !IsNil(o.Gain) {
		return true
	}

	return false
}

// SetGain gets a reference to the given int32 and assigns it to the Gain field.
func (o *ApiAudioInputsDanteReferenceGet200Response) SetGain(v int32) {
	o.Gain = &v
}

// GetFarEndAutoAdjustEnabled returns the FarEndAutoAdjustEnabled field value if set, zero value otherwise.
func (o *ApiAudioInputsDanteReferenceGet200Response) GetFarEndAutoAdjustEnabled() bool {
	if o == nil || IsNil(o.FarEndAutoAdjustEnabled) {
		var ret bool
		return ret
	}
	return *o.FarEndAutoAdjustEnabled
}

// GetFarEndAutoAdjustEnabledOk returns a tuple with the FarEndAutoAdjustEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsDanteReferenceGet200Response) GetFarEndAutoAdjustEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FarEndAutoAdjustEnabled) {
		return nil, false
	}
	return o.FarEndAutoAdjustEnabled, true
}

// HasFarEndAutoAdjustEnabled returns a boolean if a field has been set.
func (o *ApiAudioInputsDanteReferenceGet200Response) HasFarEndAutoAdjustEnabled() bool {
	if o != nil && !IsNil(o.FarEndAutoAdjustEnabled) {
		return true
	}

	return false
}

// SetFarEndAutoAdjustEnabled gets a reference to the given bool and assigns it to the FarEndAutoAdjustEnabled field.
func (o *ApiAudioInputsDanteReferenceGet200Response) SetFarEndAutoAdjustEnabled(v bool) {
	o.FarEndAutoAdjustEnabled = &v
}

func (o ApiAudioInputsDanteReferenceGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAudioInputsDanteReferenceGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gain) {
		toSerialize["gain"] = o.Gain
	}
	if !IsNil(o.FarEndAutoAdjustEnabled) {
		toSerialize["farEndAutoAdjustEnabled"] = o.FarEndAutoAdjustEnabled
	}
	return toSerialize, nil
}

type NullableApiAudioInputsDanteReferenceGet200Response struct {
	value *ApiAudioInputsDanteReferenceGet200Response
	isSet bool
}

func (v NullableApiAudioInputsDanteReferenceGet200Response) Get() *ApiAudioInputsDanteReferenceGet200Response {
	return v.value
}

func (v *NullableApiAudioInputsDanteReferenceGet200Response) Set(val *ApiAudioInputsDanteReferenceGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAudioInputsDanteReferenceGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAudioInputsDanteReferenceGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAudioInputsDanteReferenceGet200Response(val *ApiAudioInputsDanteReferenceGet200Response) *NullableApiAudioInputsDanteReferenceGet200Response {
	return &NullableApiAudioInputsDanteReferenceGet200Response{value: val, isSet: true}
}

func (v NullableApiAudioInputsDanteReferenceGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAudioInputsDanteReferenceGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


