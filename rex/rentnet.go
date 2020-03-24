package rex

import (
	yta "github.com/YstarLab/yta-go"
)

func NewRentNet(
	from yta.AccountName,
	receiver yta.AccountName,
	loanPayment yta.Asset,
	loanFund yta.Asset,
) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("rentnet"),
		Authorization: []yta.PermissionLevel{
			{Actor: from, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(RentNet{
			From:        from,
			Receiver:    receiver,
			LoanPayment: loanPayment,
			LoanFund:    loanFund,
		}),
	}
}

type RentNet struct {
	From        yta.AccountName
	Receiver    yta.AccountName
	LoanPayment yta.Asset
	LoanFund    yta.Asset
}
