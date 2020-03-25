package rex

import yta "github.com/ystar-foundation/yta-go"

func init() {
	yta.RegisterAction(REXAN, ActN("buyrex"), BuyREX{})
	yta.RegisterAction(REXAN, ActN("closerex"), CloseREX{})
	yta.RegisterAction(REXAN, ActN("cnclrexorder"), CancelREXOrder{})
	yta.RegisterAction(REXAN, ActN("consolidate"), Consolidate{})
	yta.RegisterAction(REXAN, ActN("defcpuloan"), DefundCPULoan{})
	yta.RegisterAction(REXAN, ActN("defnetloan"), DefundNetLoan{})
	yta.RegisterAction(REXAN, ActN("deposit"), Deposit{})
	yta.RegisterAction(REXAN, ActN("fundcpuloan"), FundCPULoan{})
	yta.RegisterAction(REXAN, ActN("fundnetloan"), FundNetLoan{})
	yta.RegisterAction(REXAN, ActN("mvfrsavings"), MoveFromSavings{})
	yta.RegisterAction(REXAN, ActN("mvtosavings"), MoveToSavings{})
	yta.RegisterAction(REXAN, ActN("rentcpu"), RentCPU{})
	yta.RegisterAction(REXAN, ActN("rentnet"), RentNet{})
	yta.RegisterAction(REXAN, ActN("rexexec"), REXExec{})
	yta.RegisterAction(REXAN, ActN("sellrex"), SellREX{})
	yta.RegisterAction(REXAN, ActN("unstaketorex"), UnstakeToREX{})
	yta.RegisterAction(REXAN, ActN("updaterex"), UpdateREX{})
	yta.RegisterAction(REXAN, ActN("withdraw"), Withdraw{})
}

var AN = yta.AN
var PN = yta.PN
var ActN = yta.ActN

var REXAN = AN("eosio")
