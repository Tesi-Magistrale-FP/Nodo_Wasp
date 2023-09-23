package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/contracts/wasm/autorizzazioni/go/autorizzazioni"
	"github.com/iotaledger/wasp/contracts/wasm/autorizzazioni/go/autorizzazioniimpl"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, autorizzazioni.ScName, autorizzazioniimpl.OnDispatch)
	require.NoError(t, ctx.ContractExists(autorizzazioni.ScName))
}
