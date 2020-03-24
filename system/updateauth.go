package system

import "github.com/YstarLab/yta-go"

// NewUpdateAuth creates an action from the `eosio.system` contract
// called `updateauth`.
//
// usingPermission needs to be `owner` if you want to modify the
// `owner` authorization, otherwise `active` will do for the rest.
func NewUpdateAuth(account yta.AccountName, permission, parent yta.PermissionName, authority yta.Authority, usingPermission yta.PermissionName) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("updateauth"),
		Authorization: []yta.PermissionLevel{
			{
				Actor:      account,
				Permission: usingPermission,
			},
		},
		ActionData: yta.NewActionData(UpdateAuth{
			Account:    account,
			Permission: permission,
			Parent:     parent,
			Auth:       authority,
		}),
	}

	return a
}

// UpdateAuth represents the hard-coded `updateauth` action.
//
// If you change the `active` permission, `owner` is the required parent.
//
// If you change the `owner` permission, there should be no parent.
type UpdateAuth struct {
	Account    yta.AccountName    `json:"account"`
	Permission yta.PermissionName `json:"permission"`
	Parent     yta.PermissionName `json:"parent"`
	Auth       yta.Authority      `json:"auth"`
}
