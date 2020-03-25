package rex

import (
	yta "github.com/ystar-foundation/yta-go"
)

func NewREXExec(user yta.AccountName, max uint16) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("rexexec"),
		Authorization: []yta.PermissionLevel{
			{Actor: user, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(REXExec{
			User: user,
			Max:  max,
		}),
	}
}

type REXExec struct {
	User yta.AccountName
	Max  uint16
}
