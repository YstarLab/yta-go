package rex

import (
	yta "github.com/ystar-foundation/yta-go"
)

func NewDefundCPULoan(from yta.AccountName, loanNumber uint64, amount yta.Asset) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("defcpuloan"),
		Authorization: []yta.PermissionLevel{
			{Actor: from, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(DefundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundCPULoan struct {
	From       yta.AccountName
	LoanNumber uint64
	Amount     yta.Asset
}
