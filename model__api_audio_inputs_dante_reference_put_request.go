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

// checks if the ApiAudioInputsDanteReferencePutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAudioInputsDanteReferencePutRequest{}

// ApiAudioInputsDanteReferencePutRequest struct for ApiAudioInputsDanteReferencePutRequest
type ApiAudioInputsDanteReferencePutRequest struct {
	// The manual gain for the far end detection in dB. Must not be sent if `farEndAutoAdjustEnabled` is `true`.
	Gain *int32 `json:"gain,omitempty"`
	// Enable automatic threshold adjustment based on noise floor estimation. If `true` must not be sent together with `gain`.
	FarEndAutoAdjustEnabled *bool `json:"farEndAutoAdjustEnabled,omitempty"`
}

// NewApiAudioInputsDanteReferencePutRequest instantiates a new ApiAudioInputsDanteReferencePutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAudioInputsDanteReferencePutRequest() *ApiAudioInputsDanteReferencePutRequest {
	this := ApiAudioInputsDanteReferencePutRequest{}
	return &this
}

// NewApiAudioInputsDanteReferencePutRequestWithDefaults instantiates a new ApiAudioInputsDanteReferencePutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAudioInputsDanteReferencePutRequestWithDefaults() *ApiAudioInputsDanteReferencePutRequest {
	this := ApiAudioInputsDanteReferencePutRequest{}
	return &this
}

// GetGain returns the Gain field value if set, zero value otherwise.
func (o *ApiAudioInputsDanteReferencePutRequest) GetGain() int32 {
	if o == nil || IsNil(o.Gain) {
		var ret int32
		return ret
	}
	return *o.Gain
}

// GetGainOk returns a tuple with the Gain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsDanteReferencePutRequest) GetGainOk() (*int32, bool) {
	if o == nil || IsNil(o.Gain) {
		return nil, false
	}
	return o.Gain, true
}

// HasGain returns a boolean if a field has been set.
func (o *ApiAudioInputsDanteReferencePutRequest) HasGain() bool {
	if o != nil && !IsNil(o.Gain) {
		return true
	}

	return false
}

// SetGain gets a reference to the given int32 and assigns it to the Gain field.
func (o *ApiAudioInputsDanteReferencePutRequest) SetGain(v int32) {
	o.Gain = &v
}

// GetFarEndAutoAdjustEnabled returns the FarEndAutoAdjustEnabled field value if set, zero value otherwise.
func (o *ApiAudioInputsDanteReferencePutRequest) GetFarEndAutoAdjustEnabled() bool {
	if o == nil || IsNil(o.FarEndAutoAdjustEnabled) {
		var ret bool
		return ret
	}
	return *o.FarEndAutoAdjustEnabled
}

// GetFarEndAutoAdjustEnabledOk returns a tuple with the FarEndAutoAdjustEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioInputsDanteReferencePutRequest) GetFarEndAutoAdjustEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FarEndAutoAdjustEnabled) {
		return nil, false
	}
	return o.FarEndAutoAdjustEnabled, true
}

// HasFarEndAutoAdjustEnabled returns a boolean if a field has been set.
func (o *ApiAudioInputsDanteReferencePutRequest) HasFarEndAutoAdjustEnabled() bool {
	if o != nil && !IsNil(o.FarEndAutoAdjustEnabled) {
		return true
	}

	return false
}

// SetFarEndAutoAdjustEnabled gets a reference to the given bool and assigns it to the FarEndAutoAdjustEnabled field.
func (o *ApiAudioInputsDanteReferencePutRequest) SetFarEndAutoAdjustEnabled(v bool) {
	o.FarEndAutoAdjustEnabled = &v
}

func (o ApiAudioInputsDanteReferencePutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAudioInputsDanteReferencePutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gain) {
		toSerialize["gain"] = o.Gain
	}
	if !IsNil(o.FarEndAutoAdjustEnabled) {
		toSerialize["farEndAutoAdjustEnabled"] = o.FarEndAutoAdjustEnabled
	}
	return toSerialize, nil
}

type NullableApiAudioInputsDanteReferencePutRequest struct {
	value *ApiAudioInputsDanteReferencePutRequest
	isSet bool
}

func (v NullableApiAudioInputsDanteReferencePutRequest) Get() *ApiAudioInputsDanteReferencePutRequest {
	return v.value
}

func (v *NullableApiAudioInputsDanteReferencePutRequest) Set(val *ApiAudioInputsDanteReferencePutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAudioInputsDanteReferencePutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAudioInputsDanteReferencePutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAudioInputsDanteReferencePutRequest(val *ApiAudioInputsDanteReferencePutRequest) *NullableApiAudioInputsDanteReferencePutRequest {
	return &NullableApiAudioInputsDanteReferencePutRequest{value: val, isSet: true}
}

func (v NullableApiAudioInputsDanteReferencePutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAudioInputsDanteReferencePutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


