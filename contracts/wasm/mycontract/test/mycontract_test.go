package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/contracts/wasm/mycontract/go/mycontract"
	"github.com/iotaledger/wasp/contracts/wasm/mycontract/go/mycontractimpl"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, mycontract.ScName, mycontractimpl.OnDispatch)
	require.NoError(t, ctx.ContractExists(mycontract.ScName))
}
