package system

import (
	"github.com/ystar-foundation/yta-go"
	"github.com/ystar-foundation/yta-go/ecc"
)

// NewNewAccount returns a `newaccount` action that lives on the
// `eosio.system` contract.
func NewNewAccount(creator, newAccount yta.AccountName, publicKey ecc.PublicKey) *yta.Action {
	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("newaccount"),
		Authorization: []yta.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: yta.Authority{
				Threshold: 1,
				Keys: []yta.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []yta.PermissionLevelWeight{},
			},
			Active: yta.Authority{
				Threshold: 1,
				Keys: []yta.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []yta.PermissionLevelWeight{},
			},
		}),
	}
}

// NewDelegatedNewAccount returns a `newaccount` action that lives on the
// `eosio.system` contract. It is filled with an authority structure that
// delegates full control of the new account to an already existing account.
func NewDelegatedNewAccount(creator, newAccount yta.AccountName, delegatedTo yta.AccountName) *yta.Action {
	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("newaccount"),
		Authorization: []yta.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: yta.Authority{
				Threshold: 1,
				Keys:      []yta.KeyWeight{},
				Accounts: []yta.PermissionLevelWeight{
					yta.PermissionLevelWeight{
						Permission: yta.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
			Active: yta.Authority{
				Threshold: 1,
				Keys:      []yta.KeyWeight{},
				Accounts: []yta.PermissionLevelWeight{
					yta.PermissionLevelWeight{
						Permission: yta.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
		}),
	}
}

// NewCustomNewAccount returns a `newaccount` action that lives on the
// `eosio.system` contract. You can specify your own `owner` and
// `active` permissions.
func NewCustomNewAccount(creator, newAccount yta.AccountName, owner, active yta.Authority) *yta.Action {
	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("newaccount"),
		Authorization: []yta.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner:   owner,
			Active:  active,
		}),
	}
}

// NewAccount represents a `newaccount` action on the `eosio.system`
// contract. It is one of the rare ones to be hard-coded into the
// blockchain.
type NewAccount struct {
	Creator yta.AccountName `json:"creator"`
	Name    yta.AccountName `json:"name"`
	Owner   yta.Authority   `json:"owner"`
	Active  yta.Authority   `json:"active"`
}
