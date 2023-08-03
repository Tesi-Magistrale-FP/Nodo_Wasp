// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package coreblocklog

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

type ControlAddressesCall struct {
	Func    *wasmlib.ScView
	Results ImmutableControlAddressesResults
}

type GetBlockInfoCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetBlockInfoParams
	Results ImmutableGetBlockInfoResults
}

type GetEventsForBlockCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetEventsForBlockParams
	Results ImmutableGetEventsForBlockResults
}

type GetEventsForContractCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetEventsForContractParams
	Results ImmutableGetEventsForContractResults
}

type GetEventsForRequestCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetEventsForRequestParams
	Results ImmutableGetEventsForRequestResults
}

type GetRequestIDsForBlockCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetRequestIDsForBlockParams
	Results ImmutableGetRequestIDsForBlockResults
}

type GetRequestReceiptCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetRequestReceiptParams
	Results ImmutableGetRequestReceiptResults
}

type GetRequestReceiptsForBlockCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetRequestReceiptsForBlockParams
	Results ImmutableGetRequestReceiptsForBlockResults
}

type IsRequestProcessedCall struct {
	Func    *wasmlib.ScView
	Params  MutableIsRequestProcessedParams
	Results ImmutableIsRequestProcessedResults
}

type Funcs struct{}

var ScFuncs Funcs

// Returns the current state controller and governing addresses and at what block index they were set.
func (sc Funcs) ControlAddresses(ctx wasmlib.ScViewClientContext) *ControlAddressesCall {
	f := &ControlAddressesCall{Func: wasmlib.NewScView(ctx, HScName, HViewControlAddresses)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// Returns information about the given block.
func (sc Funcs) GetBlockInfo(ctx wasmlib.ScViewClientContext) *GetBlockInfoCall {
	f := &GetBlockInfoCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetBlockInfo)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// Returns the list of events triggered during the execution of the given block.
func (sc Funcs) GetEventsForBlock(ctx wasmlib.ScViewClientContext) *GetEventsForBlockCall {
	f := &GetEventsForBlockCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetEventsForBlock)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// Returns the list of events triggered by the given contract
// during the execution of the given block range.
func (sc Funcs) GetEventsForContract(ctx wasmlib.ScViewClientContext) *GetEventsForContractCall {
	f := &GetEventsForContractCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetEventsForContract)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// Returns the list of events triggered during the execution of the given request.
func (sc Funcs) GetEventsForRequest(ctx wasmlib.ScViewClientContext) *GetEventsForRequestCall {
	f := &GetEventsForRequestCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetEventsForRequest)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// Returns a list with all request IDs in the given block.
func (sc Funcs) GetRequestIDsForBlock(ctx wasmlib.ScViewClientContext) *GetRequestIDsForBlockCall {
	f := &GetRequestIDsForBlockCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetRequestIDsForBlock)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// Returns the receipt for the request with the given ID.
func (sc Funcs) GetRequestReceipt(ctx wasmlib.ScViewClientContext) *GetRequestReceiptCall {
	f := &GetRequestReceiptCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetRequestReceipt)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// Returns all request receipts in the given block.
func (sc Funcs) GetRequestReceiptsForBlock(ctx wasmlib.ScViewClientContext) *GetRequestReceiptsForBlockCall {
	f := &GetRequestReceiptsForBlockCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetRequestReceiptsForBlock)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// Returns whether the request with ID u has been processed.
func (sc Funcs) IsRequestProcessed(ctx wasmlib.ScViewClientContext) *IsRequestProcessedCall {
	f := &IsRequestProcessedCall{Func: wasmlib.NewScView(ctx, HScName, HViewIsRequestProcessed)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

var exportMap = wasmlib.ScExportMap{
	Names: []string{
		ViewControlAddresses,
		ViewGetBlockInfo,
		ViewGetEventsForBlock,
		ViewGetEventsForContract,
		ViewGetEventsForRequest,
		ViewGetRequestIDsForBlock,
		ViewGetRequestReceipt,
		ViewGetRequestReceiptsForBlock,
		ViewIsRequestProcessed,
	},
	Funcs: []wasmlib.ScFuncContextFunction{
	},
	Views: []wasmlib.ScViewContextFunction{
		wasmlib.ViewError,
		wasmlib.ViewError,
		wasmlib.ViewError,
		wasmlib.ViewError,
		wasmlib.ViewError,
		wasmlib.ViewError,
		wasmlib.ViewError,
		wasmlib.ViewError,
		wasmlib.ViewError,
	},
}

func OnDispatch(index int32) *wasmlib.ScExportMap {
	if index < 0 {
		return exportMap.Dispatch(index)
	}

	panic("Calling core contract?")
}