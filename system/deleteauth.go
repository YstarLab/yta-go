package system

import "github.com/ystar-foundation/yta-go"

// NewDeleteAuth creates an action from the `eosio.system` contract
// called `deleteauth`.
//
// You cannot delete the `owner` or `active` permissions.  Also, if a
// permission is still linked through a previous `updatelink` action,
// you will need to `unlinkauth` first.
func NewDeleteAuth(account yta.AccountName, permission yta.PermissionName) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("deleteauth"),
		Authorization: []yta.PermissionLevel{
			{Actor: account, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(DeleteAuth{
			Account:    account,
			Permission: permission,
		}),
	}

	return a
}

// DeleteAuth represents the native `deleteauth` action, reachable
// through the `eosio.system` contract.
type DeleteAuth struct {
	Account    yta.AccountName    `json:"account"`
	Permission yta.PermissionName `json:"permission"`
}
