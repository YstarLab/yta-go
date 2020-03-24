package system

import (
	yta "github.com/YstarLab/yta-go"
)

// NewRegProxy returns a `regproxy` action that lives on the
// `eosio.system` contract.
func NewRegProxy(proxy yta.AccountName, isProxy bool) *yta.Action {
	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("regproxy"),
		Authorization: []yta.PermissionLevel{
			{Actor: proxy, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(RegProxy{
			Proxy:   proxy,
			IsProxy: isProxy,
		}),
	}
}

// RegProxy represents the `eosio.system::regproxy` action
type RegProxy struct {
	Proxy   yta.AccountName `json:"proxy"`
	IsProxy bool            `json:"isproxy"`
}
