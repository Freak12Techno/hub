// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/querier
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/types
package dvpn

import (
	"github.com/sentinel-official/hub/x/dvpn/keeper"
	"github.com/sentinel-official/hub/x/dvpn/querier"
	"github.com/sentinel-official/hub/x/dvpn/types"
)

const (
	Codespace    = types.Codespace
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	QuerierRoute = types.QuerierRoute
	StoreKey     = types.StoreKey
)

var (
	// functions aliases
	NewKeeper             = keeper.NewKeeper
	NewQuerier            = querier.NewQuerier
	RegisterCodec         = types.RegisterCodec
	ErrorMarshal          = types.ErrorMarshal
	ErrorUnmarshal        = types.ErrorUnmarshal
	ErrorUnknownMsgType   = types.ErrorUnknownMsgType
	ErrorUnknownQueryType = types.ErrorUnknownQueryType
	NewGenesisState       = types.NewGenesisState
	DefaultGenesisState   = types.DefaultGenesisState

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
)
