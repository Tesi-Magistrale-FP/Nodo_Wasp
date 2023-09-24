package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/contracts/wasm/gestioneapplicazioni/go/gestioneapplicazioni"
	"github.com/iotaledger/wasp/contracts/wasm/gestioneapplicazioni/go/gestioneapplicazioniimpl"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, gestioneapplicazioni.ScName, gestioneapplicazioniimpl.OnDispatch)
	require.NoError(t, ctx.ContractExists(gestioneapplicazioni.ScName))
}
