package msig

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewUnapprove returns a `unapprove` action that lives on the
// `eosio.msig` contract.
func NewUnapprove(proposer yta.AccountName, proposalName yta.Name, level yta.PermissionLevel) *yta.Action {
	return &yta.Action{
		Account:       yta.AccountName("eosio.msig"),
		Name:          yta.ActionName("unapprove"),
		Authorization: []yta.PermissionLevel{level},
		ActionData:    yta.NewActionData(Unapprove{proposer, proposalName, level}),
	}
}

type Unapprove struct {
	Proposer     yta.AccountName     `json:"proposer"`
	ProposalName yta.Name            `json:"proposal_name"`
	Level        yta.PermissionLevel `json:"level"`
}
