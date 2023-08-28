package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/contracts/wasm/mycontract4/go/mycontract4"
	"github.com/iotaledger/wasp/contracts/wasm/mycontract4/go/mycontract4impl"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, mycontract4.ScName, mycontract4impl.OnDispatch)
	require.NoError(t, ctx.ContractExists(mycontract4.ScName))
}
