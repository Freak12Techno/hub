package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type VpnDetails struct {
	Ip         string
	Port       string
	NetSpeed   NetSpeed
	PricePerGb int64
	EncMethod  string
	Location   Location
	Version    string
	Info       Info
}

type NetSpeed struct {
	Upload   int64
	Download int64
}

type Location struct {
	Latitude  int64
	Longitude int64
	City      string
	Country   string
}

type VpnIBCPacket struct {
	VpnId     string
	Address   sdkTypes.AccAddress
	Coin      sdkTypes.Coins
	DestChain string
}

type Info struct {
	Status bool
	BlockHeight int64
}

type IBCMsgRegisterVpn struct {
	VpnId   sdkTypes.AccAddress
	Address sdkTypes.AccAddress
	Coins   sdkTypes.Coins
}
