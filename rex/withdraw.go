package rex

import (
	yta "github.com/ystar-foundation/yta-go"
)

func NewWithdraw(owner yta.AccountName, amount yta.Asset) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("withdraw"),
		Authorization: []yta.PermissionLevel{
			{Actor: owner, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Withdraw{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Withdraw struct {
	Owner  yta.AccountName
	Amount yta.Asset
}
