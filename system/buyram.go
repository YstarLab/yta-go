package system

import (
	yta "github.com/ystar-foundation/yta-go"
)

func NewBuyRAM(payer, receiver yta.AccountName, eosQuantity uint64) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("buyram"),
		Authorization: []yta.PermissionLevel{
			{Actor: payer, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(BuyRAM{
			Payer:    payer,
			Receiver: receiver,
			Quantity: yta.NewEOSAsset(int64(eosQuantity)),
		}),
	}
	return a
}

// BuyRAM represents the `eosio.system::buyram` action.
type BuyRAM struct {
	Payer    yta.AccountName `json:"payer"`
	Receiver yta.AccountName `json:"receiver"`
	Quantity yta.Asset       `json:"quant"` // specified in YTA
}
