package system

import (
	yta "github.com/YstarLab/yta-go"
)

// NewUndelegateBW returns a `undelegatebw` action that lives on the
// `eosio.system` contract.
func NewUndelegateBW(from, receiver yta.AccountName, unstakeCPU, unstakeNet yta.Asset) *yta.Action {
	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("undelegatebw"),
		Authorization: []yta.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(UndelegateBW{
			From:       from,
			Receiver:   receiver,
			UnstakeNet: unstakeNet,
			UnstakeCPU: unstakeCPU,
		}),
	}
}

// UndelegateBW represents the `eosio.system::undelegatebw` action.
type UndelegateBW struct {
	From       yta.AccountName `json:"from"`
	Receiver   yta.AccountName `json:"receiver"`
	UnstakeNet yta.Asset       `json:"unstake_net_quantity"`
	UnstakeCPU yta.Asset       `json:"unstake_cpu_quantity"`
}
