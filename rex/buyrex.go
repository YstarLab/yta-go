package rex

import (
	yta "github.com/ystar-foundation/yta-go"
)

func NewBuyREX(from yta.AccountName, amount yta.Asset) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("buyrex"),
		Authorization: []yta.PermissionLevel{
			{Actor: from, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(BuyREX{
			From:   from,
			Amount: amount,
		}),
	}
}

type BuyREX struct {
	From   yta.AccountName
	Amount yta.Asset
}
