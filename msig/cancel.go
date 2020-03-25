package msig

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewCancel returns a `cancel` action that lives on the
// `eosio.msig` contract.
func NewCancel(proposer yta.AccountName, proposalName yta.Name, canceler yta.AccountName) *yta.Action {
	return &yta.Action{
		Account: yta.AccountName("eosio.msig"),
		Name:    yta.ActionName("cancel"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []yta.PermissionLevel{
			{Actor: canceler, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Cancel{proposer, proposalName, canceler}),
	}
}

type Cancel struct {
	Proposer     yta.AccountName `json:"proposer"`
	ProposalName yta.Name        `json:"proposal_name"`
	Canceler     yta.AccountName `json:"canceler"`
}
