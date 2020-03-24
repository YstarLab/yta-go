package system

import (
	yta "github.com/YstarLab/yta-go"
)

func NewBidname(bidder, newname yta.AccountName, bid yta.Asset) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("bidname"),
		Authorization: []yta.PermissionLevel{
			{Actor: bidder, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(Bidname{
			Bidder:  bidder,
			Newname: newname,
			Bid:     bid,
		}),
	}
	return a
}

// Bidname represents the `eosio.system_contract::bidname` action.
type Bidname struct {
	Bidder  yta.AccountName `json:"bidder"`
	Newname yta.AccountName `json:"newname"`
	Bid     yta.Asset       `json:"bid"` // specified in YTA
}
