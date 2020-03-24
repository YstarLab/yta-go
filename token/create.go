package token

import yta "github.com/YstarLab/yta-go"

func NewCreate(issuer yta.AccountName, maxSupply yta.Asset) *yta.Action {
	return &yta.Action{
		Account: AN("eosio.token"),
		Name:    ActN("create"),
		Authorization: []yta.PermissionLevel{
			{Actor: AN("eosio.token"), Permission: PN("active")},
		},
		ActionData: yta.NewActionData(Create{
			Issuer:        issuer,
			MaximumSupply: maxSupply,
		}),
	}
}

// Create represents the `create` struct on the `eosio.token` contract.
type Create struct {
	Issuer        yta.AccountName `json:"issuer"`
	MaximumSupply yta.Asset       `json:"maximum_supply"`
}
