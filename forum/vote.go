package forum

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewVote is an action representing a simple vote to be broadcast
// through the chain network.
func NewVote(voter yta.AccountName, proposalName yta.Name, voteValue uint8, voteJSON string) *yta.Action {
	a := &yta.Action{
		Account: ForumAN,
		Name:    ActN("vote"),
		Authorization: []yta.PermissionLevel{
			{Actor: voter, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Vote{
			Voter:        voter,
			ProposalName: proposalName,
			Vote:         voteValue,
			VoteJSON:     voteJSON,
		}),
	}
	return a
}

// Vote represents the `eosio.forum::vote` action.
type Vote struct {
	Voter        yta.AccountName `json:"voter"`
	ProposalName yta.Name        `json:"proposal_name"`
	Vote         uint8           `json:"vote"`
	VoteJSON     string          `json:"vote_json"`
}
