package forum

import (
	yta "github.com/ystar-foundation/yta-go"
)

// CleanProposal is an action to flush proposal and allow RAM used by it.
func NewCleanProposal(cleaner yta.AccountName, proposalName yta.Name, maxCount uint64) *yta.Action {
	a := &yta.Action{
		Account: ForumAN,
		Name:    ActN("clnproposal"),
		Authorization: []yta.PermissionLevel{
			{Actor: cleaner, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(CleanProposal{
			ProposalName: proposalName,
			MaxCount:     maxCount,
		}),
	}
	return a
}

// CleanProposal represents the `eosio.forum::clnproposal` action.
type CleanProposal struct {
	ProposalName yta.Name `json:"proposal_name"`
	MaxCount     uint64   `json:"max_count"`
}
