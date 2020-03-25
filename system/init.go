package system

import (
	"github.com/ystar-foundation/yta-go"
)

func init() {
	yta.RegisterAction(AN("eosio"), ActN("setcode"), SetCode{})
	yta.RegisterAction(AN("eosio"), ActN("setabi"), SetABI{})
	yta.RegisterAction(AN("eosio"), ActN("newaccount"), NewAccount{})
	yta.RegisterAction(AN("eosio"), ActN("delegatebw"), DelegateBW{})
	yta.RegisterAction(AN("eosio"), ActN("undelegatebw"), UndelegateBW{})
	yta.RegisterAction(AN("eosio"), ActN("refund"), Refund{})
	yta.RegisterAction(AN("eosio"), ActN("regproducer"), RegProducer{})
	yta.RegisterAction(AN("eosio"), ActN("unregprod"), UnregProducer{})
	yta.RegisterAction(AN("eosio"), ActN("regproxy"), RegProxy{})
	yta.RegisterAction(AN("eosio"), ActN("voteproducer"), VoteProducer{})
	yta.RegisterAction(AN("eosio"), ActN("claimrewards"), ClaimRewards{})
	yta.RegisterAction(AN("eosio"), ActN("buyram"), BuyRAM{})
	yta.RegisterAction(AN("eosio"), ActN("buyrambytes"), BuyRAMBytes{})
	yta.RegisterAction(AN("eosio"), ActN("linkauth"), LinkAuth{})
	yta.RegisterAction(AN("eosio"), ActN("unlinkauth"), UnlinkAuth{})
	yta.RegisterAction(AN("eosio"), ActN("deleteauth"), DeleteAuth{})
	yta.RegisterAction(AN("eosio"), ActN("rmvproducer"), RemoveProducer{})
	yta.RegisterAction(AN("eosio"), ActN("setprods"), SetProds{})
	yta.RegisterAction(AN("eosio"), ActN("setpriv"), SetPriv{})
	yta.RegisterAction(AN("eosio"), ActN("canceldelay"), CancelDelay{})
	yta.RegisterAction(AN("eosio"), ActN("bidname"), Bidname{})
	// yta.RegisterAction(AN("eosio"), ActN("nonce"), &Nonce{})
	yta.RegisterAction(AN("eosio"), ActN("sellram"), SellRAM{})
	yta.RegisterAction(AN("eosio"), ActN("updateauth"), UpdateAuth{})
	yta.RegisterAction(AN("eosio"), ActN("setramrate"), SetRAMRate{})
	yta.RegisterAction(AN("eosio"), ActN("setalimits"), Setalimits{})
}

var AN = yta.AN
var PN = yta.PN
var ActN = yta.ActN
