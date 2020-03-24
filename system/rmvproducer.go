package system

import (
	yta "github.com/YstarLab/yta-go"
)

// NewRemoveProducer returns a `rmvproducer` action that lives on the
// `eosio.system` contract.  This is to be called by the consortium of
// BPs, to oust a BP from its place.  If you want to unregister
// yourself as a BP, use `unregprod`.
func NewRemoveProducer(producer yta.AccountName) *yta.Action {
	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("rmvproducer"),
		Authorization: []yta.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: yta.NewActionData(RemoveProducer{
			Producer: producer,
		}),
	}
}

// RemoveProducer represents the `eosio.system::rmvproducer` action
type RemoveProducer struct {
	Producer yta.AccountName `json:"producer"`
}
