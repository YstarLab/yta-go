package system

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewSetalimits sets the account limits. Requires signature from `eosio@active` account.
func NewSetalimits(account yta.AccountName, ramBytes, netWeight, cpuWeight int64) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("setalimit"),
		Authorization: []yta.PermissionLevel{
			{Actor: yta.AccountName("eosio"), Permission: PN("active")},
		},
		ActionData: yta.NewActionData(Setalimits{
			Account:   account,
			RAMBytes:  ramBytes,
			NetWeight: netWeight,
			CPUWeight: cpuWeight,
		}),
	}
	return a
}

// Setalimits represents the `eosio.system::setalimit` action.
type Setalimits struct {
	Account   yta.AccountName `json:"account"`
	RAMBytes  int64           `json:"ram_bytes"`
	NetWeight int64           `json:"net_weight"`
	CPUWeight int64           `json:"cpu_weight"`
}
