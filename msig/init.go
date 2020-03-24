package msig

import (
	"github.com/YstarLab/yta-go"
)

func init() {
	yta.RegisterAction(AN("eosio.msig"), ActN("propose"), &Propose{})
	yta.RegisterAction(AN("eosio.msig"), ActN("approve"), &Approve{})
	yta.RegisterAction(AN("eosio.msig"), ActN("unapprove"), &Unapprove{})
	yta.RegisterAction(AN("eosio.msig"), ActN("cancel"), &Cancel{})
	yta.RegisterAction(AN("eosio.msig"), ActN("exec"), &Exec{})
}

var AN = yta.AN
var PN = yta.PN
var ActN = yta.ActN
