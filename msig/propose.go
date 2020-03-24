package msig

import (
	yta "github.com/YstarLab/yta-go"
)

// NewPropose returns a `propose` action that lives on the
// `eosio.msig` contract.
func NewPropose(proposer yta.AccountName, proposalName yta.Name, requested []yta.PermissionLevel, transaction *yta.Transaction) *yta.Action {
	return &yta.Action{
		Account: yta.AccountName("eosio.msig"),
		Name:    yta.ActionName("propose"),
		Authorization: []yta.PermissionLevel{
			{Actor: proposer, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Propose{proposer, proposalName, requested, transaction}),
	}
}

type Propose struct {
	Proposer     yta.AccountName       `json:"proposer"`
	ProposalName yta.Name              `json:"proposal_name"`
	Requested    []yta.PermissionLevel `json:"requested"`
	Transaction  *yta.Transaction      `json:"trx"`
}
