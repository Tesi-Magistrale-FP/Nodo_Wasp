package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/contracts/wasm/mycontract3/go/mycontract3"
	"github.com/iotaledger/wasp/contracts/wasm/mycontract3/go/mycontract3impl"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, mycontract3.ScName, mycontract3impl.OnDispatch)
	require.NoError(t, ctx.ContractExists(mycontract3.ScName))
}
