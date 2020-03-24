package token

import "github.com/YstarLab/yta-go"

func init() {
	yta.RegisterAction(AN("eosio.token"), ActN("transfer"), Transfer{})
	yta.RegisterAction(AN("eosio.token"), ActN("issue"), Issue{})
	yta.RegisterAction(AN("eosio.token"), ActN("create"), Create{})
}
