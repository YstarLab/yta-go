package system

import (
	yta "github.com/YstarLab/yta-go"
)

func NewSetRAM(maxRAMSize uint64) *yta.Action {
	a := &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("setram"),
		Authorization: []yta.PermissionLevel{
			{
				Actor:      AN("eosio"),
				Permission: yta.PermissionName("active"),
			},
		},
		ActionData: yta.NewActionData(SetRAM{
			MaxRAMSize: yta.Uint64(maxRAMSize),
		}),
	}
	return a
}

// SetRAM represents the hard-coded `setram` action.
type SetRAM struct {
	MaxRAMSize yta.Uint64 `json:"max_ram_size"`
}
