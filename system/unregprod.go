package system

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewUnregProducer returns a `unregprod` action that lives on the
// `eosio.system` contract.
func NewUnregProducer(producer yta.AccountName) *yta.Action {
	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("unregprod"),
		Authorization: []yta.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(UnregProducer{
			Producer: producer,
		}),
	}
}

// UnregProducer represents the `eosio.system::unregprod` action
type UnregProducer struct {
	Producer yta.AccountName `json:"producer"`
}
