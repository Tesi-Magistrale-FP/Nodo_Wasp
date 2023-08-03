// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package emulator

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/buffered"
	"github.com/iotaledger/wasp/packages/kv/codec"
)

const (
	keyAccountNonce    = "n"
	keyAccountCode     = "c"
	keyAccountState    = "s"
	keyAccountSuicided = "S"
)

func accountKey(prefix kv.Key, addr common.Address) kv.Key {
	return prefix + kv.Key(addr.Bytes())
}

func accountNonceKey(addr common.Address) kv.Key {
	return accountKey(keyAccountNonce, addr)
}

func accountCodeKey(addr common.Address) kv.Key {
	return accountKey(keyAccountCode, addr)
}

func accountStateKey(addr common.Address, hash common.Hash) kv.Key {
	return accountKey(keyAccountState, addr) + kv.Key(hash[:])
}

func accountSuicidedKey(addr common.Address) kv.Key {
	return accountKey(keyAccountSuicided, addr)
}

type L2Balance interface {
	Get(addr common.Address) *big.Int
	Add(addr common.Address, amount *big.Int)
	Sub(addr common.Address, amount *big.Int)
}

// StateDB implements vm.StateDB with a kv.KVStore as backend.
// The Ethereum account balance is tied to the L1 balance.
type StateDB struct {
	kv        kv.KVStore
	logs      []*types.Log
	refund    uint64
	l2Balance L2Balance
}

var _ vm.StateDB = &StateDB{}

func NewStateDB(store kv.KVStore, l2Balance L2Balance) *StateDB {
	return &StateDB{
		kv:        StateDBSubrealm(store),
		l2Balance: l2Balance,
	}
}

func (s *StateDB) CreateAccount(addr common.Address) {
	s.SetNonce(addr, 0)
}

func (s *StateDB) SubBalance(addr common.Address, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	if amount.Sign() == -1 {
		panic("unexpected negative amount")
	}
	s.l2Balance.Sub(addr, amount)
}

func (s *StateDB) AddBalance(addr common.Address, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	if amount.Sign() == -1 {
		panic("unexpected negative amount")
	}
	s.l2Balance.Add(addr, amount)
}

func (s *StateDB) GetBalance(addr common.Address) *big.Int {
	return s.l2Balance.Get(addr)
}

func GetNonce(s kv.KVStoreReader, addr common.Address) uint64 {
	return codec.MustDecodeUint64(s.Get(accountNonceKey(addr)), 0)
}

func (s *StateDB) GetNonce(addr common.Address) uint64 {
	return GetNonce(s.kv, addr)
}

func (s *StateDB) IncNonce(addr common.Address) {
	s.SetNonce(addr, s.GetNonce(addr)+1)
}

func (s *StateDB) SetNonce(addr common.Address, n uint64) {
	s.kv.Set(accountNonceKey(addr), codec.EncodeUint64(n))
}

func (s *StateDB) GetCodeHash(addr common.Address) common.Hash {
	// TODO cache the code hash?
	return crypto.Keccak256Hash(s.GetCode(addr))
}

func (s *StateDB) GetCode(addr common.Address) []byte {
	return s.kv.Get(accountCodeKey(addr))
}

func (s *StateDB) SetCode(addr common.Address, code []byte) {
	if code == nil {
		s.kv.Del(accountCodeKey(addr))
	} else {
		s.kv.Set(accountCodeKey(addr), code)
	}
}

func (s *StateDB) GetCodeSize(addr common.Address) int {
	// TODO cache the code size?
	return len(s.GetCode(addr))
}

func (s *StateDB) AddRefund(n uint64) {
	s.refund += n
}

func (s *StateDB) SubRefund(n uint64) {
	if n > s.refund {
		panic(fmt.Sprintf("Refund counter below zero (gas: %d > refund: %d)", n, s.refund))
	}
	s.refund -= n
}

func (s *StateDB) GetRefund() uint64 {
	return s.refund
}

func (s *StateDB) GetCommittedState(addr common.Address, key common.Hash) common.Hash {
	return s.GetState(addr, key)
}

func (s *StateDB) GetState(addr common.Address, key common.Hash) common.Hash {
	return common.BytesToHash(s.kv.Get(accountStateKey(addr, key)))
}

func (s *StateDB) SetState(addr common.Address, key, value common.Hash) {
	s.kv.Set(accountStateKey(addr, key), value.Bytes())
}

