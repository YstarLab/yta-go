package token

import yta "github.com/YstarLab/yta-go"

func NewIssue(to yta.AccountName, quantity yta.Asset, memo string) *yta.Action {
	return &yta.Action{
		Account: AN("eosio.token"),
		Name:    ActN("issue"),
		Authorization: []yta.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: yta.NewActionData(Issue{
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Issue represents the `issue` struct on the `eosio.token` contract.
type Issue struct {
	To       yta.AccountName `json:"to"`
	Quantity yta.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
