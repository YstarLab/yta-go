package system

import "github.com/YstarLab/yta-go"

// NewNonce returns a `nonce` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `yta-bios` boot process by the
// `eosio.system` contract.
func NewNonce(nonce string) *yta.Action {
	a := &yta.Action{
		Account:       AN("eosio"),
		Name:          ActN("nonce"),
		Authorization: []yta.PermissionLevel{
			//{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: yta.NewActionData(Nonce{
			Value: nonce,
		}),
	}
	return a
}