func (s *StateDB) Suicide(addr common.Address) bool {
	if !s.Exist(addr) {
		return false
	}

	s.kv.Del(accountNonceKey(addr))
	s.kv.Del(accountCodeKey(addr))

	keys := make([]kv.Key, 0)
	s.kv.IterateKeys(accountKey(keyAccountState, addr), func(key kv.Key) bool {
		keys = append(keys, key)
		return true
	})
	for _, k := range keys {
		s.kv.Del(k)
	}

	// for some reason the EVM engine calls AddBalance to the beneficiary address,
	// but not SubBalance for the suicided address.
	s.l2Balance.Sub(addr, s.l2Balance.Get(addr))

	s.kv.Set(accountSuicidedKey(addr), []byte{1})

	return true
}

func (s *StateDB) HasSuicided(addr common.Address) bool {
	return s.kv.Has(accountSuicidedKey(addr))
}

// Exist reports whether the given account exists in state.
// Notably this should also return true for suicided accounts.
func (s *StateDB) Exist(addr common.Address) bool {
	return s.kv.Has(accountNonceKey(addr))
}

// Empty returns whether the given account is empty. Empty
// is defined according to EIP161 (balance = nonce = code = 0).
func (s *StateDB) Empty(addr common.Address) bool {
	return s.GetNonce(addr) == 0 && s.GetBalance(addr).Sign() == 0 && s.GetCodeSize(addr) == 0
}

func (s *StateDB) PrepareAccessList(sender common.Address, dest *common.Address, precompiles []common.Address, txAccesses types.AccessList) {
	_ = sender
	_ = dest
	_ = precompiles
	_ = txAccesses
}

func (s *StateDB) AddressInAccessList(addr common.Address) bool {
	_ = addr
	return true
}

func (s *StateDB) SlotInAccessList(addr common.Address, slot common.Hash) (addressOk, slotOk bool) {
	_ = addr
	_ = slot
	return true, true
}

// AddAddressToAccessList adds the given address to the access list. This operation is safe to perform
// even if the feature/fork is not active yet
func (s *StateDB) AddAddressToAccessList(addr common.Address) {
	_ = addr
}

// AddSlotToAccessList adds the given (address,slot) to the access list. This operation is safe to perform
// even if the feature/fork is not active yet
func (s *StateDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	_ = addr
	_ = slot
}

func (s *StateDB) RevertToSnapshot(int) {}
func (s *StateDB) Snapshot() int        { return 0 }

func (s *StateDB) AddLog(log *types.Log) {
	log.Index = uint(len(s.logs))
	s.logs = append(s.logs, log)
}

func (s *StateDB) GetLogs(_ common.Hash) []*types.Log {
	return s.logs
}

func (s *StateDB) AddPreimage(common.Hash, []byte) { panic("not implemented") }

func (s *StateDB) ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) error {
	panic("not implemented")
}

func (s *StateDB) Buffered() *BufferedStateDB {
	return NewBufferedStateDB(s)
}

// GetTransientState implements vm.StateDB
func (*StateDB) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	panic("unimplemented")
}

// Prepare implements vm.StateDB
func (s *StateDB) Prepare(rules params.Rules, sender common.Address, coinbase common.Address, dest *common.Address, precompiles []common.Address, txAccesses types.AccessList) {
	// do nothing
}

// SetTransientState implements vm.StateDB
func (*StateDB) SetTransientState(addr common.Address, key common.Hash, value common.Hash) {
	panic("unimplemented")
}

// BufferedStateDB is a wrapper for StateDB that writes all mutations into an in-memory buffer,
// leaving the original state unmodified until the mutations are applied manually with Commit().
type BufferedStateDB struct {
	buf       *buffered.BufferedKVStore
	base      kv.KVStore
	l2Balance L2Balance
}

func NewBufferedStateDB(base *StateDB) *BufferedStateDB {
	return &BufferedStateDB{
		buf:       buffered.NewBufferedKVStore(base.kv),
		base:      base.kv,
		l2Balance: base.l2Balance,
	}
}

func (b *BufferedStateDB) StateDB() *StateDB {
	return &StateDB{
		kv:        b.buf,
		l2Balance: b.l2Balance,
	}
}

func (b *BufferedStateDB) Commit() {
	b.buf.Mutations().ApplyTo(b.base)
}
