package rex

import (
	yta "github.com/ystar-foundation/yta-go"
)

func NewFundNetLoan(from yta.AccountName, loanNumber uint64, payment yta.Asset) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("fundnetloan"),
		Authorization: []yta.PermissionLevel{
			{Actor: from, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(FundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundNetLoan struct {
	From       yta.AccountName
	LoanNumber uint64
	Payment    yta.Asset
}
