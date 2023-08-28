package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/contracts/wasm/mycontract1/go/mycontract1"
	"github.com/iotaledger/wasp/contracts/wasm/mycontract1/go/mycontract1impl"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, mycontract1.ScName, mycontract1impl.OnDispatch)
	require.NoError(t, ctx.ContractExists(mycontract1.ScName))
}
