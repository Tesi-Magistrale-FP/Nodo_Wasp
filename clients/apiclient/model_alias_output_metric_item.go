/*
Wasp API

REST API for the Wasp node

API version: 0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
	"time"
)

// checks if the AliasOutputMetricItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AliasOutputMetricItem{}

// AliasOutputMetricItem struct for AliasOutputMetricItem
type AliasOutputMetricItem struct {
	LastMessage Output `json:"lastMessage"`
	Messages uint32 `json:"messages"`
	Timestamp time.Time `json:"timestamp"`
}

// NewAliasOutputMetricItem instantiates a new AliasOutputMetricItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAliasOutputMetricItem(lastMessage Output, messages uint32, timestamp time.Time) *AliasOutputMetricItem {
	this := AliasOutputMetricItem{}
	this.LastMessage = lastMessage
	this.Messages = messages
	this.Timestamp = timestamp
	return &this
}

// NewAliasOutputMetricItemWithDefaults instantiates a new AliasOutputMetricItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliasOutputMetricItemWithDefaults() *AliasOutputMetricItem {
	this := AliasOutputMetricItem{}
	return &this
}

// GetLastMessage returns the LastMessage field value
func (o *AliasOutputMetricItem) GetLastMessage() Output {
	if o == nil {
		var ret Output
		return ret
	}

	return o.LastMessage
}

// GetLastMessageOk returns a tuple with the LastMessage field value
// and a boolean to check if the value has been set.
func (o *AliasOutputMetricItem) GetLastMessageOk() (*Output, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastMessage, true
}

// SetLastMessage sets field value
func (o *AliasOutputMetricItem) SetLastMessage(v Output) {
	o.LastMessage = v
}

// GetMessages returns the Messages field value
func (o *AliasOutputMetricItem) GetMessages() uint32 {
	if o == nil {
		var ret uint32
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *AliasOutputMetricItem) GetMessagesOk() (*uint32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value
func (o *AliasOutputMetricItem) SetMessages(v uint32) {
	o.Messages = v
}

// GetTimestamp returns the Timestamp field value
func (o *AliasOutputMetricItem) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *AliasOutputMetricItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *AliasOutputMetricItem) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o AliasOutputMetricItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AliasOutputMetricItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lastMessage"] = o.LastMessage
	toSerialize["messages"] = o.Messages
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullableAliasOutputMetricItem struct {
	value *AliasOutputMetricItem
	isSet bool
}

func (v NullableAliasOutputMetricItem) Get() *AliasOutputMetricItem {
	return v.value
}

func (v *NullableAliasOutputMetricItem) Set(val *AliasOutputMetricItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAliasOutputMetricItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAliasOutputMetricItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAliasOutputMetricItem(val *AliasOutputMetricItem) *NullableAliasOutputMetricItem {
	return &NullableAliasOutputMetricItem{value: val, isSet: true}
}

func (v NullableAliasOutputMetricItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAliasOutputMetricItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


