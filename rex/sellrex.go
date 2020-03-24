package rex

import (
	yta "github.com/YstarLab/yta-go"
)

func NewSellREX(from yta.AccountName, rex yta.Asset) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("sellrex"),
		Authorization: []yta.PermissionLevel{
			{Actor: from, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(SellREX{
			From: from,
			REX:  rex,
		}),
	}
}

type SellREX struct {
	From yta.AccountName
	REX  yta.Asset
}
