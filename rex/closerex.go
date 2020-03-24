package rex

import (
	yta "github.com/YstarLab/yta-go"
)

func NewCloseREX(owner yta.AccountName) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("closerex"),
		Authorization: []yta.PermissionLevel{
			{Actor: owner, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(CloseREX{
			Ownwer: owner,
		}),
	}
}

type CloseREX struct {
	Ownwer yta.AccountName
}
