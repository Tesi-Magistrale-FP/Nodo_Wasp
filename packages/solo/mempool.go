// Here we implement a trivial mempool, just for solo tests.
// We don't use the mempool from the chain/consensus because
// it does a lot of functions not needed in this context.
// The interface of this mempool has a little in common with
// the real mempool implementation.

package solo

import (
	"fmt"
	"time"

	"github.com/iotaledger/wasp/packages/isc"
)

type Mempool interface {
	ReceiveRequests(reqs ...isc.Request)
	RequestBatchProposal() []isc.Request
	RemoveRequest(reqs isc.RequestID)
	Info() MempoolInfo
}

type MempoolInfo struct {
	TotalPool      int
	InPoolCounter  int
	OutPoolCounter int
}

type mempoolImpl struct {
	requests    map[isc.RequestID]isc.Request
	info        MempoolInfo
	currentTime func() time.Time
}

func newMempool(currentTime func() time.Time) Mempool {
	return &mempoolImpl{
		requests:    map[isc.RequestID]isc.Request{},
		info:        MempoolInfo{},
		currentTime: currentTime,
	}
}

func (mi *mempoolImpl) ReceiveRequests(reqs ...isc.Request) {
	for _, req := range reqs {
		if _, ok := mi.requests[req.ID()]; !ok {
			mi.info.TotalPool++
			mi.info.InPoolCounter++
		}
		mi.requests[req.ID()] = req
	}
}

func (mi *mempoolImpl) RequestBatchProposal() []isc.Request {
	now := mi.currentTime()
	batch := []isc.Request{}
	for rid, request := range mi.requests {
		switch request := request.(type) {
		case isc.OnLedgerRequest:
			reqUnlockCondSet := request.Output().UnlockConditionSet()
			timeLock := reqUnlockCondSet.Timelock()
			expiration := reqUnlockCondSet.Expiration()
			if expiration != nil && timeLock.UnixTime >= expiration.UnixTime {
				// can never be processed, just reject
				delete(mi.requests, rid)
				continue
			}
			if timeLock == nil || timeLock.UnixTime <= uint32(now.Unix()) {
				batch = append(batch, request)
				continue
			}
		case isc.OffLedgerRequest:
			batch = append(batch, request)
		default:
			panic(fmt.Errorf("unexpected request type %T: %+v", request, request))
		}
	}
	return batch
}

func (mi *mempoolImpl) RemoveRequest(rID isc.RequestID) {
	if _, ok := mi.requests[rID]; ok {
		mi.info.OutPoolCounter++
		mi.info.TotalPool--
	}
	delete(mi.requests, rID)
}

func (mi *mempoolImpl) Info() MempoolInfo {
	return mi.info
}
