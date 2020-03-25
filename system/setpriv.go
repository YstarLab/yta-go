package system

import yta "github.com/ystar-foundation/yta-go"

// NewSetPriv returns a `setpriv` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `yta-bios` boot process by the
// `eosio.system` contract.
func NewSetPriv(account yta.AccountName) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("setpriv"),
		Authorization: []yta.PermissionLevel{
			{Actor: AN("eosio"), Permission: PN("active")},
		},
		ActionData: yta.NewActionData(SetPriv{
			Account: account,
			IsPriv:  yta.Bool(true),
		}),
	}
	return a
}

// SetPriv sets privileged account status. Used in the bios boot mechanism.
type SetPriv struct {
	Account yta.AccountName `json:"account"`
	IsPriv  yta.Bool        `json:"is_priv"`
}
