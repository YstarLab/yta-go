package forum

import (
	yta "github.com/YstarLab/yta-go"
)

// Status is an action to set a status update for a given account on the forum contract.
func NewStatus(account yta.AccountName, content string) *yta.Action {
	a := &yta.Action{
		Account: ForumAN,
		Name:    ActN("status"),
		Authorization: []yta.PermissionLevel{
			{Actor: account, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Status{
			Account: account,
			Content: content,
		}),
	}
	return a
}

// Status represents the `eosio.forum::status` action.
type Status struct {
	Account yta.AccountName `json:"account_name"`
	Content string          `json:"content"`
}
