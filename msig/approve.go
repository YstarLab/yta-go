package msig

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewApprove returns a `approve` action that lives on the
// `eosio.msig` contract.
func NewApprove(proposer yta.AccountName, proposalName yta.Name, level yta.PermissionLevel) *yta.Action {
	return &yta.Action{
		Account:       yta.AccountName("eosio.msig"),
		Name:          yta.ActionName("approve"),
		Authorization: []yta.PermissionLevel{level},
		ActionData:    yta.NewActionData(Approve{proposer, proposalName, level}),
	}
}

type Approve struct {
	Proposer     yta.AccountName     `json:"proposer"`
	ProposalName yta.Name            `json:"proposal_name"`
	Level        yta.PermissionLevel `json:"level"`
}
