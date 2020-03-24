package rex

import (
	yta "github.com/YstarLab/yta-go"
)

func NewUnstakeToREX(
	owner yta.AccountName,
	receiver yta.AccountName,
	fromNet yta.Asset,
	fromCPU yta.Asset,
) *yta.Action {
	return &yta.Action{
		Account: REXAN,
		Name:    ActN("unstaketorex"),
		Authorization: []yta.PermissionLevel{
			{Actor: owner, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(UnstakeToREX{
			Owner:    owner,
			Receiver: receiver,
			FromNet:  fromNet,
			FromCPU:  fromCPU,
		}),
	}
}

type UnstakeToREX struct {
	Owner    yta.AccountName
	Receiver yta.AccountName
	FromNet  yta.Asset
	FromCPU  yta.Asset
}
