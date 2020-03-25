package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	yta "github.com/ystar-foundation/yta-go"
)

func NewSetContract(account yta.AccountName, wasmPath, abiPath string) (out []*yta.Action, err error) {
	codeAction, err := NewSetCode(account, wasmPath)
	if err != nil {
		return nil, err
	}

	abiAction, err := NewSetABI(account, abiPath)
	if err != nil {
		return nil, err
	}

	return []*yta.Action{codeAction, abiAction}, nil
}

func NewSetCode(account yta.AccountName, wasmPath string) (out *yta.Action, err error) {
	codeContent, err := ioutil.ReadFile(wasmPath)
	if err != nil {
		return nil, err
	}

	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("setcode"),
		Authorization: []yta.PermissionLevel{
			{
				Actor:      account,
				Permission: yta.PermissionName("active"),
			},
		},
		ActionData: yta.NewActionData(SetCode{
			Account:   account,
			VMType:    0,
			VMVersion: 0,
			Code:      yta.HexBytes(codeContent),
		}),
	}, nil
}

func NewSetABI(account yta.AccountName, abiPath string) (out *yta.Action, err error) {
	abiContent, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}

	var abiPacked []byte
	if len(abiContent) > 0 {
		var abiDef yta.ABI
		if err := json.Unmarshal(abiContent, &abiDef); err != nil {
			return nil, fmt.Errorf("unmarshal ABI file: %s", err)
		}

		abiPacked, err = yta.MarshalBinary(abiDef)
		if err != nil {
			return nil, fmt.Errorf("packing ABI: %s", err)
		}
	}

	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("setabi"),
		Authorization: []yta.PermissionLevel{
			{
				Actor:      account,
				Permission: yta.PermissionName("active"),
			},
		},
		ActionData: yta.NewActionData(SetABI{
			Account: account,
			ABI:     yta.HexBytes(abiPacked),
		}),
	}, nil
}

// NewSetCodeTx is _deprecated_. Use NewSetContract instead, and build
// your transaction yourself.
func NewSetCodeTx(account yta.AccountName, wasmPath, abiPath string) (out *yta.Transaction, err error) {
	actions, err := NewSetContract(account, wasmPath, abiPath)
	if err != nil {
		return nil, err
	}
	return &yta.Transaction{Actions: actions}, nil
}

// SetCode represents the hard-coded `setcode` action.
type SetCode struct {
	Account   yta.AccountName `json:"account"`
	VMType    byte            `json:"vmtype"`
	VMVersion byte            `json:"vmversion"`
	Code      yta.HexBytes    `json:"code"`
}

// SetABI represents the hard-coded `setabi` action.
type SetABI struct {
	Account yta.AccountName `json:"account"`
	ABI     yta.HexBytes    `json:"abi"`
}
