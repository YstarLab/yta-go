package forum

import (
	yta "github.com/YstarLab/yta-go"
)

// NewExpire is an action to expire a proposal ahead of its natural death.
func NewExpire(proposer yta.AccountName, proposalName yta.Name) *yta.Action {
	a := &yta.Action{
		Account: ForumAN,
		Name:    ActN("expire"),
		Authorization: []yta.PermissionLevel{
			{Actor: proposer, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Expire{
			ProposalName: proposalName,
		}),
	}
	return a
}

// Expire represents the `eosio.forum::propose` action.
type Expire struct {
	ProposalName yta.Name `json:"proposal_name"`
}
