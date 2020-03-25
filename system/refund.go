package system

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewRefund returns a `refund` action that lives on the
// `eosio.system` contract.
func NewRefund(owner yta.AccountName) *yta.Action {
	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("refund"),
		Authorization: []yta.PermissionLevel{
			{Actor: owner, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(Refund{
			Owner: owner,
		}),
	}
}

// Refund represents the `eosio.system::refund` action
type Refund struct {
	Owner yta.AccountName `json:"owner"`
}
