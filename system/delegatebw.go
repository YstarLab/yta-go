package system

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewDelegateBW returns a `delegatebw` action that lives on the
// `eosio.system` contract.
func NewDelegateBW(from, receiver yta.AccountName, stakeCPU, stakeNet yta.Asset, transfer bool) *yta.Action {
	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("delegatebw"),
		Authorization: []yta.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(DelegateBW{
			From:     from,
			Receiver: receiver,
			StakeNet: stakeNet,
			StakeCPU: stakeCPU,
			Transfer: yta.Bool(transfer),
		}),
	}
}

// DelegateBW represents the `eosio.system::delegatebw` action.
type DelegateBW struct {
	From     yta.AccountName `json:"from"`
	Receiver yta.AccountName `json:"receiver"`
	StakeNet yta.Asset       `json:"stake_net"`
	StakeCPU yta.Asset       `json:"stake_cpu"`
	Transfer yta.Bool        `json:"transfer"`
}
