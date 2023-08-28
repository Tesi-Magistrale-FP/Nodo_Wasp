package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/contracts/wasm/mycontract2/go/mycontract2"
	"github.com/iotaledger/wasp/contracts/wasm/mycontract2/go/mycontract2impl"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
)

func setupTest(t *testing.T, init ...*wasmlib.ScInitFunc) *wasmsolo.SoloContext 
{
	ctx := wasmsolo.NewSoloContext(t, mycontract2.ScName, mycontract2impl.OnDispatch, init...)
	require.NoError(t, ctx.Err)
	return ctx
}

func TestDeploy(t *testing.T) 
{
	ctx := setupTest(t)
	require.NoError(t, ctx.ContractExists(mycontract2.ScName))
}

func TestSetProva(t *testing.T) 
{
	init := mycontract2.ScFuncs.Init(nil)
	ctx := setupTest(t, init.Func)

	cambio := mycontract2.ScFuncs.SetProva(ctx)
	cambio.Params.Prova().SetValue("Ciao")
	cambio.Func.Post()
	require.NoError(t, ctx.Err)

	checkStateProva(t, ctx, "Ciao")
}

func checkStateProva(t *testing.T, ctx *wasmsolo.SoloContext, expected interface{}) {
	getProva := mycontract2.ScFuncs.GetProva(ctx)
	getProva.Func.Call()
	require.NoError(t, ctx.Err)
	prova := getProva.Results.Prova()
	if expected == nil {
		require.False(t, prova.Exists())
		return
	}
	require.True(t, prova.Exists())
	require.EqualValues(t, expected, prova.Value())
}