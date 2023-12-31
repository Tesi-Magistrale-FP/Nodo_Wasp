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

// checks if the PublisherStateTransactionItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublisherStateTransactionItem{}

// PublisherStateTransactionItem struct for PublisherStateTransactionItem
type PublisherStateTransactionItem struct {
	LastMessage StateTransaction `json:"lastMessage"`
	Messages uint32 `json:"messages"`
	Timestamp time.Time `json:"timestamp"`
}

// NewPublisherStateTransactionItem instantiates a new PublisherStateTransactionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublisherStateTransactionItem(lastMessage StateTransaction, messages uint32, timestamp time.Time) *PublisherStateTransactionItem {
	this := PublisherStateTransactionItem{}
	this.LastMessage = lastMessage
	this.Messages = messages
	this.Timestamp = timestamp
	return &this
}

// NewPublisherStateTransactionItemWithDefaults instantiates a new PublisherStateTransactionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublisherStateTransactionItemWithDefaults() *PublisherStateTransactionItem {
	this := PublisherStateTransactionItem{}
	return &this
}

// GetLastMessage returns the LastMessage field value
func (o *PublisherStateTransactionItem) GetLastMessage() StateTransaction {
	if o == nil {
		var ret StateTransaction
		return ret
	}

	return o.LastMessage
}

// GetLastMessageOk returns a tuple with the LastMessage field value
// and a boolean to check if the value has been set.
func (o *PublisherStateTransactionItem) GetLastMessageOk() (*StateTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastMessage, true
}

// SetLastMessage sets field value
func (o *PublisherStateTransactionItem) SetLastMessage(v StateTransaction) {
	o.LastMessage = v
}

// GetMessages returns the Messages field value
func (o *PublisherStateTransactionItem) GetMessages() uint32 {
	if o == nil {
		var ret uint32
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *PublisherStateTransactionItem) GetMessagesOk() (*uint32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value
func (o *PublisherStateTransactionItem) SetMessages(v uint32) {
	o.Messages = v
}

// GetTimestamp returns the Timestamp field value
func (o *PublisherStateTransactionItem) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *PublisherStateTransactionItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *PublisherStateTransactionItem) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o PublisherStateTransactionItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublisherStateTransactionItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lastMessage"] = o.LastMessage
	toSerialize["messages"] = o.Messages
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullablePublisherStateTransactionItem struct {
	value *PublisherStateTransactionItem
	isSet bool
}

func (v NullablePublisherStateTransactionItem) Get() *PublisherStateTransactionItem {
	return v.value
}

func (v *NullablePublisherStateTransactionItem) Set(val *PublisherStateTransactionItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePublisherStateTransactionItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePublisherStateTransactionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublisherStateTransactionItem(val *PublisherStateTransactionItem) *NullablePublisherStateTransactionItem {
	return &NullablePublisherStateTransactionItem{value: val, isSet: true}
}

func (v NullablePublisherStateTransactionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublisherStateTransactionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


