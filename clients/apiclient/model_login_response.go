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

// checks if the LoginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginResponse{}

// LoginResponse struct for LoginResponse
type LoginResponse struct {
	Error string `json:"error"`
	Jwt string `json:"jwt"`
}

// NewLoginResponse instantiates a new LoginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginResponse(error_ string, jwt string) *LoginResponse {
	this := LoginResponse{}
	this.Error = error_
	this.Jwt = jwt
	return &this
}

// NewLoginResponseWithDefaults instantiates a new LoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginResponseWithDefaults() *LoginResponse {
	this := LoginResponse{}
	return &this
}

// GetError returns the Error field value
func (o *LoginResponse) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *LoginResponse) SetError(v string) {
	o.Error = v
}

// GetJwt returns the Jwt field value
func (o *LoginResponse) GetJwt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Jwt
}

// GetJwtOk returns a tuple with the Jwt field value
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetJwtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jwt, true
}

// SetJwt sets field value
func (o *LoginResponse) SetJwt(v string) {
	o.Jwt = v
}

func (o LoginResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["jwt"] = o.Jwt
	return toSerialize, nil
}

type NullableLoginResponse struct {
	value *LoginResponse
	isSet bool
}

func (v NullableLoginResponse) Get() *LoginResponse {
	return v.value
}

func (v *NullableLoginResponse) Set(val *LoginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginResponse(val *LoginResponse) *NullableLoginResponse {
	return &NullableLoginResponse{value: val, isSet: true}
}

func (v NullableLoginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


