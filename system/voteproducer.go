package system

import "github.com/ystar-foundation/yta-go"

// NewNonce returns a `nonce` action that lives on the
// `eosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `yta-bios` boot process by the
// `eosio.system` contract.
func NewVoteProducer(voter yta.AccountName, proxy yta.AccountName, producers ...yta.AccountName) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("voteproducer"),
		Authorization: []yta.PermissionLevel{
			{Actor: voter, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(
			VoteProducer{
				Voter:     voter,
				Proxy:     proxy,
				Producers: producers,
			},
		),
	}
	return a
}

// VoteProducer represents the `eosio.system::voteproducer` action
type VoteProducer struct {
	Voter     yta.AccountName   `json:"voter"`
	Proxy     yta.AccountName   `json:"proxy"`
	Producers []yta.AccountName `json:"producers"`
}
