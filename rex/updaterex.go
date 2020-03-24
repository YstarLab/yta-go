package rex

import (
	yta "github.com/YstarLab/yta-go"
)

func NewUpdateREX(owner yta.AccountName) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("updaterex"),
		Authorization: []yta.PermissionLevel{
			{Actor: owner, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(UpdateREX{
			Owner: owner,
		}),
	}
}

type UpdateREX struct {
	Owner yta.AccountName
}
