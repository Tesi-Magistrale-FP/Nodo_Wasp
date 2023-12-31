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

// checks if the PeeringTrustRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeeringTrustRequest{}

// PeeringTrustRequest struct for PeeringTrustRequest
type PeeringTrustRequest struct {
	Name string `json:"name"`
	// The peering URL of the peer
	PeeringURL string `json:"peeringURL"`
	// The peers public key encoded in Hex
	PublicKey string `json:"publicKey"`
}

// NewPeeringTrustRequest instantiates a new PeeringTrustRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeeringTrustRequest(name string, peeringURL string, publicKey string) *PeeringTrustRequest {
	this := PeeringTrustRequest{}
	this.Name = name
	this.PeeringURL = peeringURL
	this.PublicKey = publicKey
	return &this
}

// NewPeeringTrustRequestWithDefaults instantiates a new PeeringTrustRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeeringTrustRequestWithDefaults() *PeeringTrustRequest {
	this := PeeringTrustRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PeeringTrustRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PeeringTrustRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PeeringTrustRequest) SetName(v string) {
	o.Name = v
}

// GetPeeringURL returns the PeeringURL field value
func (o *PeeringTrustRequest) GetPeeringURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeeringURL
}

// GetPeeringURLOk returns a tuple with the PeeringURL field value
// and a boolean to check if the value has been set.
func (o *PeeringTrustRequest) GetPeeringURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeeringURL, true
}

// SetPeeringURL sets field value
func (o *PeeringTrustRequest) SetPeeringURL(v string) {
	o.PeeringURL = v
}

// GetPublicKey returns the PublicKey field value
func (o *PeeringTrustRequest) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *PeeringTrustRequest) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *PeeringTrustRequest) SetPublicKey(v string) {
	o.PublicKey = v
}

func (o PeeringTrustRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PeeringTrustRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["peeringURL"] = o.PeeringURL
	toSerialize["publicKey"] = o.PublicKey
	return toSerialize, nil
}

type NullablePeeringTrustRequest struct {
	value *PeeringTrustRequest
	isSet bool
}

func (v NullablePeeringTrustRequest) Get() *PeeringTrustRequest {
	return v.value
}

func (v *NullablePeeringTrustRequest) Set(val *PeeringTrustRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePeeringTrustRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePeeringTrustRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeeringTrustRequest(val *PeeringTrustRequest) *NullablePeeringTrustRequest {
	return &NullablePeeringTrustRequest{value: val, isSet: true}
}

func (v NullablePeeringTrustRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeeringTrustRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


