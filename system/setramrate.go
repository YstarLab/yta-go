package system

import (
	yta "github.com/ystar-foundation/yta-go"
)

func NewSetRAMRate(bytesPerBlock uint16) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("setram"),
		Authorization: []yta.PermissionLevel{
			{
				Actor:      AN("eosio"),
				Permission: yta.PermissionName("active"),
			},
		},
		ActionData: yta.NewActionData(SetRAMRate{
			BytesPerBlock: bytesPerBlock,
		}),
	}
	return a
}

// SetRAMRate represents the system contract's `setramrate` action.
type SetRAMRate struct {
	BytesPerBlock uint16 `json:"bytes_per_block"`
}
