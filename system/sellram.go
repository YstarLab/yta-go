package system

import (
	yta "github.com/YstarLab/yta-go"
)

// NewSellRAM will sell at current market price a given number of
// bytes of RAM.
func NewSellRAM(account yta.AccountName, bytes uint64) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("sellram"),
		Authorization: []yta.PermissionLevel{
			{Actor: account, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(SellRAM{
			Account: account,
			Bytes:   bytes,
		}),
	}
	return a
}

// SellRAM represents the `eosio.system::sellram` action.
type SellRAM struct {
	Account yta.AccountName `json:"account"`
	Bytes   uint64          `json:"bytes"`
}
