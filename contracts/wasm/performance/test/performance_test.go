package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/contracts/wasm/performance/go/performance"
	"github.com/iotaledger/wasp/contracts/wasm/performance/go/performanceimpl"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, performance.ScName, performanceimpl.OnDispatch)
	require.NoError(t, ctx.ContractExists(performance.ScName))
}
