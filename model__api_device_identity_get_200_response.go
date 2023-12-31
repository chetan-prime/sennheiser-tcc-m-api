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

// checks if the ApiDeviceIdentityGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDeviceIdentityGet200Response{}

// ApiDeviceIdentityGet200Response struct for ApiDeviceIdentityGet200Response
type ApiDeviceIdentityGet200Response struct {
	Product *string `json:"product,omitempty"`
	HardwareRevision *string `json:"hardwareRevision,omitempty"`
	Serial *string `json:"serial,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
}

// NewApiDeviceIdentityGet200Response instantiates a new ApiDeviceIdentityGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDeviceIdentityGet200Response() *ApiDeviceIdentityGet200Response {
	this := ApiDeviceIdentityGet200Response{}
	return &this
}

// NewApiDeviceIdentityGet200ResponseWithDefaults instantiates a new ApiDeviceIdentityGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDeviceIdentityGet200ResponseWithDefaults() *ApiDeviceIdentityGet200Response {
	this := ApiDeviceIdentityGet200Response{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *ApiDeviceIdentityGet200Response) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDeviceIdentityGet200Response) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *ApiDeviceIdentityGet200Response) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *ApiDeviceIdentityGet200Response) SetProduct(v string) {
	o.Product = &v
}

// GetHardwareRevision returns the HardwareRevision field value if set, zero value otherwise.
func (o *ApiDeviceIdentityGet200Response) GetHardwareRevision() string {
	if o == nil || IsNil(o.HardwareRevision) {
		var ret string
		return ret
	}
	return *o.HardwareRevision
}

// GetHardwareRevisionOk returns a tuple with the HardwareRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDeviceIdentityGet200Response) GetHardwareRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareRevision) {
		return nil, false
	}
	return o.HardwareRevision, true
}

// HasHardwareRevision returns a boolean if a field has been set.
func (o *ApiDeviceIdentityGet200Response) HasHardwareRevision() bool {
	if o != nil && !IsNil(o.HardwareRevision) {
		return true
	}

	return false
}

// SetHardwareRevision gets a reference to the given string and assigns it to the HardwareRevision field.
func (o *ApiDeviceIdentityGet200Response) SetHardwareRevision(v string) {
	o.HardwareRevision = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *ApiDeviceIdentityGet200Response) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDeviceIdentityGet200Response) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *ApiDeviceIdentityGet200Response) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *ApiDeviceIdentityGet200Response) SetSerial(v string) {
	o.Serial = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *ApiDeviceIdentityGet200Response) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDeviceIdentityGet200Response) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *ApiDeviceIdentityGet200Response) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *ApiDeviceIdentityGet200Response) SetVendor(v string) {
	o.Vendor = &v
}

func (o ApiDeviceIdentityGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDeviceIdentityGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.HardwareRevision) {
		toSerialize["hardwareRevision"] = o.HardwareRevision
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	return toSerialize, nil
}

type NullableApiDeviceIdentityGet200Response struct {
	value *ApiDeviceIdentityGet200Response
	isSet bool
}

func (v NullableApiDeviceIdentityGet200Response) Get() *ApiDeviceIdentityGet200Response {
	return v.value
}

func (v *NullableApiDeviceIdentityGet200Response) Set(val *ApiDeviceIdentityGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDeviceIdentityGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDeviceIdentityGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDeviceIdentityGet200Response(val *ApiDeviceIdentityGet200Response) *NullableApiDeviceIdentityGet200Response {
	return &NullableApiDeviceIdentityGet200Response{value: val, isSet: true}
}

func (v NullableApiDeviceIdentityGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDeviceIdentityGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


