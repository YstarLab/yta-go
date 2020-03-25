package system

import (
	yta "github.com/ystar-foundation/yta-go"
)

// NewBuyRAMBytes will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewBuyRAMBytes(payer, receiver yta.AccountName, bytes uint32) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("buyrambytes"),
		Authorization: []yta.PermissionLevel{
			{Actor: payer, Permission: yta.PermissionName("active")},
		},
		ActionData: yta.NewActionData(BuyRAMBytes{
			Payer:    payer,
			Receiver: receiver,
			Bytes:    bytes,
		}),
	}
	return a
}

// BuyRAMBytes represents the `eosio.system::buyrambytes` action.
type BuyRAMBytes struct {
	Payer    yta.AccountName `json:"payer"`
	Receiver yta.AccountName `json:"receiver"`
	Bytes    uint32          `json:"bytes"`
}
