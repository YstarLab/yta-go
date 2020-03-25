package token

import yta "github.com/ystar-foundation/yta-go"

func NewTransfer(from, to yta.AccountName, quantity yta.Asset, memo string) *yta.Action {
	return &yta.Action{
		Account: AN("eosio.token"),
		Name:    ActN("transfer"),
		Authorization: []yta.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Transfer represents the `transfer` struct on `eosio.token` contract.
type Transfer struct {
	From     yta.AccountName `json:"from"`
	To       yta.AccountName `json:"to"`
	Quantity yta.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
