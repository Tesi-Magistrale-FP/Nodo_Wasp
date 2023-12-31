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

// checks if the PublicChainMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicChainMetadata{}

// PublicChainMetadata struct for PublicChainMetadata
type PublicChainMetadata struct {
	// The description of the chain.
	Description string `json:"description"`
	// The EVM json rpc url
	EvmJsonRpcURL string `json:"evmJsonRpcURL"`
	// The EVM websocket url)
	EvmWebSocketURL string `json:"evmWebSocketURL"`
	// The name of the chain
	Name string `json:"name"`
	// The official website of the chain.
	Website string `json:"website"`
}

// NewPublicChainMetadata instantiates a new PublicChainMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicChainMetadata(description string, evmJsonRpcURL string, evmWebSocketURL string, name string, website string) *PublicChainMetadata {
	this := PublicChainMetadata{}
	this.Description = description
	this.EvmJsonRpcURL = evmJsonRpcURL
	this.EvmWebSocketURL = evmWebSocketURL
	this.Name = name
	this.Website = website
	return &this
}

// NewPublicChainMetadataWithDefaults instantiates a new PublicChainMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicChainMetadataWithDefaults() *PublicChainMetadata {
	this := PublicChainMetadata{}
	return &this
}

// GetDescription returns the Description field value
func (o *PublicChainMetadata) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PublicChainMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PublicChainMetadata) SetDescription(v string) {
	o.Description = v
}

// GetEvmJsonRpcURL returns the EvmJsonRpcURL field value
func (o *PublicChainMetadata) GetEvmJsonRpcURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvmJsonRpcURL
}

// GetEvmJsonRpcURLOk returns a tuple with the EvmJsonRpcURL field value
// and a boolean to check if the value has been set.
func (o *PublicChainMetadata) GetEvmJsonRpcURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvmJsonRpcURL, true
}

// SetEvmJsonRpcURL sets field value
func (o *PublicChainMetadata) SetEvmJsonRpcURL(v string) {
	o.EvmJsonRpcURL = v
}

// GetEvmWebSocketURL returns the EvmWebSocketURL field value
func (o *PublicChainMetadata) GetEvmWebSocketURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvmWebSocketURL
}

// GetEvmWebSocketURLOk returns a tuple with the EvmWebSocketURL field value
// and a boolean to check if the value has been set.
func (o *PublicChainMetadata) GetEvmWebSocketURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvmWebSocketURL, true
}

// SetEvmWebSocketURL sets field value
func (o *PublicChainMetadata) SetEvmWebSocketURL(v string) {
	o.EvmWebSocketURL = v
}

// GetName returns the Name field value
func (o *PublicChainMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicChainMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicChainMetadata) SetName(v string) {
	o.Name = v
}

// GetWebsite returns the Website field value
func (o *PublicChainMetadata) GetWebsite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Website
}

// GetWebsiteOk returns a tuple with the Website field value
// and a boolean to check if the value has been set.
func (o *PublicChainMetadata) GetWebsiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Website, true
}

// SetWebsite sets field value
func (o *PublicChainMetadata) SetWebsite(v string) {
	o.Website = v
}

func (o PublicChainMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicChainMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["evmJsonRpcURL"] = o.EvmJsonRpcURL
	toSerialize["evmWebSocketURL"] = o.EvmWebSocketURL
	toSerialize["name"] = o.Name
	toSerialize["website"] = o.Website
	return toSerialize, nil
}

type NullablePublicChainMetadata struct {
	value *PublicChainMetadata
	isSet bool
}

func (v NullablePublicChainMetadata) Get() *PublicChainMetadata {
	return v.value
}

func (v *NullablePublicChainMetadata) Set(val *PublicChainMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicChainMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicChainMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicChainMetadata(val *PublicChainMetadata) *NullablePublicChainMetadata {
	return &NullablePublicChainMetadata{value: val, isSet: true}
}

func (v NullablePublicChainMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicChainMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


