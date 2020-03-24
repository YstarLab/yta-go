package system

import "github.com/YstarLab/yta-go"

// NewUnlinkAuth creates an action from the `eosio.system` contract
// called `unlinkauth`.
//
// `unlinkauth` detaches a previously set permission from a
// `code::actionName`. See `linkauth`.
func NewUnlinkAuth(account, code yta.AccountName, actionName yta.ActionName) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("unlinkauth"),
		Authorization: []yta.PermissionLevel{
			{
				Actor:      account,
				Permission: yta.PermissionName("active"),
			},
		},
		ActionData: yta.NewActionData(UnlinkAuth{
			Account: account,
			Code:    code,
			Type:    actionName,
		}),
	}

	return a
}

// UnlinkAuth represents the native `unlinkauth` action, through the
// system contract.
type UnlinkAuth struct {
	Account yta.AccountName `json:"account"`
	Code    yta.AccountName `json:"code"`
	Type    yta.ActionName  `json:"type"`
}
