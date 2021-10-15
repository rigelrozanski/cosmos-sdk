package v045

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"

	v045gov "github.com/cosmos/cosmos-sdk/x/gov/migrations/v045"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func Migrate(appState types.AppMap, clientCtx client.Context) types.AppMap {
	// Migrate x/gov proposals
	if appState[v1.ModuleName] != nil {
		// unmarshal relative source genesis application state
		var oldGovState v1beta1.GenesisState
		clientCtx.Codec.MustUnmarshalJSON(appState[v1.ModuleName], &oldGovState)

		// delete deprecated x/gov genesis state
		delete(appState, types.ModuleName)

		// Migrate relative source genesis application state and marshal it into
		// the respective key.
		appState[v1.ModuleName] = clientCtx.Codec.MustMarshalJSON(v045gov.Migrate(oldGovState))
	}

	return appState
}