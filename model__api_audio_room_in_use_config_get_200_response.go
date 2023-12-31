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

// checks if the ApiAudioRoomInUseConfigGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAudioRoomInUseConfigGet200Response{}

// ApiAudioRoomInUseConfigGet200Response struct for ApiAudioRoomInUseConfigGet200Response
type ApiAudioRoomInUseConfigGet200Response struct {
	// Time in s of activity (i.e. activityLevel > threshold) required for room in use to indicate 'true'
	TriggerTime *int32 `json:"triggerTime,omitempty"`
	// Time in s of inactivity required (i.e. activity_level < threshold) for room in use to indicate 'false'
	ReleaseTime *int32 `json:"releaseTime,omitempty"`
	// Activity detection threshold in dB for room in use to turn true
	Threshold *int32 `json:"threshold,omitempty"`
}

// NewApiAudioRoomInUseConfigGet200Response instantiates a new ApiAudioRoomInUseConfigGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAudioRoomInUseConfigGet200Response() *ApiAudioRoomInUseConfigGet200Response {
	this := ApiAudioRoomInUseConfigGet200Response{}
	var triggerTime int32 = 15
	this.TriggerTime = &triggerTime
	var releaseTime int32 = 300
	this.ReleaseTime = &releaseTime
	var threshold int32 = 10
	this.Threshold = &threshold
	return &this
}

// NewApiAudioRoomInUseConfigGet200ResponseWithDefaults instantiates a new ApiAudioRoomInUseConfigGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAudioRoomInUseConfigGet200ResponseWithDefaults() *ApiAudioRoomInUseConfigGet200Response {
	this := ApiAudioRoomInUseConfigGet200Response{}
	var triggerTime int32 = 15
	this.TriggerTime = &triggerTime
	var releaseTime int32 = 300
	this.ReleaseTime = &releaseTime
	var threshold int32 = 10
	this.Threshold = &threshold
	return &this
}

// GetTriggerTime returns the TriggerTime field value if set, zero value otherwise.
func (o *ApiAudioRoomInUseConfigGet200Response) GetTriggerTime() int32 {
	if o == nil || IsNil(o.TriggerTime) {
		var ret int32
		return ret
	}
	return *o.TriggerTime
}

// GetTriggerTimeOk returns a tuple with the TriggerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioRoomInUseConfigGet200Response) GetTriggerTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.TriggerTime) {
		return nil, false
	}
	return o.TriggerTime, true
}

// HasTriggerTime returns a boolean if a field has been set.
func (o *ApiAudioRoomInUseConfigGet200Response) HasTriggerTime() bool {
	if o != nil && !IsNil(o.TriggerTime) {
		return true
	}

	return false
}

// SetTriggerTime gets a reference to the given int32 and assigns it to the TriggerTime field.
func (o *ApiAudioRoomInUseConfigGet200Response) SetTriggerTime(v int32) {
	o.TriggerTime = &v
}

// GetReleaseTime returns the ReleaseTime field value if set, zero value otherwise.
func (o *ApiAudioRoomInUseConfigGet200Response) GetReleaseTime() int32 {
	if o == nil || IsNil(o.ReleaseTime) {
		var ret int32
		return ret
	}
	return *o.ReleaseTime
}

// GetReleaseTimeOk returns a tuple with the ReleaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioRoomInUseConfigGet200Response) GetReleaseTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ReleaseTime) {
		return nil, false
	}
	return o.ReleaseTime, true
}

// HasReleaseTime returns a boolean if a field has been set.
func (o *ApiAudioRoomInUseConfigGet200Response) HasReleaseTime() bool {
	if o != nil && !IsNil(o.ReleaseTime) {
		return true
	}

	return false
}

// SetReleaseTime gets a reference to the given int32 and assigns it to the ReleaseTime field.
func (o *ApiAudioRoomInUseConfigGet200Response) SetReleaseTime(v int32) {
	o.ReleaseTime = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *ApiAudioRoomInUseConfigGet200Response) GetThreshold() int32 {
	if o == nil || IsNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAudioRoomInUseConfigGet200Response) GetThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *ApiAudioRoomInUseConfigGet200Response) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *ApiAudioRoomInUseConfigGet200Response) SetThreshold(v int32) {
	o.Threshold = &v
}

func (o ApiAudioRoomInUseConfigGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAudioRoomInUseConfigGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TriggerTime) {
		toSerialize["triggerTime"] = o.TriggerTime
	}
	if !IsNil(o.ReleaseTime) {
		toSerialize["releaseTime"] = o.ReleaseTime
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	return toSerialize, nil
}

type NullableApiAudioRoomInUseConfigGet200Response struct {
	value *ApiAudioRoomInUseConfigGet200Response
	isSet bool
}

func (v NullableApiAudioRoomInUseConfigGet200Response) Get() *ApiAudioRoomInUseConfigGet200Response {
	return v.value
}

func (v *NullableApiAudioRoomInUseConfigGet200Response) Set(val *ApiAudioRoomInUseConfigGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAudioRoomInUseConfigGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAudioRoomInUseConfigGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAudioRoomInUseConfigGet200Response(val *ApiAudioRoomInUseConfigGet200Response) *NullableApiAudioRoomInUseConfigGet200Response {
	return &NullableApiAudioRoomInUseConfigGet200Response{value: val, isSet: true}
}

func (v NullableApiAudioRoomInUseConfigGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAudioRoomInUseConfigGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


