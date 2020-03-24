package rex

import (
	yta "github.com/YstarLab/yta-go"
)

func NewDeposit(owner yta.AccountName, amount yta.Asset) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("deposit"),
		Authorization: []yta.PermissionLevel{
			{Actor: owner, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Deposit{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Deposit struct {
	Owner  yta.AccountName
	Amount yta.Asset
}
