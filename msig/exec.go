package msig

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewExec returns a `exec` action that lives on the
// `eosio.msig` contract.
func NewExec(proposer yta.AccountName, proposalName yta.Name, executer yta.AccountName) *yta.Action {
	return &yta.Action{
		Account: yta.AccountName("eosio.msig"),
		Name:    yta.ActionName("exec"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []yta.PermissionLevel{
			{Actor: executer, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Exec{proposer, proposalName, executer}),
	}
}

type Exec struct {
	Proposer     yta.AccountName `json:"proposer"`
	ProposalName yta.Name        `json:"proposal_name"`
	Executer     yta.AccountName `json:"executer"`
}
