package bnet

import (
	yta "github.com/YstarLab/yta-go"
	"github.com/YstarLab/yta-go/ecc"
)

type BNetMessageType byte

const (
	HelloType                    BNetMessageType = iota // 0
	TransactionNoticeType                               // 1
	BlockNoticeType                                     // 2
	SignedBlockPointerType                              // 3
	PackedTransactionPointerType                        // 4
	PingType                                            // 5
	PongType                                            // 6
)

//
// Message types
//

type Hello struct {
	PeerID                   ecc.PublicKey     `json:"peer_id"`
	NetworkVersion           string            `json:"network_version"`
	User                     string            `json:"user"`
	Password                 string            `json:"password"`
	Agent                    string            `json:"agent"`
	ProtocolVersion          string            `json:"protocol_version"`
	ChainID                  yta.Checksum256   `json:"chain_id"`
	RequestTransactions      yta.Bool          `json:"request_transactions"`
	LastIrreversibleBlockNum uint32            `json:"last_irr_block_num"`
	PendingBlockIDs          []yta.Checksum256 `json:"pending_block_ids"`
}

/**
 * From bnet_plugin.cpp's `trx_notice`:
 *
 * This message is sent upon successful speculative application of a transaction
 * and informs a peer not to send this message.
 */
type TransactionNotice struct {
	SignedTransactionIDs []yta.Checksum256 ///< hash of trx + sigs
}

/**
 * From bnet_plugin.cpp's `block_notice`:
 *
 * This message is sent upon successfully adding a transaction to the fork database
 * and informs the remote peer that there is no need to send this block.
 */
type BlockNotice struct {
	BlockIDs []yta.Checksum256 `json:"block_ids"`
}

type Ping struct {
	Sent                  yta.Tstamp      `json:"sent"`
	Code                  yta.Checksum256 `json:"code"`
	LastIrreversibleBlock uint32          `json:"lib"` /// last irreversible block
}

type Pong struct {
	Sent yta.Tstamp      `json:"sent"`
	Code yta.Checksum256 `json:"code"`
}

// Also use `yta.SignedBlock`
// Also use `yta.SignedTransaction`
var SignedBlock = yta.SignedTransaction{}
