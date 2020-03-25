package forum

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewUnPost is an action undoing a post that is active
func NewUnPost(poster yta.AccountName, postUUID string) *yta.Action {
	a := &yta.Action{
		Account: ForumAN,
		Name:    ActN("unpost"),
		Authorization: []yta.PermissionLevel{
			{Actor: poster, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(UnPost{
			Poster:   poster,
			PostUUID: postUUID,
		}),
	}
	return a
}

// UnPost represents the `eosio.forum::unpost` action.
type UnPost struct {
	Poster   yta.AccountName `json:"poster"`
	PostUUID string          `json:"post_uuid"`
}
