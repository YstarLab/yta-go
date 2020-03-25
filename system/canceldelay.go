package system

import "github.com/ystar-foundation/yta-go"

// NewCancelDelay creates an action from the `eosio.system` contract
// called `canceldelay`.
//
// `canceldelay` allows you to cancel a deferred transaction,
// previously sent to the chain with a `delay_sec` larger than 0.  You
// need to sign with cancelingAuth, to cancel a transaction signed
// with that same authority.
func NewCancelDelay(cancelingAuth yta.PermissionLevel, transactionID yta.Checksum256) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("canceldelay"),
		Authorization: []yta.PermissionLevel{
			cancelingAuth,
		},
		ActionData: yta.NewActionData(CancelDelay{
			CancelingAuth: cancelingAuth,
			TransactionID: transactionID,
		}),
	}

	return a
}

// CancelDelay represents the native `canceldelay` action, through the
// system contract.
type CancelDelay struct {
	CancelingAuth yta.PermissionLevel `json:"canceling_auth"`
	TransactionID yta.Checksum256     `json:"trx_id"`
}
