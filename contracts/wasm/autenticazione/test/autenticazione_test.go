package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/contracts/wasm/autenticazione/go/autenticazione"
	"github.com/iotaledger/wasp/contracts/wasm/autenticazione/go/autenticazioneimpl"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, autenticazione.ScName, autenticazioneimpl.OnDispatch)
	require.NoError(t, ctx.ContractExists(autenticazione.ScName))
}
