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

// checks if the ApiSscVersionGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSscVersionGet200Response{}

// ApiSscVersionGet200Response struct for ApiSscVersionGet200Response
type ApiSscVersionGet200Response struct {
	// The version of SSC protocol.
	Protocol *string `json:"protocol,omitempty"`
	// This is the schema version of the API of the TCC M device. Semantic versioning must be used. So additional paths or properties will increase the minor number while breaking changes increase the integer part before the decimal point.
	Schema *string `json:"schema,omitempty"`
}

// NewApiSscVersionGet200Response instantiates a new ApiSscVersionGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSscVersionGet200Response() *ApiSscVersionGet200Response {
	this := ApiSscVersionGet200Response{}
	return &this
}

// NewApiSscVersionGet200ResponseWithDefaults instantiates a new ApiSscVersionGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSscVersionGet200ResponseWithDefaults() *ApiSscVersionGet200Response {
	this := ApiSscVersionGet200Response{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApiSscVersionGet200Response) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSscVersionGet200Response) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApiSscVersionGet200Response) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ApiSscVersionGet200Response) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ApiSscVersionGet200Response) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSscVersionGet200Response) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ApiSscVersionGet200Response) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *ApiSscVersionGet200Response) SetSchema(v string) {
	o.Schema = &v
}

func (o ApiSscVersionGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSscVersionGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	return toSerialize, nil
}

type NullableApiSscVersionGet200Response struct {
	value *ApiSscVersionGet200Response
	isSet bool
}

func (v NullableApiSscVersionGet200Response) Get() *ApiSscVersionGet200Response {
	return v.value
}

func (v *NullableApiSscVersionGet200Response) Set(val *ApiSscVersionGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSscVersionGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSscVersionGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSscVersionGet200Response(val *ApiSscVersionGet200Response) *NullableApiSscVersionGet200Response {
	return &NullableApiSscVersionGet200Response{value: val, isSet: true}
}

func (v NullableApiSscVersionGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSscVersionGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


