package system

import (
	yta "github.com/YstarLab/yta-go"
	"github.com/YstarLab/yta-go/ecc"
)

// NewSetPriv returns a `setpriv` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `yta-bios` boot process by the
// `eosio.system` contract.
func NewSetProds(producers []ProducerKey) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("setprods"),
		Authorization: []yta.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: yta.NewActionData(SetProds{
			Schedule: producers,
		}),
	}
	return a
}

// SetProds is present in `eosio.bios` contract. Used only at boot time.
type SetProds struct {
	Schedule []ProducerKey `json:"schedule"`
}

type ProducerKey struct {
	ProducerName    yta.AccountName `json:"producer_name"`
	BlockSigningKey ecc.PublicKey   `json:"block_signing_key"`
}
