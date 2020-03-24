package rex

import (
	yta "github.com/YstarLab/yta-go"
)

func NewConsolidate(owner yta.AccountName) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("consolidate"),
		Authorization: []yta.PermissionLevel{
			{Actor: owner, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Consolidate{
			Owner: owner,
		}),
	}
}

type Consolidate struct {
	Owner yta.AccountName
}
