package rex

import (
	yta "github.com/YstarLab/yta-go"
)

func NewMoveFromSavings(owner yta.AccountName, rex yta.Asset) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("mvfrsavings"),
		Authorization: []yta.PermissionLevel{
			{Actor: owner, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(MoveFromSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveFromSavings struct {
	Owner yta.AccountName
	REX   yta.Asset
}
