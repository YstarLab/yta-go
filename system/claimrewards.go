package system

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewClaimRewards will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewClaimRewards(owner yta.AccountName) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("claimrewards"),
		Authorization: []yta.PermissionLevel{
			{Actor: owner, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(ClaimRewards{
			Owner: owner,
		}),
	}
	return a
}

// ClaimRewards represents the `eosio.system::claimrewards` action.
type ClaimRewards struct {
	Owner yta.AccountName `json:"owner"`
}
