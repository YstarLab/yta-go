package sudo

import (
	yta "github.com/YstarLab/yta-go"
)

// NewExec creates an `exec` action, found in the `eosio.wrap`
// contract.
//
// Given an `yta.Transaction`, call `yta.MarshalBinary` on it first,
// pass the resulting bytes as `yta.HexBytes` here.
func NewExec(executer yta.AccountName, transaction yta.Transaction) *yta.Action {
	a := &yta.Action{
		Account: yta.AccountName("eosio.wrap"),
		Name:    yta.ActionName("exec"),
		Authorization: []yta.PermissionLevel{
			{Actor: executer, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(Exec{
			Executer:    executer,
			Transaction: transaction,
		}),
	}
	return a
}

// Exec represents the `eosio.system::exec` action.
type Exec struct {
	Executer    yta.AccountName `json:"executer"`
	Transaction yta.Transaction `json:"trx"`
}
