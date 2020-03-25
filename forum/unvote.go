package forum

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewUnVote is an action representing the action to undoing a current vote
func NewUnVote(voter yta.AccountName, proposalName yta.Name) *yta.Action {
	a := &yta.Action{
		Account: ForumAN,
		Name:    ActN("unvote"),
		Authorization: []yta.PermissionLevel{
			{Actor: voter, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(UnVote{
			Voter:        voter,
			ProposalName: proposalName,
		}),
	}
	return a
}

// UnVote represents the `eosio.forum::unvote` action.
type UnVote struct {
	Voter        yta.AccountName `json:"voter"`
	ProposalName yta.Name        `json:"proposal_name"`
}
