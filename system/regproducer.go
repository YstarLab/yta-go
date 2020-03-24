package system

import (
	yta "github.com/YstarLab/yta-go"
	"github.com/YstarLab/yta-go/ecc"
)

// NewRegProducer returns a `regproducer` action that lives on the
// `eosio.system` contract.
func NewRegProducer(producer yta.AccountName, producerKey ecc.PublicKey, url string, location uint16) *yta.Action {
	return &yta.Action{
		Account: AN("eosio"),
		Name:    ActN("regproducer"),
		Authorization: []yta.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: yta.NewActionData(RegProducer{
			Producer:    producer,
			ProducerKey: producerKey,
			URL:         url,
			Location:    location,
		}),
	}
}

// RegProducer represents the `eosio.system::regproducer` action
type RegProducer struct {
	Producer    yta.AccountName `json:"producer"`
	ProducerKey ecc.PublicKey   `json:"producer_key"`
	URL         string          `json:"url"`
	Location    uint16          `json:"location"` // what,s the meaning of that anyway ?
}
