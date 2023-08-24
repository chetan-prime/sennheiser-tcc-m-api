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

// checks if the ApiAudioOutputsDanteFarEndPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAudioOutputsDanteFarEndPutRequest{}

// ApiAudioOutputsDanteFarEndPutRequest struct for ApiAudioOutputsDanteFarEndPutRequest
type ApiAudioOutputsDanteFarEndPutRequest struct {
	// Gain in dB
	Gain *int32 `json:"gain,omitempty"`
	// Applies the settings from /api/audio/noiseGate
	NoiseGateEnabled *bool `json:"noiseGateEnabled,omitempty"`
	// Applies the settings from /api/audio/equalizer
	EqualizerEnabled *bool `json:"equalizerEnabled,omitempty"`
	// Delay in ms
	Delay *int32 `json:"delay,omitempty"`
}

// NewApiAudioOutputsDanteFarEndPutRequest instantiates a new ApiAudioOutputsDanteFarEndPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAudioOutputsDanteFarEndPutRequest() *ApiAudioOutputsDanteFarEndPutRequest {
	this := ApiAudioOutputsDanteFarEndPutRequest{}
	return &this
}

// NewApiAudioOutputsDanteFarEndPutRequestWithDefaults instantiates a new ApiAudioOutputsDanteFarEndPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAudioOutputsDanteFarEndPutRequestWithDefaults() *ApiAudioOutputsDanteFarEndPutRequest {
	this := ApiAudioOutputsDanteFarEndPutRequest{}
	return &this
}

// GetGain returns the Gain field value if set, zero value otherwise.
func (o *ApiAudioOutputsDanteFarEndPutRequest) GetGain() int32 {
	if o == nil || IsNil(o.Gain) {
		var ret int32
		return ret
	}
	return *o.Gain
}

// GetGainOk returns a tuple with the Gain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioOutputsDanteFarEndPutRequest) GetGainOk() (*int32, bool) {
	if o == nil || IsNil(o.Gain) {
		return nil, false
	}
	return o.Gain, true
}

// HasGain returns a boolean if a field has been set.
func (o *ApiAudioOutputsDanteFarEndPutRequest) HasGain() bool {
	if o != nil && !IsNil(o.Gain) {
		return true
	}

	return false
}

// SetGain gets a reference to the given int32 and assigns it to the Gain field.
func (o *ApiAudioOutputsDanteFarEndPutRequest) SetGain(v int32) {
	o.Gain = &v
}

// GetNoiseGateEnabled returns the NoiseGateEnabled field value if set, zero value otherwise.
func (o *ApiAudioOutputsDanteFarEndPutRequest) GetNoiseGateEnabled() bool {
	if o == nil || IsNil(o.NoiseGateEnabled) {
		var ret bool
		return ret
	}
	return *o.NoiseGateEnabled
}

// GetNoiseGateEnabledOk returns a tuple with the NoiseGateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioOutputsDanteFarEndPutRequest) GetNoiseGateEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NoiseGateEnabled) {
		return nil, false
	}
	return o.NoiseGateEnabled, true
}

// HasNoiseGateEnabled returns a boolean if a field has been set.
func (o *ApiAudioOutputsDanteFarEndPutRequest) HasNoiseGateEnabled() bool {
	if o != nil && !IsNil(o.NoiseGateEnabled) {
		return true
	}

	return false
}

// SetNoiseGateEnabled gets a reference to the given bool and assigns it to the NoiseGateEnabled field.
func (o *ApiAudioOutputsDanteFarEndPutRequest) SetNoiseGateEnabled(v bool) {
	o.NoiseGateEnabled = &v
}

// GetEqualizerEnabled returns the EqualizerEnabled field value if set, zero value otherwise.
func (o *ApiAudioOutputsDanteFarEndPutRequest) GetEqualizerEnabled() bool {
	if o == nil || IsNil(o.EqualizerEnabled) {
		var ret bool
		return ret
	}
	return *o.EqualizerEnabled
}

// GetEqualizerEnabledOk returns a tuple with the EqualizerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioOutputsDanteFarEndPutRequest) GetEqualizerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EqualizerEnabled) {
		return nil, false
	}
	return o.EqualizerEnabled, true
}

// HasEqualizerEnabled returns a boolean if a field has been set.
func (o *ApiAudioOutputsDanteFarEndPutRequest) HasEqualizerEnabled() bool {
	if o != nil && !IsNil(o.EqualizerEnabled) {
		return true
	}

	return false
}

// SetEqualizerEnabled gets a reference to the given bool and assigns it to the EqualizerEnabled field.
func (o *ApiAudioOutputsDanteFarEndPutRequest) SetEqualizerEnabled(v bool) {
	o.EqualizerEnabled = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *ApiAudioOutputsDanteFarEndPutRequest) GetDelay() int32 {
	if o == nil || IsNil(o.Delay) {
		var ret int32
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioOutputsDanteFarEndPutRequest) GetDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.Delay) {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *ApiAudioOutputsDanteFarEndPutRequest) HasDelay() bool {
	if o != nil && !IsNil(o.Delay) {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int32 and assigns it to the Delay field.
func (o *ApiAudioOutputsDanteFarEndPutRequest) SetDelay(v int32) {
	o.Delay = &v
}

func (o ApiAudioOutputsDanteFarEndPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAudioOutputsDanteFarEndPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gain) {
		toSerialize["gain"] = o.Gain
	}
	if !IsNil(o.NoiseGateEnabled) {
		toSerialize["noiseGateEnabled"] = o.NoiseGateEnabled
	}
	if !IsNil(o.EqualizerEnabled) {
		toSerialize["equalizerEnabled"] = o.EqualizerEnabled
	}
	if !IsNil(o.Delay) {
		toSerialize["delay"] = o.Delay
	}
	return toSerialize, nil
}

type NullableApiAudioOutputsDanteFarEndPutRequest struct {
	value *ApiAudioOutputsDanteFarEndPutRequest
	isSet bool
}

func (v NullableApiAudioOutputsDanteFarEndPutRequest) Get() *ApiAudioOutputsDanteFarEndPutRequest {
	return v.value
}

func (v *NullableApiAudioOutputsDanteFarEndPutRequest) Set(val *ApiAudioOutputsDanteFarEndPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAudioOutputsDanteFarEndPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAudioOutputsDanteFarEndPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAudioOutputsDanteFarEndPutRequest(val *ApiAudioOutputsDanteFarEndPutRequest) *NullableApiAudioOutputsDanteFarEndPutRequest {
	return &NullableApiAudioOutputsDanteFarEndPutRequest{value: val, isSet: true}
}

func (v NullableApiAudioOutputsDanteFarEndPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAudioOutputsDanteFarEndPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

