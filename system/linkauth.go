package system

import "github.com/ystar-foundation/yta-go"

// NewLinkAuth creates an action from the `eosio.system` contract
// called `linkauth`.
//
// `linkauth` allows you to attach certain permission to the given
// `code::actionName`. With this set on-chain, you can use the
// `requiredPermission` to sign transactions for `code::actionName`
// and not rely on your `active` (which might be more sensitive as it
// can sign anything) for the given operation.
func NewLinkAuth(account, code yta.AccountName, actionName yta.ActionName, requiredPermission yta.PermissionName) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("linkauth"),
		Authorization: []yta.PermissionLevel{
			{
				Actor:      account,
				Permission: yta.PermissionName("active"),
			},
		},
		ActionData: yta.NewActionData(LinkAuth{
			Account:     account,
			Code:        code,
			Type:        actionName,
			Requirement: requiredPermission,
		}),
	}

	return a
}

// LinkAuth represents the native `linkauth` action, through the
// system contract.
type LinkAuth struct {
	Account     yta.AccountName    `json:"account"`
	Code        yta.AccountName    `json:"code"`
	Type        yta.ActionName     `json:"type"`
	Requirement yta.PermissionName `json:"requirement"`
}
