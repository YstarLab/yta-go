package forum

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewPropose is an action to submit a proposal for vote.
func NewPropose(proposer yta.AccountName, proposalName yta.Name, title string, proposalJSON string, expiresAt yta.JSONTime) *yta.Action {
	a := &yta.Action{
		Account: ForumAN,
		Name:    ActN("propose"),
		Authorization: []yta.PermissionLevel{
			{Actor: proposer, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Propose{
			Proposer:     proposer,
			ProposalName: proposalName,
			Title:        title,
			ProposalJSON: proposalJSON,
			ExpiresAt:    expiresAt,
		}),
	}
	return a
}

// Propose represents the `eosio.forum::propose` action.
type Propose struct {
	Proposer     yta.AccountName `json:"proposer"`
	ProposalName yta.Name        `json:"proposal_name"`
	Title        string          `json:"title"`
	ProposalJSON string          `json:"proposal_json"`
	ExpiresAt    yta.JSONTime    `json:"expires_at"`
}
