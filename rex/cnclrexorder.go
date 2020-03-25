package rex

import (
	yta "github.com/ystar-foundation/yta-go"
)

func NewCancelREXOrder(owner yta.AccountName) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("cnclrexorder"),
		Authorization: []yta.PermissionLevel{
			{Actor: owner, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(CancelREXOrder{
			Owner: owner,
		}),
	}
}

type CancelREXOrder struct {
	Owner yta.AccountName
}
