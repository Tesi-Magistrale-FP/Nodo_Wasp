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

// checks if the Event type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Event{}

// Event struct for Event
type Event struct {
	ContractID *int32 `json:"contractID,omitempty"`
	Payload []int32 `json:"payload,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	Topic *string `json:"topic,omitempty"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent() *Event {
	this := Event{}
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetContractID returns the ContractID field value if set, zero value otherwise.
func (o *Event) GetContractID() int32 {
	if o == nil || isNil(o.ContractID) {
		var ret int32
		return ret
	}
	return *o.ContractID
}

// GetContractIDOk returns a tuple with the ContractID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetContractIDOk() (*int32, bool) {
	if o == nil || isNil(o.ContractID) {
		return nil, false
	}
	return o.ContractID, true
}

// HasContractID returns a boolean if a field has been set.
func (o *Event) HasContractID() bool {
	if o != nil && !isNil(o.ContractID) {
		return true
	}

	return false
}

// SetContractID gets a reference to the given int32 and assigns it to the ContractID field.
func (o *Event) SetContractID(v int32) {
	o.ContractID = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *Event) GetPayload() []int32 {
	if o == nil || isNil(o.Payload) {
		var ret []int32
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetPayloadOk() ([]int32, bool) {
	if o == nil || isNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *Event) HasPayload() bool {
	if o != nil && !isNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given []int32 and assigns it to the Payload field.
func (o *Event) SetPayload(v []int32) {
	o.Payload = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Event) GetTimestamp() int64 {
	if o == nil || isNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Event) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *Event) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *Event) GetTopic() string {
	if o == nil || isNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTopicOk() (*string, bool) {
	if o == nil || isNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *Event) HasTopic() bool {
	if o != nil && !isNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *Event) SetTopic(v string) {
	o.Topic = &v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContractID) {
		toSerialize["contractID"] = o.ContractID
	}
	if !isNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	return toSerialize, nil
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


