package sudo

import yta "github.com/ystar-foundation/yta-go"

func init() {
	yta.RegisterAction(AN("eosio.wrap"), ActN("exec"), Exec{})
}

var AN = yta.AN
var ActN = yta.ActN
