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

// checks if the ApiAudioOutputsDanteFarEndGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAudioOutputsDanteFarEndGet200Response{}

// ApiAudioOutputsDanteFarEndGet200Response struct for ApiAudioOutputsDanteFarEndGet200Response
type ApiAudioOutputsDanteFarEndGet200Response struct {
	// Gain in dB
	Gain *int32 `json:"gain,omitempty"`
	// Indication whether the settings from /api/audio/noiseGate are applied for this output
	NoiseGateEnabled *bool `json:"noiseGateEnabled,omitempty"`
	// Indication whether the settings from /api/audio/equalizer are applied for this output
	EqualizerEnabled *bool `json:"equalizerEnabled,omitempty"`
	// Delay in ms
	Delay *int32 `json:"delay,omitempty"`
}

// NewApiAudioOutputsDanteFarEndGet200Response instantiates a new ApiAudioOutputsDanteFarEndGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAudioOutputsDanteFarEndGet200Response() *ApiAudioOutputsDanteFarEndGet200Response {
	this := ApiAudioOutputsDanteFarEndGet200Response{}
	var gain int32 = 12
	this.Gain = &gain
	var noiseGateEnabled bool = false
	this.NoiseGateEnabled = &noiseGateEnabled
	var equalizerEnabled bool = false
	this.EqualizerEnabled = &equalizerEnabled
	var delay int32 = 0
	this.Delay = &delay
	return &this
}

// NewApiAudioOutputsDanteFarEndGet200ResponseWithDefaults instantiates a new ApiAudioOutputsDanteFarEndGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAudioOutputsDanteFarEndGet200ResponseWithDefaults() *ApiAudioOutputsDanteFarEndGet200Response {
	this := ApiAudioOutputsDanteFarEndGet200Response{}
	var gain int32 = 12
	this.Gain = &gain
	var noiseGateEnabled bool = false
	this.NoiseGateEnabled = &noiseGateEnabled
	var equalizerEnabled bool = false
	this.EqualizerEnabled = &equalizerEnabled
	var delay int32 = 0
	this.Delay = &delay
	return &this
}

// GetGain returns the Gain field value if set, zero value otherwise.
func (o *ApiAudioOutputsDanteFarEndGet200Response) GetGain() int32 {
	if o == nil || IsNil(o.Gain) {
		var ret int32
		return ret
	}
	return *o.Gain
}

// GetGainOk returns a tuple with the Gain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioOutputsDanteFarEndGet200Response) GetGainOk() (*int32, bool) {
	if o == nil || IsNil(o.Gain) {
		return nil, false
	}
	return o.Gain, true
}

// HasGain returns a boolean if a field has been set.
func (o *ApiAudioOutputsDanteFarEndGet200Response) HasGain() bool {
	if o != nil && !IsNil(o.Gain) {
		return true
	}

	return false
}

// SetGain gets a reference to the given int32 and assigns it to the Gain field.
func (o *ApiAudioOutputsDanteFarEndGet200Response) SetGain(v int32) {
	o.Gain = &v
}

// GetNoiseGateEnabled returns the NoiseGateEnabled field value if set, zero value otherwise.
func (o *ApiAudioOutputsDanteFarEndGet200Response) GetNoiseGateEnabled() bool {
	if o == nil || IsNil(o.NoiseGateEnabled) {
		var ret bool
		return ret
	}
	return *o.NoiseGateEnabled
}

// GetNoiseGateEnabledOk returns a tuple with the NoiseGateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioOutputsDanteFarEndGet200Response) GetNoiseGateEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NoiseGateEnabled) {
		return nil, false
	}
	return o.NoiseGateEnabled, true
}

// HasNoiseGateEnabled returns a boolean if a field has been set.
func (o *ApiAudioOutputsDanteFarEndGet200Response) HasNoiseGateEnabled() bool {
	if o != nil && !IsNil(o.NoiseGateEnabled) {
		return true
	}

	return false
}

// SetNoiseGateEnabled gets a reference to the given bool and assigns it to the NoiseGateEnabled field.
func (o *ApiAudioOutputsDanteFarEndGet200Response) SetNoiseGateEnabled(v bool) {
	o.NoiseGateEnabled = &v
}

// GetEqualizerEnabled returns the EqualizerEnabled field value if set, zero value otherwise.
func (o *ApiAudioOutputsDanteFarEndGet200Response) GetEqualizerEnabled() bool {
	if o == nil || IsNil(o.EqualizerEnabled) {
		var ret bool
		return ret
	}
	return *o.EqualizerEnabled
}

// GetEqualizerEnabledOk returns a tuple with the EqualizerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioOutputsDanteFarEndGet200Response) GetEqualizerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EqualizerEnabled) {
		return nil, false
	}
	return o.EqualizerEnabled, true
}

// HasEqualizerEnabled returns a boolean if a field has been set.
func (o *ApiAudioOutputsDanteFarEndGet200Response) HasEqualizerEnabled() bool {
	if o != nil && !IsNil(o.EqualizerEnabled) {
		return true
	}

	return false
}

// SetEqualizerEnabled gets a reference to the given bool and assigns it to the EqualizerEnabled field.
func (o *ApiAudioOutputsDanteFarEndGet200Response) SetEqualizerEnabled(v bool) {
	o.EqualizerEnabled = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *ApiAudioOutputsDanteFarEndGet200Response) GetDelay() int32 {
	if o == nil || IsNil(o.Delay) {
		var ret int32
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioOutputsDanteFarEndGet200Response) GetDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.Delay) {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *ApiAudioOutputsDanteFarEndGet200Response) HasDelay() bool {
	if o != nil && !IsNil(o.Delay) {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int32 and assigns it to the Delay field.
func (o *ApiAudioOutputsDanteFarEndGet200Response) SetDelay(v int32) {
	o.Delay = &v
}

func (o ApiAudioOutputsDanteFarEndGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAudioOutputsDanteFarEndGet200Response) ToMap() (map[string]interface{}, error) {
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

type NullableApiAudioOutputsDanteFarEndGet200Response struct {
	value *ApiAudioOutputsDanteFarEndGet200Response
	isSet bool
}

func (v NullableApiAudioOutputsDanteFarEndGet200Response) Get() *ApiAudioOutputsDanteFarEndGet200Response {
	return v.value
}

func (v *NullableApiAudioOutputsDanteFarEndGet200Response) Set(val *ApiAudioOutputsDanteFarEndGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAudioOutputsDanteFarEndGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAudioOutputsDanteFarEndGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAudioOutputsDanteFarEndGet200Response(val *ApiAudioOutputsDanteFarEndGet200Response) *NullableApiAudioOutputsDanteFarEndGet200Response {
	return &NullableApiAudioOutputsDanteFarEndGet200Response{value: val, isSet: true}
}

func (v NullableApiAudioOutputsDanteFarEndGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAudioOutputsDanteFarEndGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


