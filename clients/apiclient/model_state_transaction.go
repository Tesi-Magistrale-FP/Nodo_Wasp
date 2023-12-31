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

// checks if the StateTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateTransaction{}

// StateTransaction struct for StateTransaction
type StateTransaction struct {
	// The state index
	StateIndex uint32 `json:"stateIndex"`
	// The transaction ID
	TxId string `json:"txId"`
}

// NewStateTransaction instantiates a new StateTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateTransaction(stateIndex uint32, txId string) *StateTransaction {
	this := StateTransaction{}
	this.StateIndex = stateIndex
	this.TxId = txId
	return &this
}

// NewStateTransactionWithDefaults instantiates a new StateTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateTransactionWithDefaults() *StateTransaction {
	this := StateTransaction{}
	return &this
}

// GetStateIndex returns the StateIndex field value
func (o *StateTransaction) GetStateIndex() uint32 {
	if o == nil {
		var ret uint32
		return ret
	}

	return o.StateIndex
}

// GetStateIndexOk returns a tuple with the StateIndex field value
// and a boolean to check if the value has been set.
func (o *StateTransaction) GetStateIndexOk() (*uint32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateIndex, true
}

// SetStateIndex sets field value
func (o *StateTransaction) SetStateIndex(v uint32) {
	o.StateIndex = v
}

// GetTxId returns the TxId field value
func (o *StateTransaction) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *StateTransaction) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *StateTransaction) SetTxId(v string) {
	o.TxId = v
}

func (o StateTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stateIndex"] = o.StateIndex
	toSerialize["txId"] = o.TxId
	return toSerialize, nil
}

type NullableStateTransaction struct {
	value *StateTransaction
	isSet bool
}

func (v NullableStateTransaction) Get() *StateTransaction {
	return v.value
}

func (v *NullableStateTransaction) Set(val *StateTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableStateTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableStateTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateTransaction(val *StateTransaction) *NullableStateTransaction {
	return &NullableStateTransaction{value: val, isSet: true}
}

func (v NullableStateTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


