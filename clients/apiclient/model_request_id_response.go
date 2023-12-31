/*
Wasp API

REST API for the Wasp node

API version: 0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the RequestIDResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestIDResponse{}

// RequestIDResponse struct for RequestIDResponse
type RequestIDResponse struct {
	// The request ID of the given transaction ID. (Hex)
	RequestId string `json:"requestId"`
}

// NewRequestIDResponse instantiates a new RequestIDResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestIDResponse(requestId string) *RequestIDResponse {
	this := RequestIDResponse{}
	this.RequestId = requestId
	return &this
}

// NewRequestIDResponseWithDefaults instantiates a new RequestIDResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestIDResponseWithDefaults() *RequestIDResponse {
	this := RequestIDResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *RequestIDResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *RequestIDResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *RequestIDResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o RequestIDResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestIDResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestId"] = o.RequestId
	return toSerialize, nil
}

type NullableRequestIDResponse struct {
	value *RequestIDResponse
	isSet bool
}

func (v NullableRequestIDResponse) Get() *RequestIDResponse {
	return v.value
}

func (v *NullableRequestIDResponse) Set(val *RequestIDResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestIDResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestIDResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestIDResponse(val *RequestIDResponse) *NullableRequestIDResponse {
	return &NullableRequestIDResponse{value: val, isSet: true}
}

func (v NullableRequestIDResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestIDResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


