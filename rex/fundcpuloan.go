package rex

import (
	yta "github.com/YstarLab/yta-go"
)

func NewFundCPULoan(from yta.AccountName, loanNumber uint64, payment yta.Asset) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("fundcpuloan"),
		Authorization: []yta.PermissionLevel{
			{Actor: from, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(FundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundCPULoan struct {
	From       yta.AccountName
	LoanNumber uint64
	Payment    yta.Asset
}
