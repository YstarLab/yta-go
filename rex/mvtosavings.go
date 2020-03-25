package rex

import (
	yta "github.com/ystar-foundation/yta-go"
)

func NewMoveToSavings(owner yta.AccountName, rex yta.Asset) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("mvtosavings"),
		Authorization: []yta.PermissionLevel{
			{Actor: owner, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(MoveToSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveToSavings struct {
	Owner yta.AccountName
	REX   yta.Asset
}
