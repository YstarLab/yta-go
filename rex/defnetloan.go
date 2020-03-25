package rex

import (
	yta "github.com/ystar-foundation/yta-go"
)

func NewDefundNetLoan(from yta.AccountName, loanNumber uint64, amount yta.Asset) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("defnetloan"),
		Authorization: []yta.PermissionLevel{
			{Actor: from, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(DefundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundNetLoan struct {
	From       yta.AccountName
	LoanNumber uint64
	Amount     yta.Asset
}
